#!/usr/bin/env bash

SETTIMED_PID=$(pgrep settimed)
SWAY_DIR="${HOME}/.config/sway"
LOGFILE="${SWAY_DIR}/.timed.log"
CONFIG="${SWAY_DIR}/mojave-timed.stw"

if [[ -z $SETTIMED_PID ]]; then
    settimed mojave-timed > "$LOGFILE"
else
    kill $SETTIMED_PID
    settimed mojave-timed > "$LOGFILE"
fi
