package util

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Response struct {
	StatusMessage string
	Timestamp     string
	Data          interface{}
}

func (res *Response) SetTimestamp(timestamp string) {
	res.Timestamp = timestamp
}

func (res Response) GetStatusMessage() string {
	return res.StatusMessage
}

func (res *Response) SetStatusMessage(statusMessage string) {
	res.StatusMessage = statusMessage
}

func (res Response) GetData() interface{} {
	return res.Data
}

func (res *Response) SetData(data interface{}) {
	res.Data = data
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
