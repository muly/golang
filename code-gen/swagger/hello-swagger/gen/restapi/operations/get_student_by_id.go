// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetStudentByIDHandlerFunc turns a function with the right signature into a get student by ID handler
type GetStudentByIDHandlerFunc func(GetStudentByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStudentByIDHandlerFunc) Handle(params GetStudentByIDParams) middleware.Responder {
	return fn(params)
}

// GetStudentByIDHandler interface for that can handle valid get student by ID params
type GetStudentByIDHandler interface {
	Handle(GetStudentByIDParams) middleware.Responder
}

// NewGetStudentByID creates a new http.Handler for the get student by ID operation
func NewGetStudentByID(ctx *middleware.Context, handler GetStudentByIDHandler) *GetStudentByID {
	return &GetStudentByID{Context: ctx, Handler: handler}
}

/* GetStudentByID swagger:route GET /api/students/{id} getStudentById

GetStudentByID get student by ID API

*/
type GetStudentByID struct {
	Context *middleware.Context
	Handler GetStudentByIDHandler
}

func (o *GetStudentByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetStudentByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}