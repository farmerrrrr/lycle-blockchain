package reward

import (
	"lycle/util"

	"github.com/go-playground/validator"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type GetPointRequest struct {
	Email string `validate:"required,email"`
}

type GetPointResponse struct {
	User User
}

func (req GetPointRequest) validate() (err error) {
	v := validator.New()
	err = v.Struct(req)
	return
}

func GetPoint(APIstub shim.ChaincodeStubInterface, request util.Request) (response util.Response, err error) {
	var req GetPointRequest
	var res GetPointResponse

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

	Usr := User{}
	key := req.Email
	err = util.GetState(APIstub, key, &Usr)
	if err != nil {
		return
	}

	res.User = Usr

	response.SetData(res)

	return
}