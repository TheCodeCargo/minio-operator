// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Operator
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServerProperties server properties
//
// swagger:model serverProperties
type ServerProperties struct {

	// commit ID
	CommitID string `json:"commitID,omitempty"`

	// drives
	Drives []*ServerDrives `json:"drives"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// network
	Network map[string]string `json:"network,omitempty"`

	// pool number
	PoolNumber int64 `json:"poolNumber,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// uptime
	Uptime int64 `json:"uptime,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this server properties
func (m *ServerProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrives(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerProperties) validateDrives(formats strfmt.Registry) error {
	if swag.IsZero(m.Drives) { // not required
		return nil
	}

	for i := 0; i < len(m.Drives); i++ {
		if swag.IsZero(m.Drives[i]) { // not required
			continue
		}

		if m.Drives[i] != nil {
			if err := m.Drives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this server properties based on the context it is used
func (m *ServerProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDrives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerProperties) contextValidateDrives(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Drives); i++ {

		if m.Drives[i] != nil {

			if swag.IsZero(m.Drives[i]) { // not required
				return nil
			}

			if err := m.Drives[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerProperties) UnmarshalBinary(b []byte) error {
	var res ServerProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
