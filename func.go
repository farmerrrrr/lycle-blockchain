package main

import (
	"encoding/json"
	f "lycle/reward"
	"lycle/util"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

func (s *SmartContract) Do(APIstub shim.ChaincodeStubInterface) (response peer.Response) {
	req := util.GenerateRequest(APIstub)

	var res util.Response
	var err error

	switch req.GetFunction() {
	case util.RegisterUser:
		res, err = f.RegisterUser(APIstub, req)
	case util.GetPoint:
		res, err = f.GetPoint(APIstub, req)
	case util.TransferPoint:
		res, err = TransferPoint(APIstub, req)
	default:
		err = util.ErrInvalidFunc
	}

	util.GenerateResponse(&res, err)

	resAsBytes, _ := json.Marshal(res)

	if err != nil {
		return shim.Error(string(resAsBytes))
	}

	shim.Success(resAsBytes)
}
