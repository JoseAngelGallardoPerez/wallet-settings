package validators

import (
	"errors"
	"strconv"

	validatorPkg "github.com/go-playground/validator/v10"
	errPkg "github.com/pkg/errors"

	"github.com/Confialink/wallet-settings/internal/db/models"
)

type SettingsValidator struct {
	validators []SettingValidator
}

func NewSettingsValidator() *SettingsValidator {
	return &SettingsValidator{}
}

type SettingValidator interface {
	Validate(data []*models.Config) *ValidationError
}

// Add adds a new validator to the registry
func (s *SettingsValidator) Add(v SettingValidator) {
	s.validators = append(s.validators, v)
}

// Validate executes all validators
func (s *SettingsValidator) Validate(data []*models.Config) (error error) {
	defer func() {
		if v := recover(); v != nil {
			error = errors.New("cannot convert validation errors. Check validator version")
		}
	}()
	for _, validator := range s.validators {
		err := validator.Validate(data)
		if err == nil {
			continue
		}
		if validationErrors, ok := err.Err.(validatorPkg.ValidationErrors); ok {
			// We should create a new list of errors and change field name
			// Because error response fields should match to request body fields
			newErrs := make(validatorPkg.ValidationErrors, 0, len(validationErrors))
			for _, validationError := range validationErrors {
				idx := err.Indexes[validationError.Namespace()]
				newErr := FieldError{
					tag: validationError.Tag(),
					// "ParentStruct" - first part of name will be cut so we can use random name
					// "Data" - name of structure in original structure.
					// "Value" - name of field in original structure
					ns:          "ParentStruct.Data[" + strconv.Itoa(idx) + "].Value",
					field:       validationError.Field(),
					structfield: validationError.StructField(),
					value:       validationError.Value(),
					param:       validationError.Param(),
					kind:        validationError.Kind(),
					typ:         validationError.Type(),
				}

				newErrs = append(newErrs, newErr)
			}
			return newErrs
		}

		return errPkg.Wrap(err.Err, "cannot convert validation errors. Check validator version")
	}

	return nil
}
