package controller

import (
	"fiber-crud/app/address/model"
	"fiber-crud/app/address/service"

	"github.com/gofiber/fiber/v2"
)

type controller struct {
	service service.IService
}

func NewController(service service.IService) controller {
	return controller{service: service}
}

func (ct controller) Inquiry(ctx *fiber.Ctx) error {
	req := model.Address{AddressId: ctx.Params("id")}
	return ctx.JSON(ct.service.Inquiry(ctx, &req))
}

// func (ct controller) Create(ctx *fiber.Ctx) error {
// 	req := model.Address{}
// 	if err := ctx.BodyParser(&req); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(model.Create_Response{Message: err.Error()})
// 	}
// 	return ctx.JSON(ct.service.Create(ctx, &req))
// }

// func (ct controller) Update(ctx *fiber.Ctx) error {
// 	req := model.Address{}
// 	if err := ctx.BodyParser(&req); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(model.Create_Response{Message: err.Error()})
// 	}
// 	return ctx.JSON(ct.service.Update(ctx, &req))
// }

// func (ct controller) Delete(ctx *fiber.Ctx) error {
// 	req := model.Address{AddressId: ctx.Params("id")}
// 	return ctx.JSON(ct.service.Delete(ctx, &req))
// }

// func (ct controller) TestDecimal(ctx *fiber.Ctx) error {
// 	req := model.TestDecimal_Request{}
// 	if err := ctx.BodyParser(&req); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(model.Create_Response{Message: err.Error()})
// 	}
// 	return ctx.JSON(ct.service.TestDecimal(ctx, &req))
// }
