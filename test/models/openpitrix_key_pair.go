// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixKeyPair openpitrix key pair
// swagger:model openpitrixKeyPair
type OpenpitrixKeyPair struct {

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// key pair id
	KeyPairID string `json:"key_pair_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// node id
	NodeID []string `json:"node_id"`

	// owner
	Owner string `json:"owner,omitempty"`

	// pub key
	PubKey string `json:"pub_key,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`
}

// Validate validates this openpitrix key pair
func (m *OpenpitrixKeyPair) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixKeyPair) validateNodeID(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixKeyPair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixKeyPair) UnmarshalBinary(b []byte) error {
	var res OpenpitrixKeyPair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}