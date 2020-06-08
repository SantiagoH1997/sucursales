// Package doc API Sucursales.
//
// Permite crear sucursales, buscar sucursales por id o por cercanía
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//	License: MIT http://opensource.org/licenses/MIT
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package doc

// Sucursal
// swagger:model sucursal
// in: body
type sucursal struct {
	// ID de la sucursal
	//
	// example: 5edb503e14e865fa8ac0cf4b
	ID string `json:"id"`
	// Dirección de la sucursal
	//
	// example: Pres. Juan Domingo Perón 4739
	Direccion string `json:"direccion"`
	// Latitud de la sucursal
	//
	// example: -34.7625601
	Latitud float64 `json:"latitud"`
	// Longitud de la sucursal
	//
	// example: -58.2192142
	Longitud float64 `json:"longitud"`
}

// swagger:model createSucursal
type createSucursal struct {
	// Dirección de la sucursal
	// example: Pres. Juan Domingo Perón 4739
	// required: true
	Direccion string `json:"direccion"`
	// Latitud de la sucursal
	// example: -34.7625601
	// required: true
	Latitud float64 `json:"latitud"`
	// Longitud de la sucursal
	// example: -58.2192142
	// required: true
	Longitud float64 `json:"longitud"`
}

// swagger:parameters getByID
type getByIDParameters struct {
	// id
	// example: 5edb503e14e865fa8ac0cf4b
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters getNearest
type getNearestParameters struct {
	// lat
	// min: -90
	// max: 90
	// example: -34.7625601
	// in: query
	// required: true
	Lat float64 `json:"lat"`
	// lon
	// min: -180
	// max: 180
	// example: -58.2192142
	// in: query
	// required: true
	Lon float64 `json:"lon"`
}

// notFoundError
// swagger:response notFoundError
// in: body
type notFoundError struct {
	// 404
	StatusCode int `json:"status_code"`
	// Mensaje del error
	// example: No se encontró la sucursal
	Message string `json:"message"`
}

// badRequestError
// swagger:response badRequestErrorID
// in: body
type badRequestErrorID struct {
	// 400
	StatusCode int `json:"status_code"`
	// Mensaje del error
	// example: El id no es válido
	Message string `json:"message"`
}

// badRequestError
// swagger:response badRequestErrorGeneric
// in: body
type badRequestErrorGeneric struct {
	// 400
	StatusCode int `json:"status_code"`
	// Mensaje del error
	// example:  La latitud debe ser un número entre -90 y 90
	Message string `json:"message"`
}

// badRequestError
// swagger:response badRequestErrorWithFields
// in: body
type badRequestErrorWithFields struct {
	// 400
	StatusCode int `json:"status_code"`
	// Mensaje del error
	// example: La latitud debe ser un número entre -90 y 90
	Message string `json:"message"`
	// Campos donde sucede el error (en caso de ser necesario)
	Fields []errorField `json:"fields"`
}

type errorField struct {
	// Campo donde sucede el error (en caso de ser necesario)
	Field string `json:"field"`
	// Mensaje del error en el campo
	Error string `json:"error"`
}
