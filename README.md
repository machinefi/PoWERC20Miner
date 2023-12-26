1. **clone the repo**:
    ```bash
   git clone https://github.com/machinefi/PoWERC20Miner.git
    ```

2. **create your new token**:
   ```bash
   cd cmd/factory
   go run main.go --privateKey=your private key
   ```
 will get a erc20 contract address

3. **mine the token**:
   ```bash
   cd cmd/miner
   go run main.go --privateKey=your private key --contractAddress=the contract created above
   ```