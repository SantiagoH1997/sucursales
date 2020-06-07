package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/santiagoh1997/challenge/entities"
	"github.com/santiagoh1997/challenge/services"
	"github.com/santiagoh1997/challenge/utils/apierrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SucursalController es el controldor para sucursales
type SucursalController interface {
	GetNearest(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	parseLatAndLong(query url.Values) (float64, float64, apierrors.APIError)
}

type controller struct {
	service services.SucursalService
}

// NewSucursalController devuelve un controlador de sucursales
func NewSucursalController(ss services.SucursalService) SucursalController {
	return &controller{ss}
}

// GetNearest toma una latitud y longitud de la query string
//  y devuelve la sucursal más cercana dentro de un área determinada
func (c *controller) GetNearest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	lat, long, apiErr := c.parseLatAndLong(query)
	if apiErr != nil {
		w.WriteHeader(apiErr.StatusCode())
		json.NewEncoder(w).Encode(apiErr.Parse())
		return
	}

	s, apiErr := c.service.GetNearest(lat, long)
	if apiErr != nil {
		w.WriteHeader(apiErr.StatusCode())
		json.NewEncoder(w).Encode(apiErr.Parse())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
	return
}

// GetByID busca una sucursal por un id enviado en la URL
func (c *controller) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["idSucursal"]
	if id == "" {
		apiErr := apierrors.NewBadRequest("Debe proporcionarse un id")
		w.WriteHeader(apiErr.StatusCode())
		json.NewEncoder(w).Encode(apiErr.Parse())
		return
	}

	idSucursal, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		apiErr := apierrors.NewBadRequest("El id proporcionado no es válido")
		w.WriteHeader(apiErr.StatusCode())
		json.NewEncoder(w).Encode(apiErr.Parse())
		return
	}

	sucursal, apiErr := c.service.GetByID(idSucursal)
	if err != nil {
		w.WriteHeader(apiErr.StatusCode())
		json.NewEncoder(w).Encode(apiErr.Parse())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sucursal)

}

// Create crea una sucursal a partir de datos enviados en el cuerpo de la request
func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	var sucursal entities.Sucursal
	err := json.NewDecoder(r.Body).Decode(&sucursal)
	if err != nil {
		apiErr := apierrors.NewBadRequest(err.Error())
		w.WriteHeader(apiErr.StatusCode())
		json.NewEncoder(w).Encode(apiErr.Parse())
		return
	}
	apiErr := c.service.Create(&sucursal)
	if apiErr != nil {
		w.WriteHeader(apiErr.StatusCode())
		json.NewEncoder(w).Encode(apiErr.Parse())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&sucursal)
}

// parseLatAndLong valida y devuelve la latitud y longitud enviada via query string
func (c *controller) parseLatAndLong(query url.Values) (float64, float64, apierrors.APIError) {
	lat, err := strconv.ParseFloat(query.Get("lat"), 64)
	if err != nil || lat < -90 || lat > 90 {
		return 0, 0, apierrors.NewBadRequest("La latitud debe ser un número entre -90 y 90")
	}
	lon, err := strconv.ParseFloat(query.Get("lon"), 64)
	if err != nil || lon < -180 || lon > 180 {
		return 0, 0, apierrors.NewBadRequest("La longitud debe ser un número entre -180 y 180")
	}
	return lat, lon, nil
}
