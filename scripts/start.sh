#!/bin/bash

cd "$(dirname "$0")" || exit 1

EXE=../bin/pausefy
PID_FILE=./pid

if [ ! -f "$EXE" ]; then
    echo -e "\033[0;31mMissing \"$EXE\". See README.md on how to build/deploy\033[0m"
    exit 2
fi

nohup "$EXE" & echo $! > "$PID_FILE"
echo "Running with PID: $!"