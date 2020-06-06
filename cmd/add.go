package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	delugeclient "github.com/gdm85/go-libdeluge"
	"github.com/spf13/cobra"
)

var Paused bool

func init() {
	addCmd.Flags().BoolVarP(&Paused, "paused", "", false, "Add torrent in paused state")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add torrent",
	Long:  `Add new torrent to Deluge`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a torrent file as first argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add new torrent")

		// args
		// first arg is path to torrent file
		filePath := args[0]

		deluge := delugeclient.NewV1(delugeclient.Settings{
			Hostname: Config.Deluge.Host,
			Port:     Config.Deluge.Port,
			Login:    Config.Deluge.Login,
			Password: Config.Deluge.Password,
		})

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
		activeDownloads, err := deluge.TorrentsStatus(delugeclient.StateDownloading, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: could not list all torrents: %v\n", err)
			os.Exit(1)
		}

		if len(activeDownloads) >= Config.Rules.MaxActiveDownloads {
			fmt.Print("too many active downloads")
			os.Exit(1)
		}

		// encode file to base64 before sending to deluge
		encodedFile := base64.StdEncoding.EncodeToString(torrentFile)

		options := delugeclient.Options{
			AddPaused: &Paused,
			// Add download save path
		}

		torrentHash, err := deluge.AddTorrentFile(filePath, encodedFile, &options)
		if err != nil {
			os.Exit(1)
		}

		fmt.Printf("Torrent successfully added! Torrenthash: %v\n", torrentHash)
	},
}
