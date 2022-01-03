package service

import (
	"fiber-crud/app/address/model"
	"fiber-crud/app/address/repository"
	"fiber-crud/pkg/utils"

	// "fiber-crud/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type IService interface {
	Inquiry(ctx *fiber.Ctx, req *model.Address) model.Inquiry_Response
	Create(ctx *fiber.Ctx, req *model.Address) model.Create_Response
	Update(ctx *fiber.Ctx, req *model.Address) model.Update_Response
	Delete(ctx *fiber.Ctx, req *model.Address) model.Delete_Response
	TestDecimal(ctx *fiber.Ctx, req *model.TestDecimal_Request) model.TestDecimal_Response
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return service{repo: repo}
}

func (s service) Inquiry(ctx *fiber.Ctx, req *model.Address) model.Inquiry_Response {
	// Validation
	if err := req.Inquiry_Validation(ctx); err != nil {
		return model.Inquiry_Response{Message: err.Error()}
	}

	// Call repository
	result, err := s.repo.Inquiry(req)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Inquiry_Response{Message: err.Error()}
	}

	return model.Inquiry_Response{Status: true, Datas: &result}
}

func (s service) Create(ctx *fiber.Ctx, req *model.Address) (resutl model.Create_Response) {
	// Validation
	if err := req.Create_Validation(ctx); err != nil {
		return model.Create_Response{Message: err.Error()}
	}

	// Init user request
	userInfo := utils.GetUserInfo(ctx)
	req.AddressId = uuid.New().String()
	req.CreateBy = userInfo.UserId

	// Call repository
	if err := s.repo.Create(req); err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Create_Response{Message: err.Error()}
	}

	return model.Create_Response{Status: true, Datas: req}
}

func (s service) Update(ctx *fiber.Ctx, req *model.Address) (resutl model.Update_Response) {
	// Validation
	if err := req.Update_Validation(ctx); err != nil {
		return model.Update_Response{Message: err.Error()}
	}

	// Init user request
	userInfo := utils.GetUserInfo(ctx)
	req.UpdateBy = userInfo.UserId

	// Call repository
	if err := s.repo.Update(req); err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Update_Response{Message: err.Error()}
	}

	return model.Update_Response{Status: true, Datas: req}
}

func (s service) Delete(ctx *fiber.Ctx, req *model.Address) (resutl model.Delete_Response) {
	// Validation
	if err := req.Delete_Validation(ctx); err != nil {
		return model.Delete_Response{Message: err.Error()}
	}

	// Init user request
	userInfo := utils.GetUserInfo(ctx)
	req.UpdateBy = userInfo.UserId

	// Call repository
	if err := s.repo.Delete(req); err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return model.Delete_Response{Message: err.Error()}
	}

	return model.Delete_Response{Status: true}
}

func (s service) TestDecimal(ctx *fiber.Ctx, req *model.TestDecimal_Request) model.TestDecimal_Response {
	total := req.Number1.Add(req.Number2)
	res := model.TestDecimal_Response{
		Number1: req.Number1,
		Number2: req.Number2,
		Total:   decimal.NullDecimal{Decimal: total, Valid: true},
	}
	res.JsonNumber()
	return res
}
