#!/usr/bin/python

import json
import os


def get_inputs():
    return json.loads(os.popen("swaymsg -t get_inputs").read())


def get_current_layout():
    inputs = get_inputs()
    for inp in inputs:
        if inp.get("xkb_active_layout_name", None) == "Russian":
            return "Ru"
    return "En"


if __name__ == "__main__":
    print(get_current_layout())
