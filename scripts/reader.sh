#!/bin/sh

while true; do
  if [ -f /data/variable.txt ]; then
    value=$(cat /data/variable.txt)
    if [ "$value" -eq "1" ]; then
      sleep 15
      k6 run /load_test.js
      break
    fi
  fi
  sleep 1
done
