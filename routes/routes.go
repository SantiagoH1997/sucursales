package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santiagoh1997/challenge/controllers"
)

// MapURLs mapea cada endpoint con la función correspondiente en el controlador
func MapURLs(r *mux.Router, sc controllers.SucursalController) {
	// swagger:route GET /sucursales/:idSucursal sucursal getByID
	// Devuelve una sucursal dado un id
	// responses:
	// 	200: weather
	// 	404: notFoundError
	// 	401: badRequestError
	r.HandleFunc("/sucursales", sc.GetNearest).Methods(http.MethodGet)
	r.HandleFunc("/sucursales/{idSucursal}", sc.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/sucursales", sc.Create).Methods(http.MethodPost)
}
