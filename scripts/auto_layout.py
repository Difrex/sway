#!/usr/bin/env python

import json
import os
import sys


LAYOUTS = {
    "manual": "[M]",
    "spiral": "[S]",
    "left": "[L]",
    "failed": "[DOWN]"
}


def switch(layout):
    out = os.popen("/home/difrex/.local/bin/swaymgr -s 'set " + layout + "'").read()
    if "ok" in out:
        os.popen(f"notify-send Swaymgr 'Switched to <b>{layout}</b> layout'")
    else:
        os.popen(f"notify-send Swaymgr 'Cannot switch to <b>{layout}</b> layout'")


def auto_layout_cmd(cmd):
    return os.popen("/home/difrex/.local/bin/swaymgr -s '" + cmd + "'").read()


def get_current_layout():
    try:
        layout = json.loads(auto_layout_cmd("get layout"))
    except Exception:
        print(LAYOUTS["failed"])
        return
    if not layout["managed"]:
        print(LAYOUTS["manual"])
    else:
        print(LAYOUTS[layout["layout"]])


if __name__ == "__main__":
    if len(sys.argv) > 1:
        switch(sys.argv[1])
    else:
        get_current_layout()
