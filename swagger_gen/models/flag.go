// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Flag flag
// swagger:model flag

type Flag struct {

	// enabled data records will get data logging in the metrics pipeline, for example, kafka.
	DataRecordsEnabled bool `json:"dataRecordsEnabled,omitempty"`

	// description
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// id
	// Read Only: true
	// Minimum: 1
	ID int64 `json:"id,omitempty"`

	// segments
	Segments FlagSegments `json:"segments"`

	// variants
	Variants FlagVariants `json:"variants"`
}

/* polymorph flag dataRecordsEnabled false */

/* polymorph flag description false */

/* polymorph flag enabled false */

/* polymorph flag id false */

/* polymorph flag segments false */

/* polymorph flag variants false */

// Validate validates this flag
func (m *Flag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Flag) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *Flag) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Flag) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", int64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Flag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Flag) UnmarshalBinary(b []byte) error {
	var res Flag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
