#!/usr/bin/env python
import subprocess
import sys


def set_volume(v):
    chan = "2"
    p = subprocess.Popen(["pactl", "set-sink-volume", chan, v])
    p.communicate()
    if p.returncode > 0:
        p = subprocess.Popen(["pactl", "set-sink-volume", "0", v])
        p.communicate()


def mute():
    chan = "1"
    p = subprocess.Popen(["pactl", "set-sink-mute", chan, "toggle"])
    p.communicate()
    if p.returncode > 0:
        p = subprocess.Popen(["pactl", "set-sink-mute", "1", "toggle"])
        p.communicate()


if __name__ == "__main__":
    if len(sys.argv) > 1:
        if sys.argv[1] == "mute":
            mute(sys.argv[1])
        else:
            set_volume(sys.argv[1])
