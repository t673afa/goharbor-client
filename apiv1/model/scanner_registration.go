// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScannerRegistration Registration represents a named configuration for invoking a scanner via its adapter.
//
//
// swagger:model ScannerRegistration
type ScannerRegistration struct {

	// An optional value of the HTTP Authorization header sent with each request to the Scanner Adapter API.
	//
	AccessCredential string `json:"access_credential,omitempty"`

	// Optional property to describe the name of the scanner registration
	Adapter string `json:"adapter,omitempty"`

	// Specify what authentication approach is adopted for the HTTP communications.
	// Supported types Basic", "Bearer" and api key header "X-ScannerAdapter-API-Key"
	//
	Auth string `json:"auth,omitempty"`

	// An optional description of this registration.
	Description string `json:"description,omitempty"`

	// Indicate whether the registration is enabled or not
	Disabled *bool `json:"disabled,omitempty"`

	// Indicate the healthy of the registration
	Health string `json:"health,omitempty"`

	// Indicate if the registration is set as the system default one
	IsDefault *bool `json:"is_default,omitempty"`

	// The name of this registration.
	Name string `json:"name,omitempty"`

	// Indicate if skip the certificate verification when sending HTTP requests
	SkipCertVerify *bool `json:"skip_certVerify,omitempty"`

	// A base URL of the scanner adapter
	URL string `json:"url,omitempty"`

	// Indicate whether use internal registry addr for the scanner to pull content or not
	UseInternalAddr *bool `json:"use_internal_addr,omitempty"`

	// The unique identifier of this registration.
	UUID string `json:"uuid,omitempty"`

	// Optional property to describe the vendor of the scanner registration
	Vendor string `json:"vendor,omitempty"`

	// Optional property to describe the version of the scanner registration
	Version string `json:"version,omitempty"`
}

// Validate validates this scanner registration
func (m *ScannerRegistration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScannerRegistration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScannerRegistration) UnmarshalBinary(b []byte) error {
	var res ScannerRegistration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
