pragma solidity ^0.4.11;
import "./Address.sol";

library ValidateUtil {
    function addressMatchPublicKey(address addr, bytes32[2] pub) internal constant returns (bool) {
        return (Address.getAddressFromPublicKeyV2(pub) == addr);
    }
    function nameHashMatchName(bytes32 hash, bytes32[4] name) internal constant returns (bool) {
        return (keccak256(name) == hash);
    }
    function publicKeyNotZero(bytes32[2] publicKey) internal constant returns (bool) {
        return (publicKey[0]!=0x0 && publicKey[1]!=0x0);
    }
    function stringNotEmpty(string str) internal constant returns (bool) {
        bytes memory temp = bytes(str);
        return(temp.length != 0x0);
    }
}
