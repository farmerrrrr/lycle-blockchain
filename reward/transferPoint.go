package reward

import (
	"strconv"
	"lycle/util"

	"github.com/go-playground/validator"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type TransferPointRequest struct {
	sender		string	`validate:"required,email"`
	recipient	string	`validate:"required,email"`
	point		int		`validate:"required,gt=0"`
}

type TransferPointResponse struct {
	sender		User
	recipient	User
}

func createTransferPointRequest(args []string) (req TransferPointRequest) {
	req.sender = args[0]
	req.recipient = args[1]
	req.point, _ = strconv.Atoi(args[2])
	return
}

func (req TransferPointRequest) validate() (err error) {
	v := validator.New()
	err = v.Struct(req)
	return
}

func TransferPoint(APIstub shim.ChaincodeStubInterface, request util.Request) (response util.Response, err error) {
	var req TransferPointRequest
	var res TransferPointResponse

	// validate request
	req = createTransferPointRequest(request.GetArguments())
	err = req.validate()
	if err != nil {
		err = util.ErrInvalidParam
		return
	}

	// validate sender in blockchain
	var sender User
	sen := req.sender
	err = util.GetState(APIstub, sen, &sender)
	if err != nil {
		return
	}

	if req.point > sender.point {
		err = util.ErrNotEnoughPoint
		return
	}

	// validate recipient in blockchain
	var recipient User
	rec := req.recipient
	err = util.GetState(APIstub, rec, &recipient)
	if err != nil {
		return
	}

	sender.point -= req.point
	recipient.point += req.point

	sender.updateAt = util.GetTimestamp()
	recipient.updateAt = sender.updateAt

	err = util.PutState(APIstub, sen, sender)
	if err != nil {
		return
	}
	err = util.PutState(APIstub, rec, recipient)
	if err != nil {
		return
	}

	res.sender = sender
	res.recipient = recipient

	response.SetData(res)

	return
}
