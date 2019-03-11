#!/usr/bin/python
from __future__ import division
import sys
import os
import logging
from subprocess import check_output
from operator import methodcaller
from systemd.journal import JournalHandler
from inspect import stack as inspect_stack
from threading import Thread
# ./module/i915/parameters/panel_ignore_lid
# tobackup:
# /etc/acpi/events/lid
# /etc/udev/rules.d/90-display-connect.rules


def logfile(loglevel=logging.DEBUG):
    r_pipe_fd, w_pipe_fd = os.pipe()
    r_pipe = os.fdopen(r_pipe_fd, 'r')
    w_pipe = os.fdopen(w_pipe_fd, 'w')

    logger = logging.getLogger()
    _, filename, lineno, func_name, _, _ = inspect_stack()[1]

    def read_pipe():
        for l in r_pipe:
            if len(l) > 0:
                if l[-1] == '\n':
                    l = l[:-1]
                if logger.isEnabledFor(loglevel):
                    logger.handle(
                        logger.makeRecord(
                            logger.name,
                            loglevel,
                            filename,
                            lineno,
                            l,
                            (),
                            None,
                            func_name
                        )
                    )

    thread = Thread(target=read_pipe)
    thread.daemon = True
    thread.start()

    return w_pipe


stream_handler = logging.StreamHandler()
stream_handler.setLevel(logging.DEBUG)

journald_handler = JournalHandler()
journald_handler.setLevel(logging.DEBUG)
journald_handler.setFormatter(logging.Formatter(
    fmt='%(levelname)s %(module)s:%(funcName)s:%(lineno)d: %(message)s'))
logging.basicConfig(
    level=logging.DEBUG,
    format='%(asctime)s %(levelname)s %(module)s:%(funcName)s:%(lineno)d: %(message)s',
    handlers=(stream_handler, journald_handler, )
)


def handle_exception(exc_type, exc_value, exc_traceback):
    logging.error("Unhandled exception", exc_info=(exc_type, exc_value, exc_traceback))
    sys.excepthook = handle_exception


if sys.stdin.isatty():
    logging.debug('launched from interactive terminal')
else:
    logging.debug('launched')

logging.info('cmdline: %s', str(sys.argv))
logging.info('environ: %s', str(os.environ))


i3_pid = int(check_output(['pgrep', '-x', 'i3'], stderr=logfile(logging.DEBUG)))
i3_env = dict(map(methodcaller('split', '=', 1), filter(
    None, open('/proc/{0}/environ'.format(i3_pid)).read().split('\0'))))

udev_drm_change = os.environ.get('ACTION') == 'change' and os.environ.get('SUBSYSTEM') == 'drm'
# if udev_drm_change:
#     sleep(1)

lid_open = 'open' in open('/proc/acpi/button/lid/LID0/state').read()
charger_connected = int(open('/sys/class/power_supply/ADP1/online').read()) == 1
battery = float(open('/sys/class/power_supply/BAT0/charge_now').read()) / \
    float(open('/sys/class/power_supply/BAT0/charge_full').read())
displays = dict(map(lambda x: (x[0], x[1] == 'connected'), filter(
    lambda x: x[1] in ('connected', 'disconnected'), map(
        methodcaller('split', ' '),
        check_output(['xrandr'],
                     env=i3_env, stderr=logfile(logging.DEBUG)).decode().splitlines()))))

if __name__ == '__main__':
    logging.debug('lid open: %s', lid_open)
    logging.debug('charger connected: %s', charger_connected)
    logging.debug('battery: %s', battery)
    logging.debug('drm change: %s', udev_drm_change)

    # for output, connected in displays.items():
    #     if output in ('eDP-1', 'eDP1',):
    #         connected = lid_open
    #         call(['xrandr', '--output', output, '--auto' if connected else '--off'],
    #              env=i3_env, stdout=logfile(logging.DEBUG), stderr=logfile(logging.DEBUG))
    #         logging.debug('%s %s', output, connected)
    # TODO: do this only if display configuration changed
    # fixes tray rendering
    # call(['i3-msg', 'restart'], env=i3_env, stdout=logfile(logging.DEBUG), stderr=logfile(logging.DEBUG))

logging.info('finished my power stuff')
