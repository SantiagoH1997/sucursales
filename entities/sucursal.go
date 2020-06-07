package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Sucursal es el struct que define la entidad Sucursal y sus propiedades
type Sucursal struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Direccion string             `json:"direccion" bson:"direccion" validate:"required"`
	Latitud   float64            `json:"latitud" bson:"latitud" validate:"required,min=-90,max=90"`
	Longitud  float64            `json:"longitud" bson:"longitud" validate:"required,min=-180,max=180"`
	Location  Location           `json:"-" bson:"location"`
}

// Location es un campo de tipo GeoJSON
type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
