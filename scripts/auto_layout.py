#!/usr/bin/env python

import json
import os


LAYOUTS = {
    "manual": "[M]",
    "spiral": "[S]"
}


def auto_layout_cmd(cmd):
    return os.popen("/home/difrex/.local/bin/swaymgr -s '" + cmd + "'").read()


def get_current_layout():
    layout = json.loads(auto_layout_cmd("get layout"))
    if not layout["managed"]:
        print(LAYOUTS["manual"])
    else:
        print(LAYOUTS[layout["layout"]])


if __name__ == "__main__":
    get_current_layout()
