package constants

import (
	"errors"
	"regexp"
)

var RegexUsername = regexp.MustCompile(`^[a-z0-9_]{3,100}$`)
var RegexPassword = regexp.MustCompile(`^.{6,32}$`)
var RegexPinCode = regexp.MustCompile(`^\d{6}$`)

const (
	READ_BODY_ERROR_MESSAGE   = "Failed to read body:"
	DESERIALIZE_ERROR_MESSAGE = "Failed to unserialize body"
	SERIALIZE_ERROR_MESSAGE   = "Failed to serialize response"

	HTTP_METHODE_NOT_ALLOWED = "Method not allowed"
	USERNAME_FORMAT_ERROR    = "Username does not respect the naming policy"
	PASSWORD_FORMAT_ERROR    = "Password does not respect the complexity policy"
	PIN_CODE_FORMAT_ERROR    = "Pincode must contains 6 digits"

	MOCK_SERVER_PORT    = "5002"
	MOCK_SERVER_ADDRESS = "mock"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)
