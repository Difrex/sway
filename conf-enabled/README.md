# Sway configuration modules


## Enable config module

```sh
cd conf-enabled
ln -s ../conf-available/config.conf
```

## Disable config module

```sh
cd conf-enabled
rm config.conf
```

## Ordered load

```
.
├── 00-workspaces.conf -> ../conf-available/workspaces.conf
├── 01-autostart.conf -> ../conf-available/autostart.conf
├── 02-apps.conf -> ../conf-available/apps.conf
├── 02-fonts.conf -> ../conf-available/fonts.conf
├── 03-keybindings.conf -> ../conf-available/keybindings.conf
├── 04-key-layout.conf -> ../conf-available/layout.conf
├── 05-terminal.conf -> /home/difrex/.config/sway/conf-available/terminal-dark.conf
├── 10-monitors.conf -> ../conf-available/monitors.conf
├── 20-colors.conf -> ../conf-available/colors-dark.conf
├── 30-gaps.conf -> ../conf-available/gaps.conf
├── 40-idle.conf -> ../conf-available/idle.conf
├── 50-screenshots.conf -> ../conf-available/screenshots.conf
├── 51-screencascsts.conf -> ../conf-available/screencasts.conf
├── 60-autolayout.conf -> ../conf-available/autolayout.conf
├── 70-scratchpads-select.conf -> ../conf-available/scratchpads-select.conf
├── 80-wallpapers.conf -> ../conf-available/wallpapers.conf
└── README.md
```
