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

package operator_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeletePVCParams creates a new DeletePVCParams object
//
// There are no default values defined in the spec.
func NewDeletePVCParams() DeletePVCParams {

	return DeletePVCParams{}
}

// DeletePVCParams contains all the bound params for the delete p v c operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeletePVC
type DeletePVCParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	PVCName string
	/*
	  Required: true
	  In: path
	*/
	Namespace string
	/*
	  Required: true
	  In: path
	*/
	Tenant string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeletePVCParams() beforehand.
func (o *DeletePVCParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPVCName, rhkPVCName, _ := route.Params.GetOK("PVCName")
	if err := o.bindPVCName(rPVCName, rhkPVCName, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	rTenant, rhkTenant, _ := route.Params.GetOK("tenant")
	if err := o.bindTenant(rTenant, rhkTenant, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPVCName binds and validates parameter PVCName from path.
func (o *DeletePVCParams) bindPVCName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.PVCName = raw

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *DeletePVCParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Namespace = raw

	return nil
}

// bindTenant binds and validates parameter Tenant from path.
func (o *DeletePVCParams) bindTenant(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Tenant = raw

	return nil
}
