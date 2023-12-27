// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

interface IDepinRC20 {
    function mine(uint256 nonce, address sender, bytes calldata proof) external;
}

contract DepinRCRouterContract {
    IDepinRC20 public depinRC20;

    function mine(address _depinRC20, uint256 nonce, address sender, bytes calldata proof) public {
        depinRC20 = IDepinRC20(_depinRC20);
        depinRC20.mine(nonce, sender, proof);
    }
}
