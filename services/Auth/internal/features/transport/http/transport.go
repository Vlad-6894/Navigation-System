package auth_http_transport

import (
	auth_domains "Auth_service/internal/core/domains"
	"context"
	pkg_jwt "navigation-pkg/jwt"
	pkg_http_server "navigation-pkg/transport/http/server"
	"net/http"
)

type AuthHTTPTransport struct {
	service AuthService
}

type AuthService interface {
	RegisterProfile(
		ctx context.Context,
		profile auth_domains.Profile,
	) (auth_domains.Profile, error)

	LoginProfile(
		ctx context.Context,
		login string,
		password string,
	) (pkg_jwt.TokensPair, error)
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
			Path:    "auth/profile/register",
			Method:  http.MethodPost,
			Handler: h.HandleRegisterProfile,
		},

		{
			Path:    "auth/profile/login",
			Method:  http.MethodPost,
			Handler: h.HandleLoginProfile,
		},
	}

	return routes
}
