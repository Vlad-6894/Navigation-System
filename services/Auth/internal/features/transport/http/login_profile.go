package auth_http_transport

import (
	"encoding/json"
	pkg_logger "navigation-pkg/logger"
	pkg_http_response "navigation-pkg/transport/http/response"
	"net/http"
)

func (h *AuthHTTPTransport) HandleLoginProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := pkg_logger.FromContext(ctx)
	responseHandler := pkg_http_response.NewHTTPResponseHandler(log, w)

	var loginProfileDto LoginProfileRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&loginProfileDto); err != nil {
		responseHandler.ErrorResponse(err, "fail decode dto!")
		return
	}

	pair, err := h.service.LoginProfile(ctx, loginProfileDto.Login, loginProfileDto.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "fail login profile!")
		return
	}

	response := TokenPareResponse{
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
	}

	responseHandler.ToJSONResponse(response, http.StatusOK)
}
