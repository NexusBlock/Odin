#!/bin/bash
# This is a postinstallation script so the service can be configured and started when requested
#
sudo -u heimdall odind init --chain=loki --home /var/lib/heimdall
sudo systemctl daemon-reload