package util

import "errors"

var ErrInvalidParam = errors.New("invalid parameter")
var ErrInvalidFunc = errors.New("invalid function")

var ErrNotEnoughPoint = errors.New("not enough point to transfer")

var ErrAlreadyExists = errors.New("already data exists")
var ErrNoDataFound = errors.New("no data found")
var ErrGetStateFalied = errors.New("falied get data")
var ErrPutStateFailed = errors.New("failed put data")

var ErrMarshalFailed = errors.New("failed to parse with json or bytes")
var ErrValidateFailed = errors.New("failed to validate request")

var ErrCannotRegisteredUser = errors.New("failed to register the user")
var ErrCannotTransferPoint = errors.New("failed to transfer")
var ErrCannotGetPoint = errors.New("failed to get point of requester")
