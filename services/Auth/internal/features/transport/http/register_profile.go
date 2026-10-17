package auth_http_transport

import (
	auth_domains "Auth_service/internal/core/domains"
	"encoding/json"
	pkg_logger "navigation-pkg/logger"
	pkg_http_response "navigation-pkg/transport/http/response"
	"net/http"
)

func (h *AuthHTTPTransport) HandleRegisterProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := pkg_logger.FromContext(ctx)
	responseHandler := pkg_http_response.NewHTTPResponseHandler(log, w)

	var profileDto RegisterProfileRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&profileDto); err != nil {
		responseHandler.ErrorResponse(err, "fail decode dto!")
		return
	}

	profile := profileRegDomainFromDto(profileDto)

	newProfile, err := h.service.RegisterProfile(ctx, profile)
	if err != nil {
		responseHandler.ErrorResponse(err, "fail register profile!")
		return
	}

	response := regProfileDtoFromDomain(newProfile)

	responseHandler.ToJSONResponse(response, http.StatusCreated)
}

func profileRegDomainFromDto(dto RegisterProfileRequestDTO) auth_domains.Profile {
	profile := auth_domains.NewUninitilizedProfile(dto.Login, dto.Password, dto.Age, dto.Role, dto.SchoolID)
	return profile
}

func regProfileDtoFromDomain(domain auth_domains.Profile) ProfileResponseDTO {
	dto := ProfileResponseDTO{
		ID:       domain.ID,
		Version:  domain.Version,
		Login:    domain.Login,
		Password: domain.Password,
		Age:      domain.Age,
		Role:     domain.Role,
		SchoolID: domain.SchoolID,
	}

	return dto
}
