package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type IRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type IService interface {
	SignIn(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
	SingUp(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
}

type IController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
