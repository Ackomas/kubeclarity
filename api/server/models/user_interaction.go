// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// UserInteraction user interaction
//
// swagger:model UserInteraction
type UserInteraction string

func NewUserInteraction(value UserInteraction) *UserInteraction {
	v := value
	return &v
}

const (

	// UserInteractionNONE captures enum value "NONE"
	UserInteractionNONE UserInteraction = "NONE"

	// UserInteractionREQUIRED captures enum value "REQUIRED"
	UserInteractionREQUIRED UserInteraction = "REQUIRED"
)

// for schema
var userInteractionEnum []interface{}

func init() {
	var res []UserInteraction
	if err := json.Unmarshal([]byte(`["NONE","REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userInteractionEnum = append(userInteractionEnum, v)
	}
}

func (m UserInteraction) validateUserInteractionEnum(path, location string, value UserInteraction) error {
	if err := validate.EnumCase(path, location, value, userInteractionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user interaction
func (m UserInteraction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserInteractionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user interaction based on context it is used
func (m UserInteraction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}