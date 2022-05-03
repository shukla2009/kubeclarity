// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/kubeclarity/api/server/models"
)

// GetDashboardTrendsVulnerabilitiesOKCode is the HTTP code returned for type GetDashboardTrendsVulnerabilitiesOK
const GetDashboardTrendsVulnerabilitiesOKCode int = 200

/*GetDashboardTrendsVulnerabilitiesOK Success

swagger:response getDashboardTrendsVulnerabilitiesOK
*/
type GetDashboardTrendsVulnerabilitiesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.NewVulnerabilitiesTrend `json:"body,omitempty"`
}

// NewGetDashboardTrendsVulnerabilitiesOK creates GetDashboardTrendsVulnerabilitiesOK with default headers values
func NewGetDashboardTrendsVulnerabilitiesOK() *GetDashboardTrendsVulnerabilitiesOK {

	return &GetDashboardTrendsVulnerabilitiesOK{}
}

// WithPayload adds the payload to the get dashboard trends vulnerabilities o k response
func (o *GetDashboardTrendsVulnerabilitiesOK) WithPayload(payload []*models.NewVulnerabilitiesTrend) *GetDashboardTrendsVulnerabilitiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dashboard trends vulnerabilities o k response
func (o *GetDashboardTrendsVulnerabilitiesOK) SetPayload(payload []*models.NewVulnerabilitiesTrend) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDashboardTrendsVulnerabilitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.NewVulnerabilitiesTrend, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetDashboardTrendsVulnerabilitiesDefault unknown error

swagger:response getDashboardTrendsVulnerabilitiesDefault
*/
type GetDashboardTrendsVulnerabilitiesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetDashboardTrendsVulnerabilitiesDefault creates GetDashboardTrendsVulnerabilitiesDefault with default headers values
func NewGetDashboardTrendsVulnerabilitiesDefault(code int) *GetDashboardTrendsVulnerabilitiesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDashboardTrendsVulnerabilitiesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get dashboard trends vulnerabilities default response
func (o *GetDashboardTrendsVulnerabilitiesDefault) WithStatusCode(code int) *GetDashboardTrendsVulnerabilitiesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get dashboard trends vulnerabilities default response
func (o *GetDashboardTrendsVulnerabilitiesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get dashboard trends vulnerabilities default response
func (o *GetDashboardTrendsVulnerabilitiesDefault) WithPayload(payload *models.APIResponse) *GetDashboardTrendsVulnerabilitiesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dashboard trends vulnerabilities default response
func (o *GetDashboardTrendsVulnerabilitiesDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDashboardTrendsVulnerabilitiesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
