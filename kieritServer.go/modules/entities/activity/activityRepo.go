package activity

import "go.mongodb.org/mongo-driver/mongo"

type ActivityRepo struct {
	Conn *mongo.Client
}

func NewMessageRepo(conn *mongo.Client) *ActivityRepo {
	return &ActivityRepo{
		Conn: conn,
	}
}
