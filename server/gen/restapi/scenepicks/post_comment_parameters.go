// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPostCommentParams creates a new PostCommentParams object
// no default values defined in spec.
func NewPostCommentParams() PostCommentParams {

	return PostCommentParams{}
}

// PostCommentParams contains all the bound params for the post comment operation
// typically these are obtained from a http.Request
//
// swagger:parameters postComment
type PostCommentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Comment PostCommentBody
	/*
	  Required: true
	  In: path
	*/
	ID string
	/*
	  Required: true
	  In: header
	*/
	Token string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostCommentParams() beforehand.
func (o *PostCommentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PostCommentBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("comment", "body", ""))
			} else {
				res = append(res, errors.NewParseError("comment", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Comment = body
			}
		}
	} else {
		res = append(res, errors.Required("comment", "body", ""))
	}
	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindToken(r.Header[http.CanonicalHeaderKey("token")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PostCommentParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}

// bindToken binds and validates parameter Token from header.
func (o *PostCommentParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("token", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("token", "header", raw); err != nil {
		return err
	}

	o.Token = raw

	return nil
}
