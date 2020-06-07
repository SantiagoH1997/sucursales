package repositories

import (
	"time"

	"github.com/santiagoh1997/challenge/entities"
	"github.com/santiagoh1997/challenge/utils/apierrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

const (
	timeout        = 5 * time.Second
	collectionName = "sucursales"
	maxDistance    = 20 * 1000
)

// SucursalRepository se comunica con la DB
type SucursalRepository interface {
	GetByID(id primitive.ObjectID) (*entities.Sucursal, apierrors.APIError)
	GetNearest(location entities.Location) (*entities.Sucursal, apierrors.APIError)
	Create(sucursal *entities.Sucursal) apierrors.APIError
}

type repository struct {
	db     *mongo.Database
	logger *zap.SugaredLogger
}

// NewSucursalRepository devuelve unrepositorio de sucursales
// con una conexión a la base de datos
func NewSucursalRepository(db *mongo.Database, l *zap.SugaredLogger) SucursalRepository {
	return &repository{db, l}
}

func (r *repository) GetByID(id primitive.ObjectID) (*entities.Sucursal, apierrors.APIError) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var s entities.Sucursal
	err := r.db.Collection(collectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&s)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apierrors.NewNotFound("No se encontró la sucursal")
		}
		r.logger.Error(err.Error())
		return nil, apierrors.NewInternalServerError("Error buscando sucursal")
	}

	return &s, nil
}

// GetNearest devuelve la Sucursal más cercana dentro de un radio de 20km
// o un error en caso de no haber ninguna.
func (r *repository) GetNearest(location entities.Location) (*entities.Sucursal, apierrors.APIError) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Filtra sucursales a menos de x km de distancia
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry":    location,
				"$maxDistance": maxDistance,
			},
		},
	}

	var sucursal entities.Sucursal
	err := r.db.Collection(collectionName).FindOne(ctx, filter).Decode(&sucursal)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apierrors.NewNotFound("No se encontraron sucursales")
		}
		r.logger.Error(err.Error())
		return nil, apierrors.NewInternalServerError("Error buscando sucursales")
	}
	return &sucursal, nil
}

// Create crea una Sucursal en la base de datos
func (r *repository) Create(sucursal *entities.Sucursal) apierrors.APIError {
	sucursal.ID = primitive.NilObjectID
	sucursal.Location = entities.Location{
		Type:        "Point",
		Coordinates: []float64{sucursal.Longitud, sucursal.Latitud},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	insertRes, err := r.db.Collection(collectionName).InsertOne(ctx, sucursal)
	if err != nil {
		r.logger.Error(err.Error())
		return apierrors.NewInternalServerError("error creating place")
	}
	sucursal.ID = insertRes.InsertedID.(primitive.ObjectID)
	return nil
}
