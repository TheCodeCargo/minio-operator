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

// OperatorSubnetLoginOKCode is the HTTP code returned for type OperatorSubnetLoginOK
const OperatorSubnetLoginOKCode int = 200

/*
OperatorSubnetLoginOK A successful response.

swagger:response operatorSubnetLoginOK
*/
type OperatorSubnetLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.OperatorSubnetLoginResponse `json:"body,omitempty"`
}

// NewOperatorSubnetLoginOK creates OperatorSubnetLoginOK with default headers values
func NewOperatorSubnetLoginOK() *OperatorSubnetLoginOK {

	return &OperatorSubnetLoginOK{}
}

// WithPayload adds the payload to the operator subnet login o k response
func (o *OperatorSubnetLoginOK) WithPayload(payload *models.OperatorSubnetLoginResponse) *OperatorSubnetLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the operator subnet login o k response
func (o *OperatorSubnetLoginOK) SetPayload(payload *models.OperatorSubnetLoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OperatorSubnetLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
OperatorSubnetLoginDefault Generic error response.

swagger:response operatorSubnetLoginDefault
*/
type OperatorSubnetLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewOperatorSubnetLoginDefault creates OperatorSubnetLoginDefault with default headers values
func NewOperatorSubnetLoginDefault(code int) *OperatorSubnetLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &OperatorSubnetLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the operator subnet login default response
func (o *OperatorSubnetLoginDefault) WithStatusCode(code int) *OperatorSubnetLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the operator subnet login default response
func (o *OperatorSubnetLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the operator subnet login default response
func (o *OperatorSubnetLoginDefault) WithPayload(payload *models.Error) *OperatorSubnetLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the operator subnet login default response
func (o *OperatorSubnetLoginDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OperatorSubnetLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
