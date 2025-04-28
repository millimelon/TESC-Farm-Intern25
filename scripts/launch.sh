#!/bin/bash

cd api
go run cmd/farmapi.go &
sleep 2
cd ..

echo 'Please select from the site list:'
nl ./scripts/data/site.list
count=$(wc -l ./scripts/data/site.list | cut -d '.' -f1)
count="${count//[$'\t\r\n ']}"
n=""
while true; do
    read -p 'Select option: ' n
    n="${n//[$'\t\r\n ']}"
    echo $n
    if [ "$n" -gt 0 ] && [ "$n" -le "$count" ]; then
        break
    fi
done
value="$(sed -n "${n}p" ./scripts/data/site.list)"
echo "Selected site $n: '$value'"

cd $value
yarn vite &
sleep 1
cd ..

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
