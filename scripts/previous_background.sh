#!/usr/bin/env bash

BACKGROUND="$(cat ~/.config/sway/.current_wallpaper)"

# Set it
swaymsg output '*' background "${BACKGROUND}" fill
