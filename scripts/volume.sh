#!/usr/bin/env bash

RP_CONNECTED="$(pactl list | grep RP)"
PODS_CONNECTED="$(pactl list | grep "AirPods")"
WIRED_CONNECTED="$(pactl list | grep -B 2 "Comet Lake PCH-LP cAVS Speaker + Headphones")"

CHANNEL="3"

if [[ -n "$RP_CONNECTED" ]]; then
    CHANNEL="bluez_sink.00_0B_97_45_42_21.a2dp_sink"
fi

if [[ -n "$PODS_CONNECTED" ]]; then
    CHANNEL="bluez_sink.AC_1D_06_88_41_20.a2dp_sink"
fi

if [[ -n "$(echo "$WIRED_CONNECTED" | grep "RUNNING")" ]]; then
    CHANNEL="alsa_output.pci-0000_00_1f.3-platform-skl_hda_dsp_generic.HiFi__hw_sofhdadsp__sink"
fi

pactl set-sink-volume "$CHANNEL" "$1"
