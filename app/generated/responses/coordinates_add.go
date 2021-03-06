package responses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"battleship/app/generated/models"

	"github.com/mailru/easyjson"
)

const (
	CoordinatesAddOKCode int = 200
)

/*CoordinatesAddOK Returns status

swagger:response coordinatesAddOK
*/
type CoordinatesAddOK struct {

	/*
	  In: Body
	*/
	Payload CoordinatesAddOKBody `json:"body,omitempty"`

	// Binary payload for response writer
	BinPayload []byte
}

// NewCoordinatesAddOK creates CoordinatesAddOK with default headers values
func NewCoordinatesAddOK() *CoordinatesAddOK {
	return &CoordinatesAddOK{}
}

// WithPayload adds the payload to the coordinates add o k response
func (o *CoordinatesAddOK) WithPayload(payload CoordinatesAddOKBody) *CoordinatesAddOK {
	o.Payload = payload
	return o
}

// WithBinPayload adds binnary payload to coordinates add o k response
func (o *CoordinatesAddOK) WithBinPayload(bytes []byte) *CoordinatesAddOK {
	o.BinPayload = bytes
	return o
}

// WriteResponse to the client
func (o *CoordinatesAddOK) WriteResponse(rw http.ResponseWriter) error {

	rw.WriteHeader(CoordinatesAddOKCode)

	if o.BinPayload != nil {
		_, err := rw.Write(o.BinPayload)
		return err
	}
	payload := o.Payload

	_, _, err := easyjson.MarshalToHTTPResponseWriter(payload, rw)
	return err

	return nil
}

/*CoordinatesAddBody coordinates add body
swagger:model CoordinatesAddBody
*/
// easyjson:json
type CoordinatesAddBody struct {

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
	Params models.CoordinatesAddParams `json:"params"`
}

/*CoordinatesAddOKBody coordinates add o k body
swagger:model CoordinatesAddOKBody
*/
// easyjson:json
type CoordinatesAddOKBody struct {
	models.BaseResponse

	// result
	// Required: true
	Result models.CoordinatesActResponse `json:"result"`
}
