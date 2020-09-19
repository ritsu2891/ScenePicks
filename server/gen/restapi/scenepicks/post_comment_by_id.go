// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostCommentByIDHandlerFunc turns a function with the right signature into a post comment by Id handler
type PostCommentByIDHandlerFunc func(PostCommentByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCommentByIDHandlerFunc) Handle(params PostCommentByIDParams) middleware.Responder {
	return fn(params)
}

// PostCommentByIDHandler interface for that can handle valid post comment by Id params
type PostCommentByIDHandler interface {
	Handle(PostCommentByIDParams) middleware.Responder
}

// NewPostCommentByID creates a new http.Handler for the post comment by Id operation
func NewPostCommentByID(ctx *middleware.Context, handler PostCommentByIDHandler) *PostCommentByID {
	return &PostCommentByID{Context: ctx, Handler: handler}
}

/*PostCommentByID swagger:route POST /dialog/{id}/comment postCommentById

PostCommentByID post comment by Id API

*/
type PostCommentByID struct {
	Context *middleware.Context
	Handler PostCommentByIDHandler
}

func (o *PostCommentByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostCommentByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostCommentByIDBody post comment by ID body
//
// swagger:model PostCommentByIDBody
type PostCommentByIDBody struct {

	// comment
	Comment string `json:"comment,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *PostCommentByIDBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// comment
		Comment string `json:"comment,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Comment = props.Comment
	return nil
}

// Validate validates this post comment by ID body
func (o *PostCommentByIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCommentByIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCommentByIDBody) UnmarshalBinary(b []byte) error {
	var res PostCommentByIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}