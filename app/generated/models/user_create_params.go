package models

// This file was generated by the gogi.
// Editing this file might prove futile when you re-run the gogi generate command

import (
	"fmt"
	"strings"

	"github.com/go-openapi/validate"
)

// This file was generated by the gogi tool.
// Editing this file might prove futile when you re-run the swagger generate command

// UserCreateParams user create params
// swagger:model UserCreateParams

// easyjson:json
type UserCreateParams struct {

	// user name
	// Required: true
	UserName string `json:"userName"`
}

// Validate validates this user create params
func (m *UserCreateParams) Validate() error {
	var res []error

	if err := m.validateUserName(); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		msgs := []string{}
		for _, err := range res {
			msgs = append(msgs, err.Error())
		}
		return fmt.Errorf("validation failure list: " + strings.Join(msgs, "; "))
	}
	return nil
}

func (m *UserCreateParams) validateUserName() error {

	if err := validate.RequiredString("userName", "", string(m.UserName)); err != nil {
		return err
	}

	return nil
}