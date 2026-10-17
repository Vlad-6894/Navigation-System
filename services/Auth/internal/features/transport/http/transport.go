package auth_http_transport

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	pkg_http_server "navigation-pkg/transport/http/server"
)

type AuthHTTPTransport struct {
	service AuthService
}

type AuthService interface {
	RegisterProfile(
		ctx context.Context,
		profile auth_domains.Profile,
	) (auth_domains.Profile, error)
}

func NewAuthHTTPTransport(
	service AuthService,
) *AuthHTTPTransport {
	return &AuthHTTPTransport{
		service: service,
	}
}

func (h *AuthHTTPTransport) Routes() []pkg_http_server.Route {
	routes := []pkg_http_server.Route{
		{
			Path:   "",
			Method: "POST",
		},
	}

	return routes
}
