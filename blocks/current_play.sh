#!/usr/bin/env bash

STATUS="$(playerctl status)"

if [[ "$STATUS" != "Playing" ]]; then
    echo "$STATUS"
else
    artist="$(playerctl metadata xesam:albumArtist)"
    title="$(playerctl metadata xesam:title)"

    echo "${artist} - ${title}"
    echo "${artist} - ${title}"
fi
