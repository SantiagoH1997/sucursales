package testutils

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/santiagoh1997/challenge/datasources/db"
	"github.com/santiagoh1997/challenge/testdata"

	// Importa las variables de entorno para testear con conexión a la DB
	_ "github.com/santiagoh1997/challenge/env"
)

var (
	mongoURI   string
	dbName     string
	collection = "weather"
)

func init() {
	mongoURI = fmt.Sprintf("mongodb://%s:%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"))
	dbName = os.Getenv("TEST_DB_NAME")
	collection = os.Getenv("TEST_DB_COLLECTION")
}

// Setup se conecta a la DB e inserta sucursales
// Devuelve una conexión a la DB, una función para cerrar la conexión, y un error
func Setup() (*mongo.Database, func(ctx context.Context) error, error) {
	database, close, err := db.Open(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Error conectándose a MongoDB: %v", err.Error())
	}
	ctx := context.Background()
	if _, err := database.Collection(collection).DeleteMany(ctx, bson.M{}); err != nil {
		log.Fatalf("Error eliminando registros de la DB: %v", err.Error())
	}
	if err := seed(ctx, database); err != nil {
		log.Fatalf("Error insertando registros a la DB: %v", err.Error())
	}
	return database, close, nil
}

func seed(ctx context.Context, db *mongo.Database) error {
	sucursales := testdata.TestSucursales
	sucursalesInterface := make([]interface{}, len(sucursales))
	for i, v := range sucursales {
		sucursalesInterface[i] = v
	}
	_, err := db.Collection(collection).InsertMany(ctx, sucursalesInterface)
	return err
}
