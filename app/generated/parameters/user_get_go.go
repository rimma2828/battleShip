package parameters

// This file was generated by the gogi.
// Editing this file might prove futile when you re-run the gogi generate command

import (
	"battleship/app/generated/models"

	"github.com/mailru/easyjson"
)

// UserGetGoParams contains all the bound params for the user get go operation
// typically these are obtained from a http.Request
//
// swagger:parameters user.get.go
type UserGetGoParams struct {
	/*
	  In: body
	*/
	Body UserGetGoBody `body:"body"`
}

func (o *UserGetGoParams) UnmarshalJSON(b []byte) error {
	return easyjson.Unmarshal(b, &o.Body)
}

// NewUserGetGoParams creates a new UserGetGoParams object
// with the default values initialized.
func NewUserGetGoParams() UserGetGoParams {
	var ()
	return UserGetGoParams{}
}

func (o *UserGetGoParams) Validate() error {

	return nil
}

/*UserGetGoBody user get go body
swagger:model UserGetGoBody
*/
// easyjson:json
type UserGetGoBody struct {

	// id
	// Required: true
	ID string `json:"id"`

	// jsonrpc
	// Required: true
	Jsonrpc string `json:"jsonrpc"`

	// method
	// Required: true
	Method string `json:"method"`

	// params
	// Required: true
	Params models.UserGetParams `json:"params"`
}

/*UserGetGoOKBody user get go o k body
swagger:model UserGetGoOKBody
*/
// easyjson:json
type UserGetGoOKBody struct {
	models.BaseResponse

	// result
	// Required: true
	Result models.UserGetResponse `json:"result"`
}
