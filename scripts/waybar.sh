#!/usr/bin/env bash

run_waybar() {
    waybar -c ~/.config/sway/waybar-enabled/config -s ~/.config/sway/waybar-enabled/style.css
}

restart_waybar() {
    pid=$(pgrep waybar)
    [ ! -z $pid ] && kill $pid
    run_waybar
}

restart_waybar
