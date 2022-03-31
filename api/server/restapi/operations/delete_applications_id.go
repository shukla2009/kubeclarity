// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteApplicationsIDHandlerFunc turns a function with the right signature into a delete applications ID handler
type DeleteApplicationsIDHandlerFunc func(DeleteApplicationsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteApplicationsIDHandlerFunc) Handle(params DeleteApplicationsIDParams) middleware.Responder {
	return fn(params)
}

// DeleteApplicationsIDHandler interface for that can handle valid delete applications ID params
type DeleteApplicationsIDHandler interface {
	Handle(DeleteApplicationsIDParams) middleware.Responder
}

// NewDeleteApplicationsID creates a new http.Handler for the delete applications ID operation
func NewDeleteApplicationsID(ctx *middleware.Context, handler DeleteApplicationsIDHandler) *DeleteApplicationsID {
	return &DeleteApplicationsID{Context: ctx, Handler: handler}
}

/* DeleteApplicationsID swagger:route DELETE /applications/{id} deleteApplicationsId

Delete Application.

*/
type DeleteApplicationsID struct {
	Context *middleware.Context
	Handler DeleteApplicationsIDHandler
}

func (o *DeleteApplicationsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteApplicationsIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
