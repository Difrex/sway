#!/usr/bin/python

import os
import random
import sys


wallpapers_dir = os.environ.get("HOME", None) + "/Изображения/wallpapers/collection/Future"
for i in os.walk(wallpapers_dir):
    r = i[0] + "/" + i[2][random.randint(0, len(i[2]))]
    sys.stderr.write(os.popen("swaymsg output '*' background '" + r + "' fill").read() + "\n")
    print(r)
