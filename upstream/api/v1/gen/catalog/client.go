// Code generated by goa v3.21.1, DO NOT EDIT.
//
// catalog client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package catalog

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "catalog" service client.
type Client struct {
	ListEndpoint goa.Endpoint
}

// NewClient initializes a "catalog" service client given the endpoints.
func NewClient(list goa.Endpoint) *Client {
	return &Client{
		ListEndpoint: list,
	}
}

// List calls the "List" endpoint of the "catalog" service.
// List may return the following errors:
//   - "internal-error" (type *goa.ServiceError): Internal Server Error
//   - error: internal error
func (c *Client) List(ctx context.Context) (res *ListResult, err error) {
	var ires any
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*ListResult), nil
}
