package message

import (
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

func (c *controller) GetMessageByEmail(ctx *fiber.Ctx) error {
	//TODO implement me
	return nil
}
