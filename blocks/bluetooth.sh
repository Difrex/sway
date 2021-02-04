#!/usr/bin/env bash

# Origin https://github.com/cwrau/linux-config/blob/master/.config/polybar/scripts/bluetooth.sh

function dbus() {
  mac=$2
  case "$1" in
    bluez)
      service=bluez
      path=dev_${mac//:/_}
      object=Battery1
      field=Percentage
      ;;
    hsphfpd)
      service=hsphfpd
      path=dev_${mac//:/_}/$3
      object=Endpoint
      field=BatteryLevel
      ;;
  esac

  dbus-send --system --dest="org.$service" --print-reply=literal /org/"$service"/hci0/"$path" org.freedesktop.DBus.Properties.Get "string:org.$service.$object" "string:$field" 2>/dev/null | awk '{print $3}' || echo -1
}

if bluetoothctl show | grep -q "Powered: yes"; then
  if bluetoothctl info | grep -q 'Device'; then
    for mac in $(bluetoothctl paired-devices | awk '{print $2}'); do
      info="$(bluetoothctl info "$mac")"
      if grep -q 'Connected: yes' <<<"$info"; then
        battery=
        alias="$(grep 'Alias' <<<"$info" | cut -d ' ' -f 2-)"
        if grep -q 'Battery Service' <<<"$info"; then
          battery=$(dbus bluez "$mac")
        else
            continue
        fi
        echo "$alias${battery:+ ($battery%)}"
      fi
    done | sort | paste -sd "," - | sed 's#,#, #g; s#^# #'
  fi
else
  echo ""
fi
