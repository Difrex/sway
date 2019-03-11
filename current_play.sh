#!/bin/bash

# Pass the password in the block instance
if [[ -n $BLOCK_INSTANCE ]]; then
    password=("-h" "$BLOCK_INSTANCE@localhost")
fi

filter() {
    echo -n '['
    tr '\n' ' ' | grep -Po '.*(?= \[playing\])|paused' | tr -d '\n'
    echo -n ']'
}

case $BLOCK_BUTTON in
    2) mpc $password toggle | filter ;;
    1) mpc $password prev   | filter ;;
    3) mpc $password next   | filter ;;
    *) mpc $password status | filter ;;
esac
