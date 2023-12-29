#!/bin/bash

privateKey=""
contractAddress=""
workerCount=10

while getopts "k:a:c:" opt; do
  case $opt in
    k) privateKey=$OPTARG ;;
    a) contractAddress=$OPTARG ;;
    c) workerCount=$OPTARG ;;
    \\?) echo "Invalid option -$OPTARG" >&2 ;;
  esac
done



if [ ! -f "./bin/miner" ]; then
    echo "Build DePinRC-20 file..."
    make build
fi

if ! command -v ioctl &> /dev/null; then
    echo "Install ioctl command..."
    brew tap iotexproject/ioctl-unstable
    brew install iotexproject/ioctl-unstable/ioctl-unstable
    alias ioctl=`which ioctl-unstable`
fi

if ! command -v jq &> /dev/null
then
    echo "Install jq command..."
    brew install jq
fi

echo "Run the miner command to mine tokens..."
./bin/miner --privateKey=$privateKey --contractAddress=$contractAddress --workerCount=$workerCount > miner.log 2>&1

send_w3bstream_message=$(grep 'Use this cmd to submit nonce' miner.log | jq -r '.msg' | grep -o 'Use this cmd to submit nonce: .*' | cut -d: -f2- | sed 's/^ *//')

if [ -z "$send_w3bstream_message" ]
then
  echo "No w3bstream message command found. See miner.log for details"
  exit 1
fi

echo "set w3bstream endpoint to devnet sprout-staging.w3bstream.com:9000"
ioctl config set wsEndpoint sprout-staging.w3bstream.com:9000

echo "Send a message to w3bstream to generate and verify a zero-knowledge proof."
eval "$send_w3bstream_message" > miner.log 2>&1

messageID=$(cat miner.log | jq -r '.messageID')
echo "messageID is $messageID"

count=0
while true; do
    outputted=$(curl -s "sprout-staging.w3bstream.com:9000/message/$messageID" | jq -r '.states[] | select(.state == "outputted")')

    if [ -n "$outputted" ]; then
        comment=$(echo $outputted | jq -r '.comment')
        echo $comment
        break
    else
        count=$((count+1))
        if [ $count -gt 5 ]; then
            echo "After 5 attempts, the outputted status information is still not available, now exiting."
            break
        fi
        echo "The current outputted status information is not available, try again after 30 seconds."
        sleep 30
    fi
done
