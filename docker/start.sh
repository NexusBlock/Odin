#!/usr/bin/env sh

# start processes
odind start > ./logs/odind.log &
odind rest-server > ./logs/odind-rest-server.log &
sleep 100
bridge start --all > ./logs/bridge.log &

# tail logs
tail -f ./logs/odind.log ./logs/odind-rest-server.log ./logs/bridge.log
