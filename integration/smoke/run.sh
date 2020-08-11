#!/usr/bin/env bash

bash ./integration/utils/start-server.sh ./integration/smoke/api.yml
# shellcheck disable=SC2046
if [ $(curl http://localhost:8000/ping) == "pong" ]; then
  exit 0
else
  exit 1
fi

