package repositories_test

import (
	"net/http"
	"testing"

	"github.com/santiagoh1997/challenge/entities"
	"github.com/santiagoh1997/challenge/repositories"
	"github.com/santiagoh1997/challenge/testdata"
	"github.com/santiagoh1997/challenge/testutils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func TestNewSucursalRepository(t *testing.T) {

	sr := repositories.NewSucursalRepository(nil, nil)
	if sr == nil {
		t.Fatalf("NewSucursalRepository want *repositories.Sucursal, got %v", sr)
	}
}

func TestGetByID(t *testing.T) {
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)

	sr := repositories.NewSucursalRepository(db, nil)

	// Success case
	t.Run("Success", func(t *testing.T) {
		sucursal, err := sr.GetByID(testdata.TestSucursalID)
		if err != nil {
			t.Fatalf("GetByID err = %v, want %v", err, nil)
		}

		got := sucursal.Direccion
		want := testdata.TestSucursal.Direccion
		if got != want {
			t.Errorf("GetByID Direccion = %s, want %s", got, want)
		}
	})

	// Fail case
	t.Run("Fail", func(t *testing.T) {
		sucursal, err := sr.GetByID(primitive.NewObjectID())
		if sucursal != nil {
			t.Errorf("GetByID Sucursal = %v, want %v", sucursal, nil)
		}
		if err == nil {
			t.Fatalf("GetByID err = %v, want error", err)
		}

		got := err.StatusCode()
		want := http.StatusNotFound
		if got != want {
			t.Errorf("GetByID StatusCode = %v, want %v", got, want)
		}
	})
}

func TestGetNearest(t *testing.T) {
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)

	sr := repositories.NewSucursalRepository(db, nil)

	// Success case
	t.Run("Success", func(t *testing.T) {
		sucursal, err := sr.GetNearest(testdata.TestSucursal.Location)
		if err != nil {
			t.Fatalf("GetNearest err = %v, want %v", err, nil)
		}

		got := sucursal.Direccion
		want := testdata.TestSucursal.Direccion
		if got != want {
			t.Errorf("GetNearest Direccion = %s, want %s", got, want)
		}
	})

	// Fail case
	t.Run("Fail", func(t *testing.T) {
		testLocation := entities.Location{
			Type:        "Point",
			Coordinates: []float64{-10, 10},
		}

		sucursal, err := sr.GetNearest(testLocation)
		if sucursal != nil {
			t.Errorf("GetNearest Sucursal = %v, want %v", sucursal, nil)
		}
		if err == nil {
			t.Fatalf("GetNearest err = %v, want error", err)
		}

		got := err.StatusCode()
		want := http.StatusNotFound
		if got != want {
			t.Errorf("GetNearest StatusCode = %v, want %v", got, want)
		}
	})
}

func TestCreate(t *testing.T) {
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)

	sr := repositories.NewSucursalRepository(db, nil)

	sucursal := entities.Sucursal{
		Direccion: "Dirección random",
		Latitud:   -10,
		Longitud:  -5,
		Location: entities.Location{
			Type:        "Point",
			Coordinates: []float64{-5, -10},
		},
	}

	apiErr := sr.Create(&sucursal)
	if err != nil {
		t.Fatalf("Create err = %v, want %v", apiErr, nil)
	}
	got := sucursal.ID.Hex()

	if got == primitive.NilObjectID.Hex() {
		t.Fatalf("Create ID got = %v, want ObjectID", got)
	}

}
