#!/usr/bin/python

import os
import sys
import json
import pyudev
import time


def get_outputs():
    return json.loads(os.popen("swaymsg -t get_outputs").read())


def connected():
    outputs = []
    for out in get_outputs():
        outputs.append(out.get("name", None))
    return outputs


def three_mons(outputs):
    print(outputs)
    if 'DP-1' in outputs and 'HDMI-A-2' in outputs and 'eDP-1' in outputs:
        return True
    return False


def two_mons(outputs):
    if 'HDMI-A-2' and 'eDP-1' in outputs and len(outputs) == 2:
        return True
    return False


def single(outputs):
    if len(outputs) == 1:
        return True
    return False


def three_mons_templ():
    return """output eDP-1 disable
output HDMI-A-2 pos 0 0 transform 270 enable
output DP-1 pos 1920 0 enable
"""


def single_templ():
    return """output eDP-1 scale 1.45 enable
output HDMI-A-2 disable
output DP-1 disable
"""


def enable_hidpi():
    return "output eDP-1 scale 1.45 enable"


def disable_hidpi():
    return "output eDP-1 scale 1.45 disable"


def two_templ():
    return """output eDP-1 pos 1920 0 scale 1.45 enable
output HDMI-A-2 pos 0 0 transform 270 enable
output DP-1 disable
"""


def two_templ_dp1():
    return """output eDP-1 pos 1920 0 scale 1.45 enable
output HDMI-A-2 pos 0 0 disable
output DP-1 pos 0 0 enable
"""


def autoswich(outputs):
    if three_mons(outputs):
        write_templ(three_mons_templ())
        swaymsg(three_mons_templ())
    else:
        write_templ(single_templ())
        swaymsg(single_templ())


def write_templ(templ):
    with open(os.environ.get("HOME", "~/") + "/.config/sway/conf.d/monitors", "w") as f:
        f.write(templ.strip("\n"))


def swaymsg(templ):
    print(json.loads(os.popen("swaymsg '" + templ + "'").read())[0]["success"])


def udev_watcher():
    context = pyudev.Context()
    monitor = pyudev.Monitor.from_netlink(context)
    monitor.filter_by('drm')
    for dev in iter(monitor.poll, None):
        print("Monitor configuration change")
        time.sleep(2)
        outputs = get_outputs()
        if len(outputs) == 3:
            print("Two external monitors detected")
            swaymsg(three_mons_templ())
        elif len(outputs) == 2:
            print("Enable one external monitor")
            for output in outputs:
                if output["name"] == "HDMI-A-2":
                    print("HDMI-A-2 detected")
                    swaymsg(two_templ())
                elif output["name"] == "DP-1":
                    swaymsg(two_templ_dp1())
                    print("DP-1 detected")
        else:
            print("Enable only internal screen")
            swaymsg(single_templ())


if __name__ == "__main__":
    # Cli
    if len(sys.argv) > 1:
        if sys.argv[1] == "1":
            # write_templ(single_templ())
            swaymsg(single_templ())
        elif sys.argv[1] == "2":
            # write_templ(two_templ())
            swaymsg(two_templ())
        elif sys.argv[1] == "3":
            # write_templ(three_mons_templ())
            swaymsg(three_mons_templ())
        elif sys.argv[1] == "h":
            swaymsg(enable_hidpi())
        elif sys.argv[1] == "d":
            swaymsg(disable_hidpi())
    else:
        udev_watcher()
