pragma solidity ^0.4.11;

library Const {
    string constant usersDataName = "UsersData";
    string constant orgsDataName = "OrgsData";
    string constant filesDataName = "FilesData";

    function getUsersDataName() internal constant returns (string) {
        return usersDataName;
    }
    function getOrgsDataName() internal constant returns (string) {
        return orgsDataName;
    }
    function getFilesDataName() internal constant returns (string) {
        return filesDataName;
    }
}
