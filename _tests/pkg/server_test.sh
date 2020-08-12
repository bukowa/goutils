#!/usr/bin/env bash
$(go run ./_tests/pkg/server.go --host=testxx --port=9090) &
sleep 1

if [ "$(curl -s -o /dev/null -w "%{http_code}" testxx:9090)" != 404 ]; then
    echo "Fails..."
    exit 1
fi

echo "Success!"
exit 0
