#!/bin/bash
export WAYLAND_DISPLAY=wayland-0
SWAYSOCK=$(/usr/bin/sway --get-socketpath)

if [[ -z $SWAYSOCK ]]; then
    echo "SWAYSOCK is not set"
    socket=$(find /run/user/1000 -type s -name 'sway-ipc*.sock' -print)
    export SWAYSOCK=$socket
    echo "SOCKET=${SWAYSOCK}"
fi
