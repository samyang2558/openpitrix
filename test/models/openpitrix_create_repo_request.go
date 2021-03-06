// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateRepoRequest openpitrix create repo request
// swagger:model openpitrixCreateRepoRequest
type OpenpitrixCreateRepoRequest struct {

	// category id
	CategoryID string `json:"category_id,omitempty"`

	// credential
	Credential string `json:"credential,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// labels
	Labels string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// providers
	Providers []string `json:"providers"`

	// selectors
	Selectors string `json:"selectors,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// visibility
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this openpitrix create repo request
func (m *OpenpitrixCreateRepoRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProviders(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixCreateRepoRequest) validateProviders(formats strfmt.Registry) error {

	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateRepoRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateRepoRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateRepoRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
