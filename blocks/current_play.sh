#!/usr/bin/env bash

echo $BLOCK_BUTTON > /tmp/block_button

GRADIO=true

# Pass the password in the block instance
if [[ -n $BLOCK_INSTANCE ]]; then
    password=("-h" "$BLOCK_INSTANCE@localhost")
fi

mpd_status=$(mpc $password status | tr '\n' ' ' | grep -Po '.*(?= \[playing\])|paused' | tr -d '\n')
pctl_status=$(playerctl status)

update_status() {
    mpd_status=$(mpc $password status | tr '\n' ' ' | grep -Po '.*(?= \[playing\])|paused' | tr -d '\n')
    pctl_status=$(playerctl status)
}

filter() {
    update_status

    echo -n '['
    status=$mpd_status
    if [[ -z $mpd_status ]] || [[ $mpd_status == "paused" ]]; then
        if [[ $GRADIO ]]; then
            if [[ $pctl_status == "Playing" ]]; then
                status="radio"
            fi
        fi
    fi
    echo -n $status
    echo -n ']'

}

play_pause() {
    update_status

    if [[ $GRADIO ]] && [[ $pctl_status == "Playing" ]]; then
        playerctl pause
    else
        mpc $password toggle
    fi

    update_status
}

case $1 in
    2) play_pause ;;
    1) mpc $password prev   | filter ;;
    3) mpc $password next   | filter ;;
    *) mpc $password status | filter ;;
esac
