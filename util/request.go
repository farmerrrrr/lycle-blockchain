package util

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Request struct {
	function  string
	arguments *[]byte
}

func (req Request) GetFunction() string {
	return req.function
}

func (req *Request) SetFunction(function string) {
	req.function = function
}

func (req Request) GetArguments() *[]byte {
	return req.arguments
}

func (req *Request) SetArguments(args *[]byte) {
	req.arguments = args
}

func GenerateRequest(APIstub shim.ChaincodeStubInterface) (req Request) {
	args := APIstub.GetArgs()
	function := string(args[FunctionName])
	parameter := &(args[Parameter])

	req.SetFunction(function)
	req.SetArguments(parameter)

	return
}

func (req Request) Parse(apiReq interface{}) (err error) {
	err = json.Unmarshal(*req.arguments, apiReq)
	if err != nil {
		err = ErrMarshalFailed
		return
	}
	return
}