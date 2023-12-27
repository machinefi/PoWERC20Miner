// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

interface IDepinRC20 {
    function mine(uint256 nonce, address sender, bytes calldata proof) external;
}

contract DepinRCRouterContract {
    IDepinRC20 public depinRC20;

    function mine(address depinRC20Address, uint256 nonce, address sender, bytes calldata proof) public {
        depinRC20 = IDepinRC20(depinRC20Address);
        depinRC20.mine(nonce, sender, proof);
    }
}
