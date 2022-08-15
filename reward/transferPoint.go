package reward

import (
	"lycle/util"

	"github.com/go-playground/validator"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type TransferPointRequest struct {
	Sender    string `validate:"required,email"`
	Recipient string `validate:"required,email"`
	Point     int	 `validate:"required,gt=0"`
}

type TransferPointResponse struct {
	Sender		User
	Recipient	User
}

func (req TransferPointRequest) validate() (err error) {
	v := validator.New()
	err = v.Struct(req)
	return
}

func TransferPoint(APIstub shim.ChaincodeStubInterface, request util.Request) (response util.Response, err error) {
	var req TransferPointRequest
	var res TransferPointResponse

	// unmarshalling request 
	err = request.Parse(&req)
	if err != nil {
		return
	}

	// validate request
	err = req.validate()
	if err != nil {
		err = util.ErrInvalidParam
		return
	}

	// validate sender in blockchain
	var Sender User
	Sen := req.Sender
	err = util.GetState(APIstub, Sen, &Sender)
	if err != nil {
		return
	}

	if req.Point > Sender.Point {
		err = util.ErrNotEnoughPoint
		return
	}

	// validate recipient in blockchain
	var Recipient User
	Rec := req.Recipient
	err = util.GetState(APIstub, Rec, &Recipient)
	if err != nil {
		return
	}

	Sender.Point -= req.Point
	Recipient.Point += req.Point

	Sender.UpdatedAt = util.GetTimestamp()
	Recipient.UpdatedAt = Sender.UpdatedAt

	err = util.PutState(APIstub, Sen, Sender)
	if err != nil {
		return
	}
	err = util.PutState(APIstub, Rec, Recipient)
	if err != nil {
		return
	}

	res.Sender = Sender
	res.Recipient = Recipient

	response.SetData(res)

	return
}
