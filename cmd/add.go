package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ludviglundgren/deluge-automanage/internal/config"

	delugeClient "github.com/gdm85/go-libdeluge"
	"github.com/spf13/cobra"
)

func RunAdd() *cobra.Command {
	var paused bool

	var command = &cobra.Command{
		Use:   "add",
		Short: "Add torrent",
		Long:  `Add new torrent to Deluge`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a torrent file as first argument")
			}

			return nil
		},
	}
	command.Flags().BoolVarP(&paused, "paused", "", false, "Add torrent in paused state")

	command.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("Add new torrent")

		// args
		// first arg is path to torrent file
		filePath := args[0]

		var deluge delugeClient.DelugeClient

		if config.Deluge.Version == "v2" {
			deluge = delugeClient.NewV2(delugeClient.Settings{
				Hostname: config.Deluge.Host,
				Port:     config.Deluge.Port,
				Login:    config.Deluge.Login,
				Password: config.Deluge.Password,
			})

		} else {
			deluge = delugeClient.NewV1(delugeClient.Settings{
				Hostname: config.Deluge.Host,
				Port:     config.Deluge.Port,
				Login:    config.Deluge.Login,
				Password: config.Deluge.Password,
			})
		}

		// perform connection to Deluge server
		err := deluge.Connect()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: connection failed: %v\n", err)
			os.Exit(1)
		}
		defer deluge.Close()

		torrentFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			os.Exit(1)
		}

		// check against rules
		activeDownloads, err := deluge.TorrentsStatus(delugeClient.StateDownloading, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: could not list all torrents: %v\n", err)
			os.Exit(1)
		}

		if len(activeDownloads) >= config.Rules.MaxActiveDownloads {
			fmt.Print("too many active downloads")
			os.Exit(1)
		}

		// encode file to base64 before sending to deluge
		encodedFile := base64.StdEncoding.EncodeToString(torrentFile)

		options := delugeClient.Options{
			AddPaused: &paused,
			// Add download save path
		}

		torrentHash, err := deluge.AddTorrentFile(filePath, encodedFile, &options)
		if err != nil {
			os.Exit(1)
		}

		fmt.Printf("Torrent successfully added! Torrenthash: %v\n", torrentHash)
	}

	return command
}
