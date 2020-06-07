package services_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/santiagoh1997/challenge/entities"
	"github.com/santiagoh1997/challenge/repositories"
	"github.com/santiagoh1997/challenge/services"
	"github.com/santiagoh1997/challenge/testdata"
	"github.com/santiagoh1997/challenge/testutils"
)

func TestCreate(t *testing.T) {
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)

	sr := repositories.NewSucursalRepository(db, nil)
	ss := services.NewSucursalService(sr, nil)

	// Success
	sucursales := testdata.TestSucursales
	for _, sucursal := range sucursales {
		t.Run(sucursal.Direccion, func(t *testing.T) {
			apiErr := ss.Create(&sucursal)
			if apiErr != nil {
				t.Errorf("Create err = %v, want %v", apiErr, nil)
			}
		})
	}

	// Fail
	tests := []struct {
		name     string
		sucursal entities.Sucursal
	}{
		{"Campos faltantes", entities.Sucursal{Direccion: "Dirección random"}},
		{"Campos faltantes", entities.Sucursal{Latitud: -10, Longitud: 10}},
		{"Longitud excede el límite", entities.Sucursal{Direccion: "Dirección random", Latitud: -10, Longitud: 190}},
		{"Longitud debajo del límite", entities.Sucursal{Direccion: "Dirección random", Latitud: -10, Longitud: -190}},
		{"Latitud excede el límite", entities.Sucursal{Direccion: "Dirección random", Latitud: 100, Longitud: 10}},
		{"Latitud debajo del límite", entities.Sucursal{Direccion: "Dirección random", Latitud: -100, Longitud: 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiErr := ss.Create(&tt.sucursal)
			if apiErr == nil {
				t.Fatalf("Create got %v, want err", apiErr)
			}

			got := apiErr.StatusCode()
			want := http.StatusBadRequest
			if got != want {
				t.Errorf("Create error.StatusCode =  %v, want %v", got, want)
			}
		})
	}
}
