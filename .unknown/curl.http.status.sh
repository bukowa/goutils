#!/usr/bin/env bash
set -e
$(go run server.go --host=testxx --port=9090) &
sleep 1

if [ "$(curl -s -o /dev/null -w "%{http_code}" testxx:9090)" == 404 ]; then
  echo "Success!"
  exit 0
fi

exit 1
