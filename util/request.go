package util

import "github.com/hyperledger/fabric/core/chaincode/shim"

type Request struct {
	function  string
	arguments []string
}

func (req Request) GetFunction() string {
	return req.function
}

func (req *Request) SetFunction(function string) {
	req.function = function
}

func (req Request) GetArguments() []string {
	return req.arguments
}

func (req *Request) SetArguments(args []string) {
	req.arguments = args
}

func GenerateRequest(APIstub shim.ChaincodeStubInterface) (req Request) {
	function, args := APIstub.GetFunctionAndParameters()

	req.SetFunction(function)
	req.SetArguments(args)

	return
}
