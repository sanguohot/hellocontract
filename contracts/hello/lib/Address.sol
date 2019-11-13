pragma solidity ^0.4.11;

library Address {
    function getAddressFromPublicKeyV1(bytes _publicKey) internal constant returns (address) {
        return address(keccak256(_publicKey));
    }
    function getAddressFromPublicKeyV2(bytes32[2] _publicKey) internal constant returns (address) {
        return address(keccak256(_publicKey));
    }
    /**
    * Returns whether the target address is a contract
    * @dev This function will return false if invoked during the constructor of a contract,
    * as the code is not actually created until after the constructor finishes.
    * @param _account address of the account to check
    * @return whether the target address is a contract
    */
    function isContract(address _account) internal returns (bool) {
        require(_account != 0x0);
        uint256 size;
        // XXX Currently there is no better way to check if there is a contract in an address
        // than to check the size of the code at that address.
        // See https://ethereum.stackexchange.com/a/14016/36603
        // for more details about how this works.
        // TODO Check this again before the Serenity release, because all addresses will be
        // contracts then.
        // solium-disable-next-line security/no-inline-assembly
        assembly { size := extcodesize(_account) }
        return size > 0;
    }
}
