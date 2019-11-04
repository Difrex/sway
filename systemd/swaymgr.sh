#!/usr/bin/env bash

/usr/bin/systemctl --user import-environment WAYLAND_DISPLAY
/usr/bin/systemctl --user import-environment XDG_SESSION_TYPE

export WAYLAND_DISPLAY=wayland-0
source /home/difrex/.config/sway/systemd/envs.sh

/home/difrex/.local/bin/swaymgr
