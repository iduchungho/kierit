package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

var root *mongo.Client
var lock = &sync.Mutex{}

type Database struct {
	conn *mongo.Client
}

func CreateConnectDatabase() *Database {
	return &Database{
		conn: newConn(),
	}
}

func (db *Database) GetConn() *mongo.Client {
	return db.conn
}

func (db *Database) NewConn() {
	db.conn = newConn()
}

func (db *Database) Close(context context.Context) {
	err := db.conn.Disconnect(context)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb disconnected")
}
