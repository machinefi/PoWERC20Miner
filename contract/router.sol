
function mine(uint256 nonce, address sender, bytes calldata proof, contractAddress) public {
       contractAddress.mine(nonce,sender,proof)
    }