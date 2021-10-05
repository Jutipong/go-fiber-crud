package service

import (
	"fiber-crud/app/address/model"
	"fiber-crud/app/address/repository"
	"fiber-crud/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type IService interface {
	Inquiry(ctx *fiber.Ctx, req *model.Address) model.Inquiry_Response
	// Create(ctx *fiber.Ctx, req *model.Address) (model.Address, error)
	// Update(ctx *fiber.Ctx, req *model.Address) (model.Address, error)
	// Delete(ctx *fiber.Ctx, req *model.Address) error
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return &service{repo}
}

func (s *service) Inquiry(ctx *fiber.Ctx, req *model.Address) model.Inquiry_Response {
	// Validation
	err := req.Inquiry_Validation(ctx)
	if err != nil {
		return model.Inquiry_Response{Message: err.Error()}
	}

	// Call repository
	result, err := s.repo.Inquiry(req)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Inquiry_Response{Message: err.Error()}
	}

	return model.Inquiry_Response{Status: true, Datas: result}
}
