package utils

import (
	"errors"
	"finalProject/constants"
	"testing"
)

func Test_checkUserName(t *testing.T) {
	type test struct {
		username string
		expected error
	}
	var tests = []test{
		{"username", nil},
		{"tes", nil},
		{"t", errors.New(constants.USERNAME_FORMAT_ERROR)},
		{"te", errors.New(constants.USERNAME_FORMAT_ERROR)},
		{"azertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklmwxcvbn", errors.New(constants.USERNAME_FORMAT_ERROR)},
		{"&é'!'ç((", errors.New(constants.USERNAME_FORMAT_ERROR)},
	}

	for _, test := range tests {
		output := checkUserName(test.username)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestCheckPassword(t *testing.T) {
	type test struct {
		password string
		expected error
	}
	var tests = []test{
		{"password", nil},
		{"pass", errors.New(constants.PASSWORD_FORMAT_ERROR)},
		{"azertyuiopqsdfghjklmwxcvbnazertyuiopqsdf", errors.New(constants.PASSWORD_FORMAT_ERROR)},
	}

	for _, test := range tests {
		output := CheckPassword(test.password)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func Test_checkPinCode(t *testing.T) {
	type test struct {
		pincode  string
		expected error
	}
	var tests = []test{
		{"123456", nil},
		{"123", errors.New(constants.PIN_CODE_FORMAT_ERROR)},
		{"123456789", errors.New(constants.PIN_CODE_FORMAT_ERROR)},
		{"pincode", errors.New(constants.PIN_CODE_FORMAT_ERROR)},
	}

	for _, test := range tests {
		output := CheckPassword(test.pincode)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
