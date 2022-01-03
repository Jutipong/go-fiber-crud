package service

import (
	"fiber-crud/app/address/model"
	"fiber-crud/app/address/repository"
	"testing"

	"github.com/gofiber/fiber/v2"
	"gotest.tools/v3/assert"
)

func Test_Inquiry_NotFound(t *testing.T) {
	t.Run("Inquiry NotFound", func(t *testing.T) {
		//repo
		repoMock := repository.NewRepositoryMock()
		repoMock.On("Inquiry").Return([]model.Address{}, fiber.NewError(fiber.StatusBadRequest, "err"))

		//service
		service := NewService(repoMock)

		// //Act
		resule := service.Inquiry(&fiber.Ctx{}, &model.Address{AddressId: "1"})
		assert.Equal(t, false, resule.Status)
		assert.Equal(t, "err", resule.Message)
	})
}

func Test_Inquiry_Found(t *testing.T) {

	t.Run("Inquiry Found", func(t *testing.T) {
		//repo
		repoMock := repository.NewRepositoryMock()
		repoMock.On("Inquiry").Return([]model.Address{
			{AddressId: "1", IsActive: true},
		}, nil)

		//service
		service := NewService(repoMock)

		// //Act
		resule := service.Inquiry(&fiber.Ctx{}, &model.Address{AddressId: "1"})
		assert.Equal(t, true, resule.Status)
		assert.Equal(t, 1, len(*resule.Datas))
	})
}
