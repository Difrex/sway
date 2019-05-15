#!/usr/bin/env bash

run_waybar() {
    waybar -c ~/.config/sway/waybar/config -s ~/.config/sway/waybar/style.css
}

restart_waybar() {
    pid=$(pgrep waybar)
    [ ! -z $pid ] && kill $pid
    run_waybar
}

restart_waybar
