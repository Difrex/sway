#!/usr/bin/env bash

SETTIMED_PID=$(pgrep settimed)

if [[ -z $SETTIMED_PID ]]; then
    settimed mojave-timed
else
    kill $SETTIMED_PID
    settimed mojave-timed
fi
