package responses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"battleship/app/generated/models"

	"github.com/mailru/easyjson"
)

const (
	UserCreateOKCode int = 200
)

/*UserCreateOK Returns created user

swagger:response userCreateOK
*/
type UserCreateOK struct {

	/*
	  In: Body
	*/
	Payload UserCreateOKBody `json:"body,omitempty"`

	// Binary payload for response writer
	BinPayload []byte
}

// NewUserCreateOK creates UserCreateOK with default headers values
func NewUserCreateOK() *UserCreateOK {
	return &UserCreateOK{}
}

// WithPayload adds the payload to the user create o k response
func (o *UserCreateOK) WithPayload(payload UserCreateOKBody) *UserCreateOK {
	o.Payload = payload
	return o
}

// WithBinPayload adds binnary payload to user create o k response
func (o *UserCreateOK) WithBinPayload(bytes []byte) *UserCreateOK {
	o.BinPayload = bytes
	return o
}

// WriteResponse to the client
func (o *UserCreateOK) WriteResponse(rw http.ResponseWriter) error {

	rw.WriteHeader(UserCreateOKCode)

	if o.BinPayload != nil {
		_, err := rw.Write(o.BinPayload)
		return err
	}
	payload := o.Payload

	_, _, err := easyjson.MarshalToHTTPResponseWriter(payload, rw)
	return err

	return nil
}

/*UserCreateBody user create body
swagger:model UserCreateBody
*/
// easyjson:json
type UserCreateBody struct {

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
	Params models.UserCreateParams `json:"params"`
}

/*UserCreateOKBody user create o k body
swagger:model UserCreateOKBody
*/
// easyjson:json
type UserCreateOKBody struct {
	models.BaseResponse

	// result
	// Required: true
	Result models.UserResponse `json:"result"`
}