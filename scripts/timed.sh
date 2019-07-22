#!/usr/bin/env bash

SETTIMED_PID=$(pgrep settimed)
SWAY_DIR="${HOME}/.config/sway"
LOGFILE="${SWAY_DIR}/.timed.log"
CONFIG="${SWAY_DIR}/mojave-timed.stw"
SETTIMED="${HOME}/.local/bin/settimed"

if [[ -z $SETTIMED_PID ]]; then
    $SETTIMED mojave-timed > "$LOGFILE"
else
    kill $SETTIMED_PID
    $SETTIMED mojave-timed > "$LOGFILE"
fi
