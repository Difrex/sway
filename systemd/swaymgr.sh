#!/usr/bin/env bash

export WAYLAND_DISPLAY=wayland-0
source /home/difrex/.config/sway/systemd/envs.sh

/usr/bin/systemctl --user import-environment WAYLAND_DISPLAY
/usr/bin/systemctl --user import-environment XDG_SESSION_TYPE

rm -f /tmp/swaymgr.sock

/usr/bin/swaymgr
