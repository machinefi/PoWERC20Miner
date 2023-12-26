1. **create new token**:
   ```bash
   cd cmd/factory
   go run main.go --privateKey=your private key
   ```
 will get a erc20 contract address

2. **mine the token**:
   ```bash
   cd cmd/miner
   go run main.go --privateKey=your private key --contractAddress=the contract created above
   ```