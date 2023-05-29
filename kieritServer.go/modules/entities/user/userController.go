package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	Service IService
}

func NewController(service IService) IController {
	return &controller{
		Service: service,
	}
}
func (ctrl *controller) Login(ctx *fiber.Ctx) error {
	var req = new(LoginUserReq)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}
	response, err := ctrl.Service.SignIn(context.Background(), req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   response,
		"status": true,
	})
}

func (ctrl *controller) Register(ctx *fiber.Ctx) error {
	var user = new(User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}
	response, err := ctrl.Service.SingUp(context.Background(), &CreateUserReq{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":   response,
		"status": true,
	})
}
