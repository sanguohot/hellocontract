pragma solidity ^0.4.11;

contract EasyCns {
    address owner;
    mapping(string => address) nameToAddressMap;
    function EasyCns() public {
        owner = msg.sender;
    }
    function get(string name) constant public returns(address) {
        return nameToAddressMap[name];
    }
}
