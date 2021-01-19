#!/usr/bin/env bash

PID="$(pgrep wf-recorder)"

if [[ -n "$PID" ]]; then
    echo ""
    echo ""
fi
