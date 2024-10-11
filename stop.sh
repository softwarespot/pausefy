#!/bin/bash

cd "$(dirname "$0")" || exit 1

PID_FILE=./pid
NOHUP_FILE=./nohup.out

if [ -f "$PID_FILE" ]; then
    pkill -F "$PID_FILE"
    echo "Stopped PID: $(< "$PID_FILE")"
    rm "$PID_FILE"
fi

# Ignore if the file doesn't exist
rm --force "$NOHUP_FILE"
