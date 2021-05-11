// Code generated by go-swagger; DO NOT EDIT.

package baremetal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// NewPostCiNodesParams creates a new PostCiNodesParams object
// with the default values initialized.
func NewPostCiNodesParams() *PostCiNodesParams {
	var ()
	return &PostCiNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCiNodesParamsWithTimeout creates a new PostCiNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCiNodesParamsWithTimeout(timeout time.Duration) *PostCiNodesParams {
	var ()
	return &PostCiNodesParams{

		timeout: timeout,
	}
}

// NewPostCiNodesParamsWithContext creates a new PostCiNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCiNodesParamsWithContext(ctx context.Context) *PostCiNodesParams {
	var ()
	return &PostCiNodesParams{

		Context: ctx,
	}
}

// NewPostCiNodesParamsWithHTTPClient creates a new PostCiNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCiNodesParamsWithHTTPClient(client *http.Client) *PostCiNodesParams {
	var ()
	return &PostCiNodesParams{
		HTTPClient: client,
	}
}

/*PostCiNodesParams contains all the parameters to send to the API endpoint
for the post ci nodes operation typically these are written to a http.Request
*/
type PostCiNodesParams struct {

	/*CiNode
	  CI node to create

	*/
	CiNode *models.MutableCiNode

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ci nodes params
func (o *PostCiNodesParams) WithTimeout(timeout time.Duration) *PostCiNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ci nodes params
func (o *PostCiNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ci nodes params
func (o *PostCiNodesParams) WithContext(ctx context.Context) *PostCiNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ci nodes params
func (o *PostCiNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ci nodes params
func (o *PostCiNodesParams) WithHTTPClient(client *http.Client) *PostCiNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ci nodes params
func (o *PostCiNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCiNode adds the ciNode to the post ci nodes params
func (o *PostCiNodesParams) WithCiNode(ciNode *models.MutableCiNode) *PostCiNodesParams {
	o.SetCiNode(ciNode)
	return o
}

// SetCiNode adds the ciNode to the post ci nodes params
func (o *PostCiNodesParams) SetCiNode(ciNode *models.MutableCiNode) {
	o.CiNode = ciNode
}

// WriteToRequest writes these params to a swagger request
func (o *PostCiNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CiNode != nil {
		if err := r.SetBodyParam(o.CiNode); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
