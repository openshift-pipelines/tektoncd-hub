// Code generated by goa v3.21.1, DO NOT EDIT.
//
// resource endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package resource

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "resource" service endpoints.
type Endpoints struct {
	Query                              goa.Endpoint
	List                               goa.Endpoint
	VersionsByID                       goa.Endpoint
	ByCatalogKindNameVersion           goa.Endpoint
	ByCatalogKindNameVersionReadme     goa.Endpoint
	ByCatalogKindNameVersionYaml       goa.Endpoint
	ByVersionID                        goa.Endpoint
	ByCatalogKindName                  goa.Endpoint
	ByID                               goa.Endpoint
	GetRawYamlByCatalogKindNameVersion goa.Endpoint
	GetLatestRawYamlByCatalogKindName  goa.Endpoint
}

// GetRawYamlByCatalogKindNameVersionResponseData holds both the result and the
// HTTP response body reader of the "GetRawYamlByCatalogKindNameVersion" method.
type GetRawYamlByCatalogKindNameVersionResponseData struct {
	// Body streams the HTTP response body.
	Body io.ReadCloser
}

// GetLatestRawYamlByCatalogKindNameResponseData holds both the result and the
// HTTP response body reader of the "GetLatestRawYamlByCatalogKindName" method.
type GetLatestRawYamlByCatalogKindNameResponseData struct {
	// Body streams the HTTP response body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "resource" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Query:                              NewQueryEndpoint(s),
		List:                               NewListEndpoint(s),
		VersionsByID:                       NewVersionsByIDEndpoint(s),
		ByCatalogKindNameVersion:           NewByCatalogKindNameVersionEndpoint(s),
		ByCatalogKindNameVersionReadme:     NewByCatalogKindNameVersionReadmeEndpoint(s),
		ByCatalogKindNameVersionYaml:       NewByCatalogKindNameVersionYamlEndpoint(s),
		ByVersionID:                        NewByVersionIDEndpoint(s),
		ByCatalogKindName:                  NewByCatalogKindNameEndpoint(s),
		ByID:                               NewByIDEndpoint(s),
		GetRawYamlByCatalogKindNameVersion: NewGetRawYamlByCatalogKindNameVersionEndpoint(s),
		GetLatestRawYamlByCatalogKindName:  NewGetLatestRawYamlByCatalogKindNameEndpoint(s),
	}
}

// Use applies the given middleware to all the "resource" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Query = m(e.Query)
	e.List = m(e.List)
	e.VersionsByID = m(e.VersionsByID)
	e.ByCatalogKindNameVersion = m(e.ByCatalogKindNameVersion)
	e.ByCatalogKindNameVersionReadme = m(e.ByCatalogKindNameVersionReadme)
	e.ByCatalogKindNameVersionYaml = m(e.ByCatalogKindNameVersionYaml)
	e.ByVersionID = m(e.ByVersionID)
	e.ByCatalogKindName = m(e.ByCatalogKindName)
	e.ByID = m(e.ByID)
	e.GetRawYamlByCatalogKindNameVersion = m(e.GetRawYamlByCatalogKindNameVersion)
	e.GetLatestRawYamlByCatalogKindName = m(e.GetLatestRawYamlByCatalogKindName)
}

// NewQueryEndpoint returns an endpoint function that calls the method "Query"
// of service "resource".
func NewQueryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*QueryPayload)
		res, err := s.Query(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResources(res, "default")
		return vres, nil
	}
}

// NewListEndpoint returns an endpoint function that calls the method "List" of
// service "resource".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPayload)
		res, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResources(res, "default")
		return vres, nil
	}
}

// NewVersionsByIDEndpoint returns an endpoint function that calls the method
// "VersionsByID" of service "resource".
func NewVersionsByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*VersionsByIDPayload)
		res, err := s.VersionsByID(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResourceVersions(res, "default")
		return vres, nil
	}
}

// NewByCatalogKindNameVersionEndpoint returns an endpoint function that calls
// the method "ByCatalogKindNameVersion" of service "resource".
func NewByCatalogKindNameVersionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByCatalogKindNameVersionPayload)
		res, err := s.ByCatalogKindNameVersion(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResourceVersion(res, "default")
		return vres, nil
	}
}

// NewByCatalogKindNameVersionReadmeEndpoint returns an endpoint function that
// calls the method "ByCatalogKindNameVersionReadme" of service "resource".
func NewByCatalogKindNameVersionReadmeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByCatalogKindNameVersionReadmePayload)
		res, err := s.ByCatalogKindNameVersionReadme(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResourceVersionReadme(res, "default")
		return vres, nil
	}
}

// NewByCatalogKindNameVersionYamlEndpoint returns an endpoint function that
// calls the method "ByCatalogKindNameVersionYaml" of service "resource".
func NewByCatalogKindNameVersionYamlEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByCatalogKindNameVersionYamlPayload)
		res, err := s.ByCatalogKindNameVersionYaml(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResourceVersionYaml(res, "default")
		return vres, nil
	}
}

// NewByVersionIDEndpoint returns an endpoint function that calls the method
// "ByVersionId" of service "resource".
func NewByVersionIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByVersionIDPayload)
		res, err := s.ByVersionID(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResourceVersion(res, "default")
		return vres, nil
	}
}

// NewByCatalogKindNameEndpoint returns an endpoint function that calls the
// method "ByCatalogKindName" of service "resource".
func NewByCatalogKindNameEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByCatalogKindNamePayload)
		res, err := s.ByCatalogKindName(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResource(res, "default")
		return vres, nil
	}
}

// NewByIDEndpoint returns an endpoint function that calls the method "ById" of
// service "resource".
func NewByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByIDPayload)
		res, err := s.ByID(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResource(res, "default")
		return vres, nil
	}
}

// NewGetRawYamlByCatalogKindNameVersionEndpoint returns an endpoint function
// that calls the method "GetRawYamlByCatalogKindNameVersion" of service
// "resource".
func NewGetRawYamlByCatalogKindNameVersionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetRawYamlByCatalogKindNameVersionPayload)
		body, err := s.GetRawYamlByCatalogKindNameVersion(ctx, p)
		if err != nil {
			return nil, err
		}
		return &GetRawYamlByCatalogKindNameVersionResponseData{Body: body}, nil
	}
}

// NewGetLatestRawYamlByCatalogKindNameEndpoint returns an endpoint function
// that calls the method "GetLatestRawYamlByCatalogKindName" of service
// "resource".
func NewGetLatestRawYamlByCatalogKindNameEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetLatestRawYamlByCatalogKindNamePayload)
		body, err := s.GetLatestRawYamlByCatalogKindName(ctx, p)
		if err != nil {
			return nil, err
		}
		return &GetLatestRawYamlByCatalogKindNameResponseData{Body: body}, nil
	}
}
