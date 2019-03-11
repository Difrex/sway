#!/bin/bash

hdmi2=`xrandr | grep 'HDMI2 connected'`
dp1=`xrandr | grep -e '^DP1 connected'`

if [[ -z $hdmi2 ]]; then
    xrandr --output eDP1 --mode 1920x1200
elif [[ ! -z $hdmi2 ]] && [[ ! -z $dp1 ]]; then
    xrandr --output eDP1 --off --output HDMI2 --auto --output DP1 --auto --right-of HDMI2
else
    xrandr --output eDP1 --mode 1920x1200 --output HDMI2 --mode 1920x1200 --same-as eD1
fi
