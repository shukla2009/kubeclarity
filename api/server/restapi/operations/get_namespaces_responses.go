// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/kubeclarity/api/server/models"
)

// GetNamespacesOKCode is the HTTP code returned for type GetNamespacesOK
const GetNamespacesOKCode int = 200

/*GetNamespacesOK Success

swagger:response getNamespacesOK
*/
type GetNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Namespace `json:"body,omitempty"`
}

// NewGetNamespacesOK creates GetNamespacesOK with default headers values
func NewGetNamespacesOK() *GetNamespacesOK {

	return &GetNamespacesOK{}
}

// WithPayload adds the payload to the get namespaces o k response
func (o *GetNamespacesOK) WithPayload(payload []*models.Namespace) *GetNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get namespaces o k response
func (o *GetNamespacesOK) SetPayload(payload []*models.Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Namespace, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetNamespacesDefault unknown error

swagger:response getNamespacesDefault
*/
type GetNamespacesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetNamespacesDefault creates GetNamespacesDefault with default headers values
func NewGetNamespacesDefault(code int) *GetNamespacesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetNamespacesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get namespaces default response
func (o *GetNamespacesDefault) WithStatusCode(code int) *GetNamespacesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get namespaces default response
func (o *GetNamespacesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get namespaces default response
func (o *GetNamespacesDefault) WithPayload(payload *models.APIResponse) *GetNamespacesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get namespaces default response
func (o *GetNamespacesDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNamespacesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
