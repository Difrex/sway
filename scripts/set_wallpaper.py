#!/usr/bin/python

import os
import random
import sys


def write_current(path):
    with open(os.environ.get("HOME", "~") + "/.config/sway/.current_wallpaper", "w") as f:
        f.write(path)
        f.flush()
        f.close()


wallpapers_dir = os.environ.get("HOME", "~") + "/.config/sway/wallpapers"
for i in os.walk(wallpapers_dir):
    r = i[0] + "/" + i[2][random.randint(0, len(i[2]))]
    sys.stderr.write(os.popen("swaymsg output '*' background '" + r + "' fill").read() + "\n")
    write_current(r)
    break
