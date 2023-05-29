package user

import (
	"context"
	"example.com/m/utils"
	"os"
	"strconv"
	"time"
)

type service struct {
	repository IRepository
	timeout    time.Duration
}

func NewService(userRepo IRepository) IService {
	return &service{
		repository: userRepo,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (service *service) SignIn(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	user, err := service.repository.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return nil, err
	}

	err = utils.ComparePass(user.Password, req.Password)
	if err != nil {
		return nil, err
	}

	token := utils.GenerateToken(strconv.FormatInt(user.ID, 10))
	tokenString, errToken := token.SignedString([]byte(os.Getenv("SECRET")))
	if errToken != nil {
		return nil, errToken
	}
	return &LoginUserRes{
		ID:          strconv.FormatInt(user.ID, 10),
		AccessToken: tokenString,
		Username:    user.Username,
	}, nil
}

func (service *service) SingUp(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	hash, _ := utils.GenPassword(req.Password)
	user, err := service.repository.CreateUser(ctx, &User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hash),
	})
	if err != nil {
		return nil, err
	}
	return &CreateUserRes{
		Username: user.Username,
		ID:       strconv.FormatInt(user.ID, 10),
		Email:    user.Email,
	}, nil
}
