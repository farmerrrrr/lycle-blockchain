package util

import (
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func PutState(APIstub shim.ChaincodeStubInterface, key string, object interface{}) (err error) {
	// json -> bytes
	objectAsBytes, err := json.Marshal(object)
	if err != nil {
		return ErrMarshalFailed
	}

	err = APIstub.PutState(key, objectAsBytes)
	if err != nil {
		fmt.Println("PutState() failed. key: " + key)
		return ErrPutStateFailed
	}

	fmt.Println("PutState() success. key: " + key)

	return
}
