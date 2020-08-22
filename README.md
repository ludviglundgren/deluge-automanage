# deluge-automanage

A cli to automate Deluge. Currently very minimal.

## Install

Download binary and put somewhere in $PATH.

    wget https://github.com/ludviglundgren/deluge-automanage/releases/download/v0.1.1/deluge-automanage_0.1.1_linux_amd64.tar.gz

Extract binary

    tar -xzvf deluge-automanage_0.1.1_linux_amd64.tar.gz

Move to somewhere in `$PATH`. Need sudo if not already root. Or put it in your user `$HOME/bin` or similar.

    sudo mv deluge-automanage /usr/bin/

Verify that it runs

    deluge-automanage help

This should print `Could not read config file`.

## Configuration

Create a new configuration file `.deluge-automanage.toml` in `$HOME/.config/deluge-automanage/`.

    mkdir -p ~/.config/deluge-automanage && touch ~/.config/deluge-automanage/.deluge-automanage.toml

A bare minimum config.

```toml
[deluge]
host     = "localhost"  # deluge daemon hostname/ip
port     = 30000        # deluge daemon port
login    = "my-user"    # deluge daemon user
password = "my-pass"    # deluge daemon password
 
[rules]
enabled              = true   # enable or disable rules
max_active_downloads = 1      # set max active downloads
```

* If running on HDDs and 1Gbit - `max_active_downloads = 2` is a good setting to not overload the disks and gives as much bandwidth as possible to the torrents. 
* For SSDs and 1Gbit+ you can increase this value.

### rutorrent-autodl-irssi setup

In rutrorrent, go to autodl-irssi `Preferences`, and then the `Action` tab. Put in the following for the global action. This can be set in a specific filter as well.

```
Choose .torrent action: Run Program
Command: /usr/bin/deluge-automanage
Arguments: add "$(TorrentPathName)"
```

If you want to grab as many new torrents as possible you can change your filters to grab every torrent that matches and then `max_active_downloads` will decide if it gets added or not, depending on how many active downloads there currently are.

If there are some filters that you want to be 100% sure gets downloaded, then either use the old deluge-console/watch action for those, or do the inverse and use the global action as deluge-console and specific filters with deluge-automanage.

## Usage

Use `deluge-automange help` to find out more about how to use.

Commands:
  - add

Flags:
  * `--config` - use other config file

### Add

Add a new torrent to deluge.

    deluge-automanage add my-torrent-file.torrent

Flags:
  * `--paused` - add torrent in paused state
