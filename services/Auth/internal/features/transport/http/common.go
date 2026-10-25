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

type LoginProfileRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenPareResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type DispetchToParentRequestDTO struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	Surname    string    `json:"surname"`
	MiddleName string    `json:"middle_name"`
}

type DispetchToStudentRequestDTO struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	Surname    string    `json:"surname"`
	MiddleName string    `json:"middle_name"`
	ClassID    int       `json:"class_id"`
	ParentID   int       `json:"parent_id"`
}

type SuccessResponse struct {
	IsSuccess string `json:"is_success"`
}
