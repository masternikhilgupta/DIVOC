// Code generated by go-swagger; DO NOT EDIT.

package certification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/divoc/api/swagger_gen/models"
)

// TestBulkCertifyHandlerFunc turns a function with the right signature into a test bulk certify handler
type TestBulkCertifyHandlerFunc func(TestBulkCertifyParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn TestBulkCertifyHandlerFunc) Handle(params TestBulkCertifyParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// TestBulkCertifyHandler interface for that can handle valid test bulk certify params
type TestBulkCertifyHandler interface {
	Handle(TestBulkCertifyParams, *models.JWTClaimBody) middleware.Responder
}

// NewTestBulkCertify creates a new http.Handler for the test bulk certify operation
func NewTestBulkCertify(ctx *middleware.Context, handler TestBulkCertifyHandler) *TestBulkCertify {
	return &TestBulkCertify{Context: ctx, Handler: handler}
}

/*TestBulkCertify swagger:route POST /test/bulkCertify certification testBulkCertify

Upload test certification csv for bulk ingestion

certify all the data in uploaded csv

*/
type TestBulkCertify struct {
	Context *middleware.Context
	Handler TestBulkCertifyHandler
}

func (o *TestBulkCertify) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTestBulkCertifyParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
