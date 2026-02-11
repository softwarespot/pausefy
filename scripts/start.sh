#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")" || exit 1

EXE=../bin/pausefy
PID_FILE=./pid

if [ ! -f "$EXE" ]; then
    echo -e "\033[0;31mMissing \"$EXE\". See README.md on how to build/deploy\033[0m"
    exit 2
fi

# Start the executable in the background and store the process ID to the PID file.
# Use "nohup" to ensure it keeps running after the terminal has closed
nohup "$EXE" &
PID=$!
echo "$PID" > "$PID_FILE"
echo "Running with PID: $PID"
