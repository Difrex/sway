
#!/bin/bash

MAILDIR=/home/difrex/Mail; export MAILDIR
COUNT="$(/usr/bin/mu find date:15m..now and flag:unread 2>/dev/null | wc -l)"

# Write notify lock
write_notify_lock() {
    mkdir -p ~/.cache
    echo $COUNT > ~/.cache/.mail_notify_lock
}

# Cleanup lock
clean_notify_lock() {
    rm -f ~/.cache/.mail_notify_lock
}

if [[ $COUNT -gt 0 ]]; then
    if [[ ! -f ~/.cache/.mail_notify_lock ]] || [[ $COUNT -gt $(/usr/bin/cat ~/.cache/.mail_notify_lock) ]]; then
        write_notify_lock
        /usr/bin/notify-send 'New mail' "$(/usr/bin/mu find date:15m..now and flag:unread 2>/dev/null | tail -1)" --icon=/usr/share/icons/oxygen/base/32x32/status/mail-unread-new.png
    fi
else
    if [[ -f ~/.cache/.mail_notify_lock ]]; then
        clean_notify_lock
    fi
fi

echo $COUNT
