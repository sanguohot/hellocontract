pragma solidity ^0.4.11;

contract Hello{
    string saySomething;

    event onSaySomethingElse(string newSaying);

    function Hello() public  {
        saySomething = "Hello World!";
    }

    function speak() public constant returns(string itSays) {
        return saySomething;
    }

    function saySomethingElse(string newSaying) public returns(bool success) {
        saySomething = newSaying;
        onSaySomethingElse(newSaying);
        return true;
    }
}