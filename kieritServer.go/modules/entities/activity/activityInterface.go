package activity

import "context"

type IRepo interface {
	Save(ctx context.Context, activity *Activity) (*Activity, error)
	GetByUsername(ctx context.Context, username string) ([]Activity, error)
}

type IService interface {
	GetUserHistory(ctx context.Context, username string) ([]Activity, error)
}
