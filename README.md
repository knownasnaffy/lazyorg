<h1 align="center" style="font-size: 3rem; font-weight: bold; margin-bottom: 0.5rem;">
  Lazyorg
</h1>

<p align="center" style="margin: 0;">
  <img src="./assets/logo.svg" alt="Lazyorg Logo" width="25%" style="display: inline-block;">
</p>

<div align="center" style="font-size: 1.2rem; margin-top: 1rem; color: #555;">
  A simple terminal-based calendar and note-taking application.
</div>

<hr style="margin: 2rem auto; width: 80%; border: 1px solid #ddd;">

<p align="center">
  <img src="./assets/demo.gif" alt="Lazyorg Demo" width="100%" style="border-radius: 8px; box-shadow: 0 4px 8px rgba(0,0,0,0.1);">
</p>

> Fork of the original [HubertBel/lazyorg](https://github.com/HubertBel/lazyorg). Modified to suit my jkl; style keybinds instead of hjkl ones. Also changes button and input colors for better readability with tokyonight theme for the terminal.in


## Features
- 📅 Terminal-based calendar interface
- ✨ Event creation and management
- 📝 Integrated simple notepad
- ⌨️ Vim-style keybindings

## Installation

### Prerequisites
- Go 1.23 or higher

### Arch
```bash
yay -S lazyorg-bin
```

### Docker Image

```bash
docker pull defnotgustavom/lazyorg
docker run -it --log-driver none --cap-drop=ALL --net none --security-opt=no-new-privileges --name lazyorg -v /usr/share/zoneinfo/Your/Location:/usr/share/zoneinfo/Your/Location:ro -e TZ=Your/Location defnotgustavom/lazyorg
```
Switch **Your/location** to your current location. Use ```timedatectl list-timezones``` to fetch a list of possible locations.

To rerun the container:
```bash
docker start -ai lazyorg
```
### Binary Installation
Download pre-compiled binary from the latest release.
MacOS and Windows have not been tested yet.

### From Source
```bash
git clone https://github.com/HubertBel/lazyorg.git
cd lazyorg
go build
```

## Usage

### Navigation
- `h/l` - Previous/Next day
- `H/L` - Previous/Next week
- `j/k` - Move time cursor down/up

### Events
- `a` - Add new event
- `d` - Delete current event
- `D` - Delete all events with same name

When creating a new event (`a`), you'll be prompted to fill in the following fields:
- **Name**: Title of event
- **Time**: Date and time of the event
- **Location** (optional): Location of the event
- **Duration**: Duration of the event in hours (0.5 is 30 minutes)
- **Frequency**: The frequency of the event in days, by default 7 or once a week
- **Occurence**: The number of occurence of the event, by default 1
- **Description** (optional): Additional notes or details about the event

### View Controls
- `Ctrl+s` - Show/Hide side view
- `Ctrl+n` - Open/Close notepad
- `Ctrl+r` - Clear notepad content
- `?` - Toggle help menu

### Global
- `Ctrl+c` - Quit

## Configuration

Configuration file will come in future releases. For now, when you open the app for the first time, the database is created at `~/.local/share/lazyorg/data.db` in case you want to do a backup.

## Contributing
Please feel free to submit a Pull Request!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/NewFeature`)
3. Commit your Changes (`git commit -m 'Add some NewFeature'`)
4. Push to the Branch (`git push origin feature/NewFeature`)
5. Open a Pull Request

## Acknowledgments
- Inspired by [lazygit](https://github.com/jesseduffield/lazygit)
- Built with [gocui](https://github.com/jroimartin/gocui) TUI framework
- Thanks to _defnotgustavom_ for the [docker image](https://hub.docker.com/r/defnotgustavom/lazyorg)
- Thanks to [zeckrust](https://github.com/zeckrust) for the logo

## Roadmap

- [ ] Time range modification
- [ ] CLI help
- [ ] Undo/Redo
- [ ] Configuration file
- [ ] Synchronization between devices
- [ ] Import calendar from other apps (Google Calendar, Outlook)

## Buy me a coffee
<a href="https://www.buymeacoffee.com/hubertb" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>
