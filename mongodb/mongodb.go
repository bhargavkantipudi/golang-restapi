package mongodb

import (
	"context"
	"errors"
	"fmt"
	"sf-heating-process-service/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// DBCon is the connection handle
	DBCon *mongo.Database
)

// DbInit Initiage the db connection
func DbInit(uri string, dbname string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Info.Println("db: couldn't connect to mongo: ", err)
		return nil, fmt.Errorf("db: couldn't connect to mongo: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		logger.Info.Println("DB Error ", err)
		return nil, fmt.Errorf("db: mongo client couldnt connect with background context: %v", err)
	}
	DBCon = client.Database(dbname)
	return DBCon, nil
}

// UpdateDb ...
func UpdateDb(tableName string, filter interface{}, input interface{}) error {
	upd := DBCon.Collection(tableName).FindOneAndUpdate(context.Background(), filter, &input, options.FindOneAndUpdate().SetUpsert(true))
	if upd.Err() != nil {
		return errors.New("Update failed")
	}
	return nil
}

// InsertDb ...
func InsertDb(tableName string, input interface{}) (interface{}, error) {
	item, err := DBCon.Collection(tableName).InsertOne(context.Background(), input)
	if err != nil {
		panic(err)
	}
	return item, err
}

// InsertManyDb ...
func InsertManyDb(tableName string, input []interface{}) (*mongo.InsertManyResult, error) {
	item, err := DBCon.Collection(tableName).InsertMany(context.Background(), input)
	if err != nil {
		panic(err)
	}
	return item, err
}

// FindItemDb ...
func FindItemDb(tableName string, filter interface{}) (*mongo.SingleResult, error) {
	item := DBCon.Collection(tableName).FindOne(context.Background(), filter)
	if item.Err() != nil {
		return nil, errors.New("Item not found")
	}
	return item, nil
}

// FindItemsDb ...
func FindItemsDb(tableName string, filter interface{}) (*mongo.Cursor, error) {
	cur, err := DBCon.Collection(tableName).Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.New("Items not found")
	}
	return cur, nil
}

// InsertIfNotExistInDb ...
func InsertIfNotExistInDb(tableName string, filter interface{}, input interface{}) error {
	_, err := DBCon.Collection(tableName).UpdateOne(context.Background(), filter, &input, options.Update().SetUpsert(true))
	if err != nil {
		return errors.New("Update failed")
	}
	return nil
}
