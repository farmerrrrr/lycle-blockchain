package util

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetState(APIstub shim.ChaincodeStubInterface, key string, object interface{}) (err error) {
	objectAsBytes, err := APIstub.GetState(key)
	if objectAsBytes == nil {
		return ErrNoDataFound
	}
	if err != nil {
		return ErrGetStateFalied
	}

	// bytes -> json
	err = json.Unmarshal(objectAsBytes, object)
	if err != nil {
		return ErrMarshalFailed
	}

	return
}
