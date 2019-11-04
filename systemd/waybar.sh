#!/usr/bin/env bash

/usr/bin/systemctl --user import-environment WAYLAND_DISPLAY
/usr/bin/systemctl --user import-environment XDG_SESSION_TYPE
/usr/bin/systemctl --user import-environment all

source /home/difrex/.config/sway/systemd/envs.sh
/usr/bin/waybar -c /home/difrex/.config/sway/waybar-enabled/config -s /home/difrex/.config/sway/waybar-enabled/style.css
