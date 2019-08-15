#!/usr/bin/env bash

case $1 in dark)
               rm -f $HOME/.config/sway/conf-enabled/colors-light.conf
               rm -f $HOME/.config/sway/conf-enabled/terminal-light.conf
               rm -f $HOME/.config/sway/conf-enabled/mako-light.conf
               ln -s $HOME/.config/sway/conf-available/colors-dark.conf $HOME/.config/sway/conf-enabled/colors-dark.conf
               ln -s $HOME/.config/sway/conf-available/terminal-dark.conf $HOME/.config/sway/conf-enabled/terminal-dark.conf
               ln -s $HOME/.config/sway/conf-available/mako-dark.conf $HOME/.config/sway/conf-enabled/mako-dark.conf
               rm -f $HOME/.config/sway/waybar-enabled/style.css
               ln -s $HOME/.config/sway/waybar-available/style-dark.css $HOME/.config/sway/waybar-enabled/style.css
               swaymsg reload

               # Emacs
               emacsclient -e '(set-dark-theme)'

               # GTK
               gsettings set org.gnome.desktop.interface gtk-theme 'Numix Solarized'
               gsettings set org.gnome.desktop.interface icon-theme 'breeze-dark'

               # Terminals
               wal -e -f solarized

               # Rofi
               sed -i 's@rofi.theme: /usr/share/rofi/themes/gruvbox-light-soft.rasi@rofi.theme: /usr/share/rofi/themes/solarized.rasi@' $HOME/.config/rofi/config
               ;;
            light)
               rm -f $HOME/.config/sway/conf-enabled/colors-dark.conf
               rm -f $HOME/.config/sway/conf-enabled/terminal-dark.conf
               rm -f $HOME/.config/sway/conf-enabled/mako-dark.conf
               ln -s $HOME/.config/sway/conf-available/colors-light.conf $HOME/.config/sway/conf-enabled/colors-light.conf
               ln -s $HOME/.config/sway/conf-available/terminal-light.conf $HOME/.config/sway/conf-enabled/terminal-light.conf
               ln -s $HOME/.config/sway/conf-available/mako-light.conf $HOME/.config/sway/conf-enabled/mako-light.conf
               rm -f $HOME/.config/sway/waybar-enabled/style.css
               ln -s $HOME/.config/sway/waybar-available/style-light.css $HOME/.config/sway/waybar-enabled/style.css
               swaymsg reload
               emacsclient -e '(set-light-theme)'
               gsettings set org.gnome.desktop.interface gtk-theme 'Numix Solarized Light'
               gsettings set org.gnome.desktop.interface icon-theme 'breeze'
               wal -e -l -f solarized
               sed -i 's@rofi.theme: /usr/share/rofi/themes/solarized.rasi@rofi.theme: /usr/share/rofi/themes/gruvbox-light-soft.rasi@' $HOME/.config/rofi/config
               ;;
            *)
                echo "Nothing to do"
                ;;
esac
