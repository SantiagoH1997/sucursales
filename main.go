package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	docsMiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/santiagoh1997/challenge/controllers"
	"github.com/santiagoh1997/challenge/datasources/db"
	"github.com/santiagoh1997/challenge/logger"
	"github.com/santiagoh1997/challenge/middleware"
	"github.com/santiagoh1997/challenge/repositories"
	"github.com/santiagoh1997/challenge/routes"
	"github.com/santiagoh1997/challenge/services"
)

func main() {
	// Creando logger
	l := logger.NewLogger()
	defer l.Sync()

	// Conectando a DB
	mongoURI := fmt.Sprintf("mongodb://%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	mongoDB, close, err := db.Open(mongoURI, os.Getenv("DB_NAME"))
	if err != nil {
		l.Error("Error conectando a la DB")
		panic(err)
	}
	defer close(context.Background())
	l.Info("Conectado a la DB")

	// Configurando repositorio
	sr := repositories.NewSucursalRepository(mongoDB, l)

	// Configurando servicio
	ss := services.NewSucursalService(sr, l)

	// Configurando controlador
	sc := controllers.NewSucursalController(ss)

	// Mux setup
	r := mux.NewRouter().StrictSlash(true)

	routes.MapURLs(r, sc)

	// Middleware
	r.Use(middleware.Logging(l))
	r.Use(middleware.ContentTypeJSON)
	// Swagger docs
	if os.Getenv("ENV") == "dev" {
		opts := docsMiddleware.RedocOpts{SpecURL: "/swagger/swagger.yaml"}
		sh := docsMiddleware.Redoc(opts, nil)
		r.Handle("/docs", sh)
		r.Handle("/swagger/swagger.yaml", http.FileServer(http.Dir("./")))
	}

	l.Fatal(http.ListenAndServe(":8080", r))
}
