package models

import (
	"fmt"
	"strings"
)

// This file was generated by the gogi.
// Editing this file might prove futile when you re-run the gogi generate command

// This file was generated by the gogi tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ShipListParams ship list params
// swagger:model ShipListParams

// easyjson:json
type ShipListParams struct {

	// count
	Count int64 `json:"count,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// length
	Length int64 `json:"length,omitempty"`
}

// Validate validates this ship list params
func (m *ShipListParams) Validate() error {
	var res []error

	if len(res) > 0 {
		msgs := []string{}
		for _, err := range res {
			msgs = append(msgs, err.Error())
		}
		return fmt.Errorf("validation failure list: " + strings.Join(msgs, "; "))
	}
	return nil
}
