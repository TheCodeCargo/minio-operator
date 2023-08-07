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

	"github.com/go-openapi/runtime"

	"github.com/TheCodeCargo/minio-operator/models"
)

// TenantConfigurationOKCode is the HTTP code returned for type TenantConfigurationOK
const TenantConfigurationOKCode int = 200

/*
TenantConfigurationOK A successful response.

swagger:response tenantConfigurationOK
*/
type TenantConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.TenantConfigurationResponse `json:"body,omitempty"`
}

// NewTenantConfigurationOK creates TenantConfigurationOK with default headers values
func NewTenantConfigurationOK() *TenantConfigurationOK {

	return &TenantConfigurationOK{}
}

// WithPayload adds the payload to the tenant configuration o k response
func (o *TenantConfigurationOK) WithPayload(payload *models.TenantConfigurationResponse) *TenantConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tenant configuration o k response
func (o *TenantConfigurationOK) SetPayload(payload *models.TenantConfigurationResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TenantConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
TenantConfigurationDefault Generic error response.

swagger:response tenantConfigurationDefault
*/
type TenantConfigurationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTenantConfigurationDefault creates TenantConfigurationDefault with default headers values
func NewTenantConfigurationDefault(code int) *TenantConfigurationDefault {
	if code <= 0 {
		code = 500
	}

	return &TenantConfigurationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the tenant configuration default response
func (o *TenantConfigurationDefault) WithStatusCode(code int) *TenantConfigurationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the tenant configuration default response
func (o *TenantConfigurationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the tenant configuration default response
func (o *TenantConfigurationDefault) WithPayload(payload *models.Error) *TenantConfigurationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tenant configuration default response
func (o *TenantConfigurationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TenantConfigurationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
