// Code generated by goa v3.21.1, DO NOT EDIT.
//
// status HTTP server
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	"context"
	"net/http"

	status "github.com/tektoncd/hub/api/gen/status"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the status service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Status http.Handler
	CORS   http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the status service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *status.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Status", "GET", "/"},
			{"Status", "GET", "/v1"},
			{"CORS", "OPTIONS", "/"},
			{"CORS", "OPTIONS", "/v1"},
		},
		Status: NewStatusHandler(e.Status, mux, decoder, encoder, errhandler, formatter),
		CORS:   NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "status" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Status = m(s.Status)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return status.MethodNames[:] }

// Mount configures the mux to serve the status endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountStatusHandler(mux, h.Status)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the status endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountStatusHandler configures the mux to serve the "status" service "Status"
// endpoint.
func MountStatusHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStatusOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/", f)
	mux.Handle("GET", "/v1", f)
}

// NewStatusHandler creates a HTTP handler which loads the HTTP request and
// calls the "status" service "Status" endpoint.
func NewStatusHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeStatusResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Status")
		ctx = context.WithValue(ctx, goa.ServiceKey, "status")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service status.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleStatusOrigin(h)
	mux.Handle("OPTIONS", "/", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleStatusOrigin applies the CORS response headers corresponding to the
// origin for the service status.
func HandleStatusOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
