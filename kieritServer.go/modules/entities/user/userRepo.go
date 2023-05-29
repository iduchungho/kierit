package user

import (
	"context"
	"example.com/m/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type repository struct {
	db *database.Database
}

func NewRepo(db *database.Database) IRepository {
	return &repository{
		db: db,
	}
}

func (repo *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	collection := repo.db.GetConn().Database("MessengerDB").Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, err
	}
	var result User
	err = collection.FindOne(context.Background(), bson.M{}, options.FindOne().SetSort(bson.M{"id": -1})).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Fatal(err)
	}
	user.ID = result.ID + 1
	_, err = collection.InsertOne(ctx, *user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	collection := repo.db.GetConn().Database("MessengerDB").Collection("users")
	var userRes = new(User)
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(userRes)
	if err != nil {
		return nil, err
	}
	return userRes, nil
}
