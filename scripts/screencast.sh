#!/usr/bin/env bash

file_name() {
    echo "screencast-$(date "+%d-%m-%y %H-%M").mp4"
}

case $1 in
    r)
        wf-recorder -f ~/Видео/screencasts/"$(file_name)" -g "$(slurp)"
        ;;
    a)
        wf-recorder -f ~/Видео/screencasts/"$(file_name)"
        ;;
    s)
        killall wf-recorder
        ;;
esac
