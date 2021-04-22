package validators

import (
	"strconv"

	"github.com/Confialink/wallet-settings/internal/db/models"
)

type AutoLogout struct {
	Status            string `binding:"oneof=yes no"`
	Timeout           int    `binding:""`
	Padding           int    `binding:""`
	InactivityMessage string `binding:"required"`
	Message           string `binding:"required_with=Padding"`
}

type AutoLogoutValidator struct {
	structValidator Validator
}

func NewAutoLogoutValidator(structValidator Validator) *AutoLogoutValidator {
	return &AutoLogoutValidator{structValidator: structValidator}
}

func (s *AutoLogoutValidator) Validate(data []*models.Config) *ValidationError {
	input := &AutoLogout{}
	structName := "AutoLogout"
	validationError := &ValidationError{Indexes: make(map[string]int)}
	for i, pair := range data {
		switch pair.Path {
		case models.PathAutoLogoutStatus:
			input.Status = pair.Value
			validationError.Indexes[structName+".Status"] = i
		case models.PathAutoLogoutTimeout:
			input.Timeout, _ = strconv.Atoi(pair.Value)
			validationError.Indexes[structName+".Timeout"] = i
		case models.PathAutoLogoutPadding:
			input.Padding, _ = strconv.Atoi(pair.Value)
			validationError.Indexes[structName+".Padding"] = i
		case models.PathAutoLogoutInactivityMessage:
			input.InactivityMessage = pair.Value
			validationError.Indexes[structName+".InactivityMessage"] = i
		case models.PathAutoLogoutMessage:
			input.Message = pair.Value
			validationError.Indexes[structName+".Message"] = i
		}
	}

	if input.Status != "yes" {
		return nil
	}

	err := s.structValidator.Struct(input)
	if err == nil {
		return nil
	}

	validationError.Err = err

	return validationError
}
