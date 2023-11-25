#!/usr/bin/env sh

if [ "$1" = 'odincli' ]; then
    shift
    exec odincli --home=$HEIMDALL_DIR "$@"
fi

exec odind --home=$HEIMDALL_DIR "$@"
