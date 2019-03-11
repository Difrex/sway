#!/usr/bin/python

import json
import os


def get_inputs():
    return json.loads(os.popen("swaymsg -t get_inputs").read())


def get_current_layout():
    inputs = get_inputs()
    for inp in inputs:
        if inp.get("identifier", None) == "1118:1974:Microsoft_Comfort_Curve_Keyboard_3000":
            if inp.get("xkb_active_layout_name", None) == "Russian":
                return "Ru"
    for inp in inputs:
        if inp.get("identifier", None) == "1452:627:Apple_Inc._Apple_Internal_Keyboard_/_Trackpad":
            if inp.get("xkb_active_layout_name", None) == "Russian":
                return "Ru"
    return "En"


if __name__ == "__main__":
    print(get_current_layout())
