// Code generated by goa v3.21.1, DO NOT EDIT.
//
// catalog client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package catalog

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "catalog" service client.
type Client struct {
	RefreshEndpoint      goa.Endpoint
	RefreshAllEndpoint   goa.Endpoint
	CatalogErrorEndpoint goa.Endpoint
}

// NewClient initializes a "catalog" service client given the endpoints.
func NewClient(refresh, refreshAll, catalogError goa.Endpoint) *Client {
	return &Client{
		RefreshEndpoint:      refresh,
		RefreshAllEndpoint:   refreshAll,
		CatalogErrorEndpoint: catalogError,
	}
}

// Refresh calls the "Refresh" endpoint of the "catalog" service.
// Refresh may return the following errors:
//   - "internal-error" (type *goa.ServiceError): Internal Server Error
//   - "not-found" (type *goa.ServiceError): Resource Not Found Error
//   - error: internal error
func (c *Client) Refresh(ctx context.Context, p *RefreshPayload) (res *Job, err error) {
	var ires any
	ires, err = c.RefreshEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Job), nil
}

// RefreshAll calls the "RefreshAll" endpoint of the "catalog" service.
// RefreshAll may return the following errors:
//   - "internal-error" (type *goa.ServiceError): Internal Server Error
//   - "not-found" (type *goa.ServiceError): Resource Not Found Error
//   - error: internal error
func (c *Client) RefreshAll(ctx context.Context, p *RefreshAllPayload) (res []*Job, err error) {
	var ires any
	ires, err = c.RefreshAllEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Job), nil
}

// CatalogError calls the "CatalogError" endpoint of the "catalog" service.
// CatalogError may return the following errors:
//   - "internal-error" (type *goa.ServiceError): Internal Server Error
//   - "not-found" (type *goa.ServiceError): Resource Not Found Error
//   - error: internal error
func (c *Client) CatalogError(ctx context.Context, p *CatalogErrorPayload) (res *CatalogErrorResult, err error) {
	var ires any
	ires, err = c.CatalogErrorEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CatalogErrorResult), nil
}
