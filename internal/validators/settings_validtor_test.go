package validators

import (
	"errors"

	validatorPkg "github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Confialink/wallet-settings/internal/db/models"
)

var _ = Describe("Setting Validator", func() {
	Context("Add", func() {
		When("successfully added new validator", func() {
			It("should change quantity of validators", func() {
				mock := &MockSettingValidator{}
				validator := NewSettingsValidator()
				Expect(len(validator.validators)).Should(Equal(0))
				validator.Add(mock)
				Expect(len(validator.validators)).Should(Equal(1))
			})
		})
	})

	Context("Validate", func() {
		When("there is successful validation without errors", func() {
			It("should not return an error", func() {
				data := []*models.Config{
					{
						Path:  models.PathAutoLogoutInactivityMessage,
						Value: "random_value",
					},
				}
				mock := &MockSettingValidator{}
				mock.On("Validate", data).Return(nil)
				validator := NewSettingsValidator()
				validator.Add(mock)

				Expect(validator.Validate(data)).ShouldNot(HaveOccurred())
			})
		})

		When("a particular validator return an unknown error", func() {
			It("should return an error", func() {
				data := []*models.Config{
					{
						Path:  models.PathAutoLogoutInactivityMessage,
						Value: "random_value",
					},
				}
				mock := &MockSettingValidator{}
				mock.On("Validate", data).Return(errors.New("random error"))
				validator := NewSettingsValidator()
				validator.Add(mock)

				err := validator.Validate(data)
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("cannot convert validation errors"))
			})
		})

		When("a particular validator return validation errors", func() {
			It("should return validation errors", func() {
				data := []*models.Config{
					{
						Path:  models.PathAutoLogoutInactivityMessage,
						Value: "random_value",
					},
				}
				mock := &MockSettingValidator{}
				validationErr := &ValidationError{
					Err:     make(validatorPkg.ValidationErrors, 0),
					Indexes: make(map[string]int),
				}
				mock.On("Validate", data).Return(validationErr)
				validator := NewSettingsValidator()
				validator.Add(mock)

				err := validator.Validate(data)
				Expect(err).Should(HaveOccurred())
				Expect(err).To(BeEquivalentTo(make(validatorPkg.ValidationErrors, 0)))
			})
		})
	})
})
