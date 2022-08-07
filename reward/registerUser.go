package reward

import (
	"lycle/util"

	"github.com/go-playground/validator"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const defaultPoint int = 0

type RegisterUserRequest struct {
	email string `validate:"required,email"`
}

type RegisterUserResponse struct {
	user User
}

func createRegisterUserRequest(args []string) (req RegisterUserRequest) {
	req.email = args[0]
	return
}

func (req RegisterUserRequest) validate() (err error) {
	v := validator.New()
	err = v.Struct(req)
	return
}

func RegisterUser(APIstub shim.ChaincodeStubInterface, request util.Request) (response util.Response, err error) {
	var req RegisterUserRequest
	var res RegisterUserResponse

	// validate request
	req = createRegisterUserRequest(request.GetArguments())
	err = req.validate()
	if err != nil {
		err = util.ErrInvalidParam
		return
	}

	// check to user in blockchain
	var user User
	key := req.email
	err = util.GetState(APIstub, key, &user)
	if err != nil {
		if err != util.ErrNoDataFound { // no data found면, continue
			return
		}
	}

	user.email = req.email
	user.point = defaultPoint
	user.registeredAt = util.GetTimestamp() // getTimestamp 에러 처리 필요
	user.updateAt = user.registeredAt

	err = util.PutState(APIstub, key, user)
	if err != nil {
		return
	}

	res.user = user

	response.SetData(res)

	return
}
