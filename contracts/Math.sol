pragma solidity 0.4.23;

/**
 * @title Math
 * @dev Math operations with safety checks that throw on error
 */

library Math {

    function max(uint256 a, uint256 b)
        internal
        pure
        returns (uint256) 
    {
        return a >= b ? a : b;
    }
}