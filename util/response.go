package util

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Response struct {
	statusMessage string
	timestamp     string
	data          interface{}
}

func (res *Response) SetTimestamp(timestamp string) {
	res.timestamp = timestamp
}

func (res Response) GetStatusMessage() string {
	return res.statusMessage
}

func (res *Response) SetStatusMessage(statusMessage string) {
	res.statusMessage = statusMessage
}

func (res Response) GetData() interface{} {
	return res.data
}

func (res Response) SetData(data interface{}) {
	res.data = data
}

func GenerateResponse(res Response, err error) peer.Response {
	res.SetTimestamp(GetTimestamp())

	if err != nil {
		res.SetStatusMessage(err.Error())
	} else {
		res.SetStatusMessage(SuccessMessage)
	}

	resAsBytes, _ := json.Marshal(res)

	if err != nil {
		return shim.Error(string(resAsBytes))
	}

	return shim.Success(resAsBytes)
}
