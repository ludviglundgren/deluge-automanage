# deluge-automanage

A cli to automate Deluge. Currently very minimal.

## Install

Download binary and put somewhere in $PATH.

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