package database

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/msantosfelipe/money-controller/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *Database
var once sync.Once

type Database struct {
	client *mongo.Client
}

// GetInstance -- Creates a new dabatase instance
func GetInstance() *Database {
	once.Do(func() {
		instance = newInstance()
	})
	return instance
}

func newInstance() *Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(config.ENV.DB_URL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		client: client,
	}
}
