package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santiagoh1997/challenge/controllers"
	"github.com/santiagoh1997/challenge/utils/apierrors"
)

// MapURLs mapea cada endpoint con la función correspondiente en el controlador
func MapURLs(r *mux.Router, sc controllers.SucursalController) {
	// swagger:route GET /sucursales sucursal getNearest
	// Toma una latitud y una longitud, devuelve la sucursal más cercana a ese punto
	// responses:
	// 	200: sucursal
	// 	400: badRequestErrorGeneric
	// 	404: notFoundError
	r.HandleFunc("/sucursales", sc.GetNearest).Methods(http.MethodGet)
	// swagger:route GET /sucursales/:idSucursal sucursal getByID
	// Devuelve una sucursal dado un id
	// responses:
	// 	200: sucursal
	// 	400: badRequestErrorID
	// 	404: notFoundError
	r.HandleFunc("/sucursales/{idSucursal}", sc.GetByID).Methods(http.MethodGet)
	// swagger:route POST /sucursales sucursal createSucursal
	// Crea una sucursal en la base de datos
	// parameters:
	// + name: sucursal
	//   in: body
	//   type: createSucursal
	// responses:
	// 	201: sucursal
	// 	400: badRequestErrorWithFields
	r.HandleFunc("/sucursales", sc.Create).Methods(http.MethodPost)

	// NotFound Handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := apierrors.NewNotFound("Endpoint no encontrado")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode())
		json.NewEncoder(w).Encode(err.Parse())
	})
}
