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
	response := ct.service.Inquiry(ctx, &req)
	return ctx.JSON(response)
}
