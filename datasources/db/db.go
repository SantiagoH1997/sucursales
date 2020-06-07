package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	collectionName = "sucursales"
	indexKey       = "location_name"
)

// Open abre una conexión a la base de datos
// Devuelve una *mongo.Database, una función para cerrar la conexión, y un error
func Open(mongoURI, dbName string) (*mongo.Database, func(ctx context.Context) error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}
	db := client.Database(dbName)
	indexOpts := options.CreateIndexes().SetMaxTime(time.Second * 10)
	pointIndexModel := mongo.IndexModel{
		Options: options.Index().SetBackground(true),
		Keys:    bsonx.MDoc{"location": bsonx.String("2dsphere")},
	}
	pointIndexes := db.Collection(collectionName).Indexes()
	_, err = pointIndexes.CreateOne(
		ctx,
		pointIndexModel,
		indexOpts,
	)
	if err != nil {
		return nil, nil, err
	}
	return db, client.Disconnect, nil
}
