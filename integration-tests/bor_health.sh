#!/bin/bash
set -e

while true
do
    peers=$(docker exec bor0 bash -c "bor attach /root/.bor/data/kernel.ipc -exec 'admin.peers'")
    block=$(docker exec bor0 bash -c "bor attach /root/.bor/data/kernel.ipc -exec 'eth.blockNumber'")

    if [[ -n "$peers" ]] && [[ -n "$block" ]]; then
        break
    fi
done

echo $peers
echo $block
