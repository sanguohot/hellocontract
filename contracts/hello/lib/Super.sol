pragma solidity ^0.4.11;

contract Super {
    address owner;
    address[] supers;
    mapping(address => bool) superm;
    
    event onAddSuper(address addr);
    event onDelSuper(address addr);
    
    function Super() public {
        owner = msg.sender;
    }
    function isSuper(address addr) internal constant returns (bool valid) {
        return addr!=0x0 && superm[addr];
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    modifier onlySuper() {
        require(isSuper(msg.sender));
        _;
    }
    modifier onlySuperOrOwner() {
        require(msg.sender==owner || isSuper(msg.sender));
        _;
    }
    function notOutOfIndex(uint256 index) internal constant returns (bool valid) {
        return (index>=0 && index<supers.length);
    }

    // this should be done first if it is called by other contracts.
    function addSuper(address addr) public onlyOwner {
        require(addr != 0x0);
        supers.push(addr);
        superm[addr] = true;
        onAddSuper(addr);
    }
    function delSuper(address addr) public onlyOwner {
        require(addr != 0x0);
        require(isSuper(addr));
        bool found = false;
        for (uint i=0; i<supers.length; i++) {
            if (addr == supers[i]){
                found = true;
                break;
            }
        }
        assert(found);
        delete supers[i];
        delete superm[addr];
        onDelSuper(addr);
    }
    function getSuperSize() public constant returns (uint256) {
        return supers.length;
    }
    function getSuperByIndex(uint256 index) public constant returns (address) {
        require(notOutOfIndex(index));
        return supers[index];
    }
}
