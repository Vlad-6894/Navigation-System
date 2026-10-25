package auth_domains

import (
	"fmt"
	pkg_http_errors "navigation-pkg/errors/http"
)

type Names struct {
	FirstName  string
	Surname    string
	MiddleName string
}

func NewNames(
	firstName string,
	surname string,
	middleName string,
) Names {
	return Names{
		FirstName:  firstName,
		Surname:    surname,
		MiddleName: middleName,
	}
}

func (n Names) Validate() error {
	if len([]rune(n.FirstName)) < 1 || len([]rune(n.FirstName)) > 50 {
		return fmt.Errorf("fail to validate first name because it must be more 1 and less 50 : %w", pkg_http_errors.ErrInvalidArgument)
	}

	if len([]rune(n.Surname)) < 1 || len([]rune(n.Surname)) > 50 {
		return fmt.Errorf("fail to validate surnamename because it must be more 1 and less 50 : %w", pkg_http_errors.ErrInvalidArgument)
	}

	if len([]rune(n.MiddleName)) < 1 || len([]rune(n.MiddleName)) > 50 {
		return fmt.Errorf("fail to validate middle name because it must be more 1 and less 50 : %w", pkg_http_errors.ErrInvalidArgument)
	}

	return nil
}
