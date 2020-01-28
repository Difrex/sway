# Swaywm configuration

This is my daily setup of the Sway.

## Components

* [Waybar](https://github.com/Alexays/Waybar) -- panel
* [mako](https://github.com/emersion/mako) -- notification daemon
* [swayidle](https://github.com/swaywm/swayidle) -- idle management daemon. Replacement for the `xautolock`
* [rofi](https://github.com/davatorium/rofi) -- application launcher and dmenu replacement
* [swaylock](https://github.com/swaywm/swaylock) -- lock screen

### Installation

```
pacman -S sway swayidle swaybg waybar mako rofi
```

## Configs

Some interesting configs.

### Transparent windows

Build:
```sh
go get github.com/Difrex/gosway/ipc
cd scripts && go build -o ~/.local/bin/tr_in transparent_inctive.go
```

Enable:
```sh
cd conf-enabled
ln -s ../conf-available/transparent-windows.conf
```

### List of floating windows(scratchpads) in Rofi

Build:
```sh
go get github.com/Difrex/gosway/ipc
cd scripts && go build -o ~/.config/sway/scripts/floating_windows floating_windows.go
```

Enable:
```sh
cd conf-enabled
ln -s ../conf-available/scratchpads-select.conf
```

Keybinding: **$mod+Alt+w**

### Autolayouting

You need a [swaymgr](https://github.com/Difrex/swaymgr) installed.

Enable:
```sh
cd conf-enabled
ln -s ../conf-available/autolayout.conf
```

### Other configs

* `autostart.conf` -- start applications on the startup
* `idle.conf` -- swayidle settings
* `wallpapers.conf` -- wallpapers configuration. Keybing to setup random wallpaper is: **$mod+Shift+w**
* `layout.conf` -- keyboard layout configuration

## Themes

I always use Solarized themes.

### Dark colors

```sh
cd conf-enabled
ln -s ../conf-available/colors-dark.conf
ln -s ../conf-available/mako-dark.conf
ln -s ../conf-available/termninal-dark.conf
cd ../waybar-enabled
ln -s ../waybar-available/style-dark.css style.css
```

### Light colors

```sh
cd conf-enabled
ln -s ../conf-available/colors-light.conf
ln -s ../conf-available/mako-light.conf
ln -s ../conf-available/termninal-light.conf
cd ../waybar-enabled
ln -s ../waybar-available/style-light.css style.css
```

## Keybindigns

`$mod` is a Super(with windows logo) key.

* Kill focused window: **$mod+q**
* Launch a terminal: **$mod+Return**
* Rofi launcher: **$mod+d**
* Emacsclient: **$mod+a**
* Copy password with the [pm](https://github.com/himidori/pm): **$mod+p**
* Lock screen with the swaylock: **$mod+l**
* Resize: **$mod+r**
* Make a screenshot(you need the [grim](https://github.com/emersion/grim) and [slurp](https://github.com/emersion/slurp)): **$mod+c**
* Move window to the scratchpad: **$mod+Shift+minus**
* Show scratchpad: **$mod+minus**
* Show modeline for monitors configuration: **$mod+x**

For a theme switching you need a [wal](https://github.com/dylanaraps/pywal)
* Switch to light theme: **$mod+Ctrl+Shift+l**
* Switch to dark theme: **$mod+Ctrl+Shift+d**

* Reload config: **$mod+Shift+c**
* Shutdown a session: **$mod+Shift+e**
