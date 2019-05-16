// Code generated by go-swagger; DO NOT EDIT.

package currenttime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCurrentTimeHandlerFunc turns a function with the right signature into a get current time handler
type GetCurrentTimeHandlerFunc func(GetCurrentTimeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCurrentTimeHandlerFunc) Handle(params GetCurrentTimeParams) middleware.Responder {
	return fn(params)
}

// GetCurrentTimeHandler interface for that can handle valid get current time params
type GetCurrentTimeHandler interface {
	Handle(GetCurrentTimeParams) middleware.Responder
}

// NewGetCurrentTime creates a new http.Handler for the get current time operation
func NewGetCurrentTime(ctx *middleware.Context, handler GetCurrentTimeHandler) *GetCurrentTime {
	return &GetCurrentTime{Context: ctx, Handler: handler}
}

/*GetCurrentTime swagger:route GET /current-time currenttime getCurrentTime

Sign up Api takes params and register user

*/
type GetCurrentTime struct {
	Context *middleware.Context
	Handler GetCurrentTimeHandler
}

func (o *GetCurrentTime) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCurrentTimeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}