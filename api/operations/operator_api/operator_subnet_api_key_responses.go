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

// OperatorSubnetAPIKeyOKCode is the HTTP code returned for type OperatorSubnetAPIKeyOK
const OperatorSubnetAPIKeyOKCode int = 200

/*
OperatorSubnetAPIKeyOK A successful response.

swagger:response operatorSubnetApiKeyOK
*/
type OperatorSubnetAPIKeyOK struct {

	/*
	  In: Body
	*/
	Payload *models.OperatorSubnetAPIKey `json:"body,omitempty"`
}

// NewOperatorSubnetAPIKeyOK creates OperatorSubnetAPIKeyOK with default headers values
func NewOperatorSubnetAPIKeyOK() *OperatorSubnetAPIKeyOK {

	return &OperatorSubnetAPIKeyOK{}
}

// WithPayload adds the payload to the operator subnet Api key o k response
func (o *OperatorSubnetAPIKeyOK) WithPayload(payload *models.OperatorSubnetAPIKey) *OperatorSubnetAPIKeyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the operator subnet Api key o k response
func (o *OperatorSubnetAPIKeyOK) SetPayload(payload *models.OperatorSubnetAPIKey) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OperatorSubnetAPIKeyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
OperatorSubnetAPIKeyDefault Generic error response.

swagger:response operatorSubnetApiKeyDefault
*/
type OperatorSubnetAPIKeyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewOperatorSubnetAPIKeyDefault creates OperatorSubnetAPIKeyDefault with default headers values
func NewOperatorSubnetAPIKeyDefault(code int) *OperatorSubnetAPIKeyDefault {
	if code <= 0 {
		code = 500
	}

	return &OperatorSubnetAPIKeyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the operator subnet Api key default response
func (o *OperatorSubnetAPIKeyDefault) WithStatusCode(code int) *OperatorSubnetAPIKeyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the operator subnet Api key default response
func (o *OperatorSubnetAPIKeyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the operator subnet Api key default response
func (o *OperatorSubnetAPIKeyDefault) WithPayload(payload *models.Error) *OperatorSubnetAPIKeyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the operator subnet Api key default response
func (o *OperatorSubnetAPIKeyDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OperatorSubnetAPIKeyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
