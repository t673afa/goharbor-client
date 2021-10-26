// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CVEWhitelistItem The item in CVE whitelist
//
// swagger:model CVEWhitelistItem
type CVEWhitelistItem struct {

	// The ID of the CVE, such as "CVE-2019-10164"
	CveID string `json:"cve_id,omitempty"`
}

// Validate validates this c v e whitelist item
func (m *CVEWhitelistItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CVEWhitelistItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CVEWhitelistItem) UnmarshalBinary(b []byte) error {
	var res CVEWhitelistItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
