package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
//	mspid, _ := cid.GetMSPID(APIstub)
//	fmt.Printf("lycle initialized. tx:%v, by:%v", APIstub.GetTxID(), mspid)

	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	return s.Do(APIstub)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error loading lycle: %s", err)
	}
}
