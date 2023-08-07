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

// ListTenantsOKCode is the HTTP code returned for type ListTenantsOK
const ListTenantsOKCode int = 200

/*
ListTenantsOK A successful response.

swagger:response listTenantsOK
*/
type ListTenantsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListTenantsResponse `json:"body,omitempty"`
}

// NewListTenantsOK creates ListTenantsOK with default headers values
func NewListTenantsOK() *ListTenantsOK {

	return &ListTenantsOK{}
}

// WithPayload adds the payload to the list tenants o k response
func (o *ListTenantsOK) WithPayload(payload *models.ListTenantsResponse) *ListTenantsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list tenants o k response
func (o *ListTenantsOK) SetPayload(payload *models.ListTenantsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTenantsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListTenantsDefault Generic error response.

swagger:response listTenantsDefault
*/
type ListTenantsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListTenantsDefault creates ListTenantsDefault with default headers values
func NewListTenantsDefault(code int) *ListTenantsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListTenantsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list tenants default response
func (o *ListTenantsDefault) WithStatusCode(code int) *ListTenantsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list tenants default response
func (o *ListTenantsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list tenants default response
func (o *ListTenantsDefault) WithPayload(payload *models.Error) *ListTenantsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list tenants default response
func (o *ListTenantsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTenantsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
