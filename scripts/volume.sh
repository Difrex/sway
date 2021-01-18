#!/usr/bin/env bash

RP_CONNECTED="$(pactl list | grep RP)"
PODS_CONNECTED="$(pactl list | grep "AirPods")"
CHANNEL="3"

if [[ -n "$RP_CONNECTED" ]]; then
    CHANNEL="bluez_sink.00_0B_97_45_42_21.a2dp_sink"
fi

if [[ -n "$PODS_CONNECTED" ]]; then
    CHANNEL="bluez_sink.AC_1D_06_88_41_20.a2dp_sink"
fi

pactl set-sink-volume "$CHANNEL" "$1"
