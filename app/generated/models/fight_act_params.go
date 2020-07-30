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

// FightActParams fight act params
// swagger:model FightActParams

// easyjson:json
type FightActParams struct {

	// ship Id
	ShipID int64 `json:"shipId,omitempty"`

	// user Id
	// Required: true
	UserID int64 `json:"userId"`

	// x coord
	// Required: true
	// Maximum: 10
	// Minimum: 1
	XCoord int64 `json:"xCoord"`

	// y coord
	// Required: true
	// Maximum: 10
	// Minimum: 1
	YCoord int64 `json:"yCoord"`
}

// Validate validates this fight act params
func (m *FightActParams) Validate() error {
	var res []error

	if err := m.validateUserID(); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateXCoord(); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateYCoord(); err != nil {
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

func (m *FightActParams) validateUserID() error {

	if err := validate.Required("userId", "", int64(m.UserID)); err != nil {
		return err
	}

	return nil
}

func (m *FightActParams) validateXCoord() error {

	if err := validate.Required("xCoord", "", int64(m.XCoord)); err != nil {
		return err
	}

	if err := validate.MinimumInt("xCoord", "", int64(m.XCoord), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("xCoord", "", int64(m.XCoord), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *FightActParams) validateYCoord() error {

	if err := validate.Required("yCoord", "", int64(m.YCoord)); err != nil {
		return err
	}

	if err := validate.MinimumInt("yCoord", "", int64(m.YCoord), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("yCoord", "", int64(m.YCoord), 10, false); err != nil {
		return err
	}

	return nil
}
