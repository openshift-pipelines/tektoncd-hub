// Code generated by goa v3.21.1, DO NOT EDIT.
//
// catalog HTTP client types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package client

import (
	catalog "github.com/tektoncd/hub/api/v1/gen/catalog"
	goa "goa.design/goa/v3/pkg"
)

// ListResponseBody is the type of the "catalog" service "List" endpoint HTTP
// response body.
type ListResponseBody struct {
	Data []*CatalogResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// ListInternalErrorResponseBody is the type of the "catalog" service "List"
// endpoint HTTP response body for the "internal-error" error.
type ListInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CatalogResponseBody is used to define fields on response body types.
type CatalogResponseBody struct {
	// ID is the unique id of the catalog
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of catalog
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Type of catalog
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// URL of catalog
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// Provider of catalog
	Provider *string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
}

// NewListResultOK builds a "catalog" service "List" endpoint result from a
// HTTP "OK" response.
func NewListResultOK(body *ListResponseBody) *catalog.ListResult {
	v := &catalog.ListResult{}
	if body.Data != nil {
		v.Data = make([]*catalog.Catalog, len(body.Data))
		for i, val := range body.Data {
			v.Data[i] = unmarshalCatalogResponseBodyToCatalogCatalog(val)
		}
	}

	return v
}

// NewListInternalError builds a catalog service List endpoint internal-error
// error.
func NewListInternalError(body *ListInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateListResponseBody runs the validations defined on ListResponseBody
func ValidateListResponseBody(body *ListResponseBody) (err error) {
	for _, e := range body.Data {
		if e != nil {
			if err2 := ValidateCatalogResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateListInternalErrorResponseBody runs the validations defined on
// List_internal-error_Response_Body
func ValidateListInternalErrorResponseBody(body *ListInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCatalogResponseBody runs the validations defined on
// CatalogResponseBody
func ValidateCatalogResponseBody(body *CatalogResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	if body.Provider == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("provider", "body"))
	}
	if body.Type != nil {
		if !(*body.Type == "official" || *body.Type == "community") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.type", *body.Type, []any{"official", "community"}))
		}
	}
	return
}
