// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAPIInventoryAPIIDAPIInfoHandlerFunc turns a function with the right signature into a get API inventory API ID API info handler
type GetAPIInventoryAPIIDAPIInfoHandlerFunc func(GetAPIInventoryAPIIDAPIInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAPIInventoryAPIIDAPIInfoHandlerFunc) Handle(params GetAPIInventoryAPIIDAPIInfoParams) middleware.Responder {
	return fn(params)
}

// GetAPIInventoryAPIIDAPIInfoHandler interface for that can handle valid get API inventory API ID API info params
type GetAPIInventoryAPIIDAPIInfoHandler interface {
	Handle(GetAPIInventoryAPIIDAPIInfoParams) middleware.Responder
}

// NewGetAPIInventoryAPIIDAPIInfo creates a new http.Handler for the get API inventory API ID API info operation
func NewGetAPIInventoryAPIIDAPIInfo(ctx *middleware.Context, handler GetAPIInventoryAPIIDAPIInfoHandler) *GetAPIInventoryAPIIDAPIInfo {
	return &GetAPIInventoryAPIIDAPIInfo{Context: ctx, Handler: handler}
}

/* GetAPIInventoryAPIIDAPIInfo swagger:route GET /apiInventory/{apiId}/apiInfo getApiInventoryApiIdApiInfo

Get api info from apiId

*/
type GetAPIInventoryAPIIDAPIInfo struct {
	Context *middleware.Context
	Handler GetAPIInventoryAPIIDAPIInfoHandler
}

func (o *GetAPIInventoryAPIIDAPIInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAPIInventoryAPIIDAPIInfoParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
