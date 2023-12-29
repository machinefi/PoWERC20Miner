## Start with script
1. **create your new token**:
```shell
sh factory.sh -k $your_private_key
```
will get a depinrc-20 contract address

2. **mine the token**:
```shell
sh miner.sh -k $your_private_key -a $depinrc-20_contract_address
```
will get a transaction address

## Run with source code
1. **clone the repo**:
    ```bash
   git clone https://github.com/machinefi/PoWERC20Miner.git
    ```

2. **create your new token**:
   ```bash
   cd cmd/factory
   go run main.go --privateKey=your private key
   ```
 will get a depinrc-20 contract address

3. **mine the token**:
   ```bash
   cd cmd/miner
   go run main.go --privateKey=your private key --contractAddress=the contract created above
   ```
will get a ioctl command

4. **send message to w3bstream to generate and verify a zero-knowledge proof**
copy the command from the previous step and execute. like this:

```shell
ioctl ws message send --project-id 20000 --project-version "0.1" --data "{...}"
```
