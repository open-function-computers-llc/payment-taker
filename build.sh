#! /bin/bash

echo "Killing any old processes"
killall payment-osx > /dev/null 2>&1
killall payment-linux > /dev/null 2>&1

echo "Purging old binaries"
rm dist/*

echo "Building new binaries"
cd httpd
GOOS=darwin GOARCH=amd64 go build -o ../dist/payment-osx
GOOS=linux GOARCH=amd64 go build -o ../dist/payment-linux
cd ..

echo "Starting new built binary"
if [[ "$OSTYPE" == "linux-gnu" ]]; then
    dist/payment-linux &
elif [[ "$OSTYPE" == "darwin"* ]]; then
    dist/payment-osx &
else
    echo "i have no idea what's going on"
fi
