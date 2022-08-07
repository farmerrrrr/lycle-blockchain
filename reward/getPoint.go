package reward

import (
	"lycle/util"

	"github.com/go-playground/validator"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type GetPointRequest struct {
	email string `validate:"required,email"`
}

type GetPointResponse struct {
	user User
}

func createGetPointRequest(args []string) (req GetPointRequest) {
	req.email = args[0]
	return
}

func (req GetPointRequest) validate() (err error) {
	v := validator.New()
	err = v.Struct(req)
	return
}

func GetPoint(APIstub shim.ChaincodeStubInterface, request util.Request) (response util.Response, err error) {
	var req GetPointRequest
	var res GetPointResponse

	// validate request
	req = createGetPointRequest(request.GetArguments())
	err = req.validate()
	if err != nil {
		err = util.ErrInvalidParam
		return
	}

	var user User
	key := req.email
	err = util.GetState(APIstub, key, &user)
	if err != nil {
		return
	}

	res.user = user

	response.SetData(res)
}
