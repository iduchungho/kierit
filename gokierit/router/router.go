package router

import (
	"example.com/m/modules/entities/message"
	"example.com/m/modules/entities/user"
	"example.com/m/modules/middleware"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	handle *fiber.App
}

func New(app *fiber.App) *Router {
	return &Router{
		handle: app,
	}
}

func (r *Router) UserRouter(controller user.IController) {
	r.handle.Post("/api/signup", controller.Register)
	r.handle.Post("/api/login", controller.Login)
}

func (r *Router) MessageRouter(controller message.IController) {
	r.handle.Get("/api/message/get", middleware.RequiredUser, controller.GetMessageByEmail)
}
