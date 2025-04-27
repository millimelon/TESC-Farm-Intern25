#!/bin/bash

cd api
go run cmd/farmapi.go &
cd ../taskpanel
vite &

close() {
  echo "Exiting on Ctrl+C"
  pkill -P $$
  exit 0
}
trap close SIGINT

echo "Development environment launched. Press Ctrl+C to exit."

while true; do
    sleep 1
done
