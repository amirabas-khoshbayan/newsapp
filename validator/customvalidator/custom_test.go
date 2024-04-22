package customvalidator

import (
	"github.com/go-playground/validator/v10"
	"newsapp/param/userparam"
	"testing"
)

func TestCustomValidatorCreateUser(t *testing.T) {
	type input struct {
		request *userparam.CreateNewUserRequest
		result  bool
	}

	var testCases = []input{
		{
			request: &userparam.CreateNewUserRequest{
				PhoneNumber: "test",
				Password:    "",
			},
			result: false,
		},
		{
			request: &userparam.CreateNewUserRequest{
				PhoneNumber: "",
				Password:    "test",
			},
			result: false,
		},
		{
			request: &userparam.CreateNewUserRequest{
				PhoneNumber: "09158888888",
				Password:    "test",
			},
			result: true,
		},
	}

	custValidate := new(Custom)
	custValidate.Validator = validator.New()
	for _, tc := range testCases {
		res := custValidate.Validate(tc.request)
		if (res != nil && tc.result) || (res == nil && !tc.result) {
			t.Fatalf("want %v , got %v", tc.result, res)
		}
	}
}
