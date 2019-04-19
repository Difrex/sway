#!/usr/bin/env python3
#
# Copyright (C) 2017 Marcel Patzwahl
# Licensed under the terms of the GNU GPL v3 only.
#
# i3blocks blocklet script to see the available updates of pacman and the AUR
import subprocess
from subprocess import check_output
import argparse
import re


def create_argparse():
    parser = argparse.ArgumentParser(description='Check for pacman updates')
    parser.add_argument(
        '-b',
        '--base_color',
        default='green',
        help='base color of the output(default=green)'
    )
    parser.add_argument(
        '-u',
        '--updates_available_color',
        default='yellow',
        help='color of the output, when updates are available(default=yellow)'
    )
    parser.add_argument(
        '-a',
        '--aur',
        action='store_true',
        help='Include AUR packages. Attn: Yaourt must be installed'
    )
    parser.add_argument(
        '-q',
        '--quiet',
        action='store_true',
        help='Do not produce output when system is up to date'
    )
    parser.add_argument(
        '-w',
        '--watch',
        nargs='*',
        default=[],
        help='Explicitly watch for specified packages. '
        'Listed elements are treated as regular expressions for matching.'
    )
    return parser.parse_args()


def get_updates():
    output = None
    try:
        output = check_output(['checkupdates']).decode('utf-8')
    except Exception:
        pass
    if not output:
        return []

    updates = [line.split(' ')[0]
               for line in output.split('\n')
               if line]

    return updates


def get_aur_updates():
    output = ''
    try:
        output = check_output(['yaourt', '-Qua']).decode('utf-8')
    except subprocess.CalledProcessError as exc:
        # yaourt exits with 1 and no output if no updates are available.
        # we ignore this case and go on
        if not (exc.returncode == 1 and not exc.output):
            raise exc
    if not output:
        return []

    aur_updates = [line.split(' ')[0]
                   for line in output.split('\n')
                   if line.startswith('aur/')]

    return aur_updates


def matching_updates(updates, watch_list):
    matches = set()
    for u in updates:
        for w in watch_list:
            if re.match(w, u):
                matches.add(u)

    return matches


message = "<span color='{0}'>{1}</span>"
args = create_argparse()

updates = get_updates()
if args.aur:
    updates += get_aur_updates()

update_count = len(updates)
if update_count > 0:
    info = str(update_count) + ' updates available'
    # matches = matching_updates(updates, args.watch)
    # if matches:
    #     info += ' [{0}]'.format(', '.join(matches))
    print(message.format(args.updates_available_color, info))
elif not args.quiet:
    print(message.format(args.base_color, 'system up to date'))
