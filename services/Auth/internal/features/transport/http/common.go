package auth_http_transport

import "github.com/google/uuid"

type RegisterProfileRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Role     string `json:"role"`
	SchoolID int    `json:"school_id"`
}

type ProfileResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Version  int       `json:"version"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Age      int       `json:"age"`
	Role     string    `json:"role"`
	SchoolID int       `json:"school_id"`
}
