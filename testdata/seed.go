package testdata

import (
	"github.com/santiagoh1997/challenge/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// TestSucursales son Sucursales a ser insertadas en la DB de prueba
	TestSucursales []entities.Sucursal
	// TestSucursalID es el ID de la sucursal a usarse en los tests
	TestSucursalID primitive.ObjectID
	// TestSucursal es una Sucursal a utilizarse durante los tests
	TestSucursal = entities.Sucursal{
		Direccion: "Gral. Guemes 897",
		Latitud:   -34.6764802,
		Longitud:  -58.3696455,
		Location: entities.Location{
			Type:        "Point",
			Coordinates: []float64{-58.3696455, -34.6764802},
		},
	}
)

func init() {
	TestSucursalID = primitive.NewObjectID()
	TestSucursal.ID = TestSucursalID
	TestSucursales = []entities.Sucursal{
		TestSucursal,
		{
			Direccion: "Pres. Juan Domingo Perón 4739",
			Latitud:   -34.7625601,
			Longitud:  -58.2192142,
			Location: entities.Location{
				Type:        "Point",
				Coordinates: []float64{-58.2192142, -34.7625601},
			},
		},
		{
			Direccion: "Paraná 3822",
			Latitud:   -34.5094551,
			Longitud:  -58.5314937,
			Location: entities.Location{
				Type:        "Point",
				Coordinates: []float64{-58.5314937, -34.5094551},
			},
		},
		{
			Direccion: "Antonio Saenz 2041",
			Latitud:   -34.5091195,
			Longitud:  -58.5753022,
			Location: entities.Location{
				Type:        "Point",
				Coordinates: []float64{-58.5753022, -34.5091195},
			},
		},
	}
}
