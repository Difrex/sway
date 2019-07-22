#!/usr/bin/env python

import json
import os


def run_notify_send(title, body):
    """Runs a notify-send programm"""
    os.popen(f"notify-send '{title}' '{body}'")


def build_command(*args):
    """Returns a string with the makoctl command"""
    cmd = ["makoctl"]
    for arg in args[0]:
        cmd.append(arg)
    return " ".join(cmd)


def run_mako_ctl(*args):
    """Runs the makoctl command with the provided arguments"""
    cmd = build_command(args)
    return json.loads(os.popen(cmd).read())


def get_list():
    """returns a notificatoins list from the makoctl"""
    l = run_mako_ctl("list")
    for i in l["data"]:
        if len(i) > 0:
            for j in i:
                run_notify_send(j["summary"]["data"], j["body"]["data"])
    return


if __name__ == "__main__":
    get_list()
