package message

import (
	"context"
	"example.com/m/database"
)

type repository struct {
	Conn *database.Database
}

func NewRepo(conn *database.Database) IRepository {
	return &repository{
		Conn: conn,
	}
}

func (r *repository) Save(ctx context.Context, message *Message) (*Message, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetMessageByEmail(ctx context.Context, username string) ([]Message, error) {
	//TODO implement me
	panic("implement me")
}
