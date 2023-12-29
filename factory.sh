#!/bin/bash

name="DePinRC20-test"
symbol="DePinRC20-test"
totalSupply=1000000000
decimals=18
difficulty=50
mintLimitPerAddress=200000000
limitPerMint=1000
privateKey=""
factoryContractAddress="0x2a96010335A9A35ca29CA47D35ed8730cC839b39"
verifierContractAddress="0xb9c809BA641EFdAcd6ca34515F87affd1d8Cc1a5"

while getopts "n:s:t:d:D:a:l:k:f:v:" opt; do
  case $opt in
    n) name=$OPTARG ;;
    s) symbol=$OPTARG ;;
    t) totalSupply=$OPTARG ;;
    d) decimals=$OPTARG ;;
    D) difficulty=$OPTARG ;;
    a) mintLimitPerAddress=$OPTARG ;;
    l) limitPerMint=$OPTARG ;;
    k) privateKey=$OPTARG ;;
    f) factoryContractAddress=$OPTARG ;;
    v) verifierContractAddress=$OPTARG ;;
    \\?) echo "Invalid option -$OPTARG" >&2 ;;
  esac
done

if [ ! -f "./bin/factory" ]; then
    echo "Build DePinRC-20 file..."
    make build
fi

echo "Run the factory command to generate DePinRC20 token contract..."
./bin/factory --name=$name --symbol=$symbol --totalSupply=$totalSupply --decimals=$decimals \
      --difficulty=$difficulty --mintLimitPerAddress=$mintLimitPerAddress --limitPerMint=$limitPerMint \
      --privateKey=$privateKey --factoryContractAddress=$factoryContractAddress --verifierContractAddress=$verifierContractAddress  > factory.log 2>&1

contract_address=$(grep 'New DePinRC20 contract' factory.log | awk '{print $NF}' | tr -d '"')

if [ -z "$contract_address" ]
then
  echo "No contract address found. See factory.log for details"
  exit 1
else
  echo "DePinRC20 contract address: $contract_address"
fi
