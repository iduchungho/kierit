package message

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type IRepository interface {
	Save(ctx context.Context, message *Message) (*Message, error)
	GetMessageByEmail(ctx context.Context, username string) ([]Message, error)
}

type IService interface {
	SaveMessage(message *Message) (*Message, error)
	GetMessageByEmail(ctx context.Context, username string) ([]Message, error)
}

type IController interface {
	GetMessageByEmail(ctx *fiber.Ctx) error
}
