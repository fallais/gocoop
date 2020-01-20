#!/bin/bash

# Start Caddy
/usr/local/bin/caddy --conf /etc/Caddyfile &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start caddy: $status"
  exit $status
fi

# Start GoCoop
/usr/bin/gocoop --config_file /usr/bin/config.yml
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start gocoop: $status"
  exit $status
fi
