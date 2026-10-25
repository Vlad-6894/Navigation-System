package auth_http_transport

import (
	auth_domains "Auth_service/internal/core/domains"
	"encoding/json"
	pkg_logger "navigation-pkg/logger"
	pkg_http_response "navigation-pkg/transport/http/response"
	"net/http"
)

func (h *AuthHTTPTransport) HandleDispetchProfileToParent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := pkg_logger.FromContext(ctx)
	responseHandler := pkg_http_response.NewHTTPResponseHandler(log, w)

	var dispetchParent DispetchToParentRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&dispetchParent); err != nil {
		responseHandler.ErrorResponse(err, "fail decode dto!")
		return
	}

	names := auth_domains.NewNames(dispetchParent.FirstName, dispetchParent.Surname, dispetchParent.MiddleName)

	if err := h.service.DispetchToParent(ctx, names, dispetchParent.ID.String()); err != nil {
		responseHandler.ErrorResponse(err, "fail to dispetch to parent!")
		return
	}

	response := SuccessResponse{
		IsSuccess: "Success disperch to parent!",
	}

	responseHandler.ToJSONResponse(response, http.StatusCreated)
}
