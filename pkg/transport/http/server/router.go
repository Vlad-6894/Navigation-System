package pkg_http_server

import (
	"fmt"
	"net/http"
)

type ApiVersion string

type PersonRole string

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	personRole PersonRole
}

func NewAPIVersionRouter(apiVersion ApiVersion, personRole PersonRole) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
		personRole: personRole,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes []Route) {
	for _, route := range routes {
		prefix := fmt.Sprintf("%s api/v%s/%s/%s", route.Method, r.apiVersion, r.personRole, route.Path)

		r.HandleFunc(prefix, route.Handler)
	}
}
