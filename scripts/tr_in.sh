#!/usr/bin/env bash

PID=$(pgrep tr_in)

if [[ -z $PID ]]; then
    ${HOME}/.local/bin/tr_in
else
    killall tr_in
    ${HOME}/.local/bin/tr_in
fi
