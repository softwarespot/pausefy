#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")" || exit 1

PID_FILE=./pid
NOHUP_FILE=./nohup.out

if [ -f "$PID_FILE" ]; then
    # Try to kill by PID file; ignore any error e.g. no matching process
    pkill -F "$PID_FILE" >/dev/null 2>&1 || true
    echo "Stopped PID: $(< "$PID_FILE")"
    rm "$PID_FILE"
fi

# Ignore if the file doesn't exist
rm -f "$NOHUP_FILE"
