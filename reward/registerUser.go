package reward

import (
	"lycle/util"

	"github.com/go-playground/validator"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const defaultPoint int = 0

type RegisterUserRequest struct {
	Email string `validate:"required,email"`
}

type RegisterUserResponse struct {
	User User
}

func (req RegisterUserRequest) validate() (err error) {
	v := validator.New()
	err = v.Struct(req)
	return
}

func RegisterUser(APIstub shim.ChaincodeStubInterface, request util.Request) (response util.Response, err error) {
	var req RegisterUserRequest
	var res RegisterUserResponse

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

	// check to user in blockchain
	var Usr User
	key := req.Email
	err = util.GetState(APIstub, key, &Usr)
	if err != nil {
		if err != util.ErrNoDataFound { // no data found면, continue
			return
		}
	} else {
		err = util.ErrAlreadyExists
		return
	}

	Usr.Email = req.Email
	Usr.Point = defaultPoint
	Usr.RegisteredAt = util.GetTimestamp() // getTimestamp 에러 처리 필요
	Usr.UpdatedAt = Usr.RegisteredAt

	err = util.PutState(APIstub, key, Usr)
	if err != nil {
		return
	}

	res.User = Usr

	response.SetData(res)

	return
}
