// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateAppRequest openpitrix create app request
// swagger:model openpitrixCreateAppRequest
type OpenpitrixCreateAppRequest struct {

	// category id
	CategoryID string `json:"category_id,omitempty"`

	// chart name
	ChartName string `json:"chart_name,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// home
	Home string `json:"home,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// keywords
	Keywords string `json:"keywords,omitempty"`

	// maintainers
	Maintainers string `json:"maintainers,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// package
	Package strfmt.Base64 `json:"package,omitempty"`

	// readme
	Readme string `json:"readme,omitempty"`

	// repo id
	RepoID string `json:"repo_id,omitempty"`

	// screenshots
	Screenshots string `json:"screenshots,omitempty"`

	// sources
	Sources string `json:"sources,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this openpitrix create app request
func (m *OpenpitrixCreateAppRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateAppRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
