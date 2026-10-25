package auth_http_transport

import (
	auth_domains "Auth_service/internal/core/domains"
	"encoding/json"
	pkg_logger "navigation-pkg/logger"
	pkg_http_response "navigation-pkg/transport/http/response"
	"net/http"
)

func (h *AuthHTTPTransport) HandleDispetchProfileToStudent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := pkg_logger.FromContext(ctx)
	responseHandler := pkg_http_response.NewHTTPResponseHandler(log, w)

	var dispetchStudent DispetchToStudentRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&dispetchStudent); err != nil {
		responseHandler.ErrorResponse(err, "fail decode dto!")
		return
	}

	names := auth_domains.NewNames(dispetchStudent.FirstName, dispetchStudent.Surname, dispetchStudent.MiddleName)

	if err := h.service.DispetchToStudent(ctx, names, dispetchStudent.ID.String(), dispetchStudent.ClassID, dispetchStudent.ParentID); err != nil {
		responseHandler.ErrorResponse(err, "fail to dispetch to student!")
		return
	}

	response := SuccessResponse{
		IsSuccess: "Success disperch to student!",
	}

	responseHandler.ToJSONResponse(response, http.StatusCreated)
}
