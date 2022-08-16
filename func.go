package main

import (
	"fmt"
	f "lycle/reward"
	"lycle/util"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

func (s *SmartContract) Do(APIstub shim.ChaincodeStubInterface) (peer.Response) {
	fmt.Println("** Transaction ID: " + APIstub.GetTxID())

	req := util.GenerateRequest(APIstub)

	fmt.Print("function name: " + req.GetFunction() + ", arguments: ")
	fmt.Println(string(*req.GetArguments()))

	var res util.Response
	var err error

	switch req.GetFunction() {
	case util.RegisterUser:
		res, err = f.RegisterUser(APIstub, req)
	case util.GetPoint:
		res, err = f.GetPoint(APIstub, req)
	case util.TransferPoint:
		res, err = f.TransferPoint(APIstub, req)
	case util.DepositPoint:
		res, err = f.DepositPoint(APIstub, req)
	default:
		err = util.ErrInvalidFunc
	}

	return util.GenerateResponse(res, err)
}
