package database

import (
	"context"

	"github.com/msantosfelipe/money-controller/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *Database) getCollection(collection string) *mongo.Collection {
	return db.client.Database(config.ENV.Database).Collection(collection)
}

// FindAll --
func (db *Database) FindAll(collection string) (*mongo.Cursor, error) {
	return db.Find(collection, bson.D{})
}

// Find --
func (db *Database) Find(collection string, filter interface{}) (*mongo.Cursor, error) {
	cur, err := db.getCollection(collection).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

// FindOne --
func (db *Database) FindOne(collection string, filter interface{}) *mongo.SingleResult {
	singleResult := db.getCollection(collection).FindOne(context.Background(), filter)
	return singleResult
}

// InsertOne -- Inserts one iten to a collection
func (db *Database) InsertOne(collection string, data interface{}) (*mongo.InsertOneResult, error) {
	result, err := db.getCollection(collection).InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateByID
func (db *Database) UpdateByID(collection string, id primitive.ObjectID, data interface{}) (*mongo.UpdateResult, error) {
	result, err := db.getCollection(collection).UpdateByID(context.Background(), id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
