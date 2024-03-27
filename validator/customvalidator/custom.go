package customvalidator

import "github.com/go-playground/validator/v10"

type Custom struct {
	Validator *validator.Validate
}

func (c *Custom) Validate(i interface{}) error {
	if err := c.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}
