package util

import (
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetState(APIstub shim.ChaincodeStubInterface, key string, object interface{}) (err error) {
	objectAsBytes, err := APIstub.GetState(key)
	if objectAsBytes == nil {
		fmt.Println("GetState() failed. key: " + key)
		return ErrNoDataFound
	}
	if err != nil {
		fmt.Println("GetState() failed. key: " + key)
		return ErrGetStateFalied
	}

	fmt.Println("GetState() success. key: " + key)

	// bytes -> json
	err = json.Unmarshal(objectAsBytes, object)
	if err != nil {
		return ErrMarshalFailed
	}

	return
}
