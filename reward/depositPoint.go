package reward

import (
	"lycle/util"

	"github.com/go-playground/validator"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type DepositPointRequest struct {
	Email string `validate:"required,email"`
	Point int	 `validate:"required,gt=0"`
}

type DepositPointResponse struct {
	User User
}

func (req DepositPointRequest) validate() (err error) {
	v := validator.New()
	err = v.Struct(req)
	return
}

func DepositPoint(APIstub shim.ChaincodeStubInterface, request util.Request) (response util.Response, err error) {
	var req DepositPointRequest
	var res DepositPointResponse

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

	// validate user in blockchain
	var Usr User
	key := req.Email
	err = util.GetState(APIstub, key, &Usr)
	if err != nil {
		return
	}

	Usr.Point += req.Point
	Usr.UpdatedAt = util.GetTimestamp()

	err = util.PutState(APIstub, key, Usr)
	if err != nil {
		return
	}

	res.User = Usr

	response.SetData(res)

	return
}
