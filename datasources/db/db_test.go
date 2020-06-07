package db_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/santiagoh1997/challenge/datasources/db"
	_ "github.com/santiagoh1997/challenge/env"
)

var (
	mongoURI string
	dbName   string
)

func init() {
	mongoURI = fmt.Sprintf("mongodb://%s:%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"))
	dbName = os.Getenv("TEST_DB_NAME")
}

func TestOpen(t *testing.T) {
	tests := []struct {
		name string
		URI  string
		DB   string
	}{
		{"Success", mongoURI, dbName},
		{"Wrong DB name", mongoURI, ""},
		{"Wrong URI", "", dbName},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, close, err := db.Open(tt.URI, tt.DB)
			if err != nil {
				if tt.URI != "" && tt.DB != "" {
					t.Errorf("Open err = %v, want %v", err, nil)
				}
			} else {
				defer close(context.Background())
			}
		})
	}
}
