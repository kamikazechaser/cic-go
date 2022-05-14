// SPDX-License-Identifier: MIT

// Adapted from https://github.com/MyCryptoHQ/eth-scan

pragma solidity >= 0.8.0;

contract Balances {
    struct Result {
        bool success;
        uint256 balance;
    }

    function tokensBalance(address owner, address[] calldata contracts) external view returns (Result[] memory results) {
        results = new Result[](contracts.length);

        bytes memory data = abi.encodeWithSignature("balanceOf(address)", owner);

        for (uint256 i = 0; i < contracts.length; i++) {
            results[i] = staticCall(contracts[i], data, 8000000);
        }
    }

    function staticCall(
        address target,
        bytes memory data,
        uint256 gas
    ) private view returns (Result memory) {
        uint256 size = codeSize(target);

        if (size > 0) {
            (bool success, bytes memory result) = target.staticcall{ gas: gas }(data);
            if (success) {
                uint256 balance = abi.decode(result, (uint256));
                return Result(success, balance);
            }
        }

        return Result(false, 0);
    }

    function codeSize(address _address) private view returns (uint256 size) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
         size := extcodesize(_address)
      }
    }
}

