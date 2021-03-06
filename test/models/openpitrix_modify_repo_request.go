// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyRepoRequest openpitrix modify repo request
// swagger:model openpitrixModifyRepoRequest
type OpenpitrixModifyRepoRequest struct {

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

	// repo id
	RepoID string `json:"repo_id,omitempty"`

	// selectors
	Selectors string `json:"selectors,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// visibility
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this openpitrix modify repo request
func (m *OpenpitrixModifyRepoRequest) Validate(formats strfmt.Registry) error {
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

func (m *OpenpitrixModifyRepoRequest) validateProviders(formats strfmt.Registry) error {

	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyRepoRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyRepoRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyRepoRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
