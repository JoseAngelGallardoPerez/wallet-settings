package validators

import (
	validatorPkg "github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/Confialink/wallet-settings/internal/db/models"
)

var _ = Describe("Auto logout", func() {
	Context("Validate", func() {
		When("there is Status field is not `yes`", func() {
			It("should not return an error", func() {
				data := []*models.Config{
					{
						Path:  models.PathAutoLogoutStatus,
						Value: "no",
					},
				}
				mock := &MockValidator{}
				service := NewAutoLogoutValidator(mock)
				Expect(service.Validate(data)).Should(BeNil())
			})
		})

		When("struct validator does not return validation errors", func() {
			It("should not return an error", func() {
				data := []*models.Config{
					{
						Path:  models.PathAutoLogoutStatus,
						Value: "yes",
					},
				}
				validator := &MockValidator{}
				validator.On("Struct", mock.Anything).Return(nil)
				service := NewAutoLogoutValidator(validator)
				Expect(service.Validate(data)).Should(BeNil())
			})
		})

		When("struct validator returns validation errors", func() {
			It("should return an error", func() {
				data := []*models.Config{
					{
						Path:  models.PathAutoLogoutStatus,
						Value: "yes",
					},
				}
				validationErrs := make(validatorPkg.ValidationErrors, 0)
				validator := &MockValidator{}
				validator.On("Struct", mock.Anything).Return(validationErrs)
				service := NewAutoLogoutValidator(validator)
				err := service.Validate(data)
				Expect(err).ShouldNot(BeNil())
				Expect(err.Err).To(BeEquivalentTo(validationErrs))
			})
		})
	})
})
