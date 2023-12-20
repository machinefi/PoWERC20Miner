## Quick Start

1. **mine the default token**:
   ```bash
   cd cmd/miner
   go run main.go --privateKey=your private key
   ```
the dependent cmd's source code is https://github.com/machinefi/rust-poseidon/tree/main/gen-poseidon-pse

## Advancement

1. **create new token**:
   ```bash
   cd cmd/factory
   go run main.go --privateKey=your private key
   ```

2. **mine the default token**:
   ```bash
   cd cmd/miner
   go run main.go --privateKey=your private key --contractAddress=the contract created above
   ```
  
