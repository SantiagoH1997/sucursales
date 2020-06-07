package services

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/santiagoh1997/challenge/entities"
	"github.com/santiagoh1997/challenge/repositories"
	"github.com/santiagoh1997/challenge/utils/apierrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

}

// SucursalService es el servicio de las sucursales
type SucursalService interface {
	Validate(s *entities.Sucursal) apierrors.APIError
	GetByID(id primitive.ObjectID) (*entities.Sucursal, apierrors.APIError)
	GetNearest(lat, lon float64) (*entities.Sucursal, apierrors.APIError)
	Create(s *entities.Sucursal) apierrors.APIError
}

type service struct {
	repo   repositories.SucursalRepository
	logger *zap.SugaredLogger
}

// NewSucursalService devuelve un servicio de sucursales
func NewSucursalService(r repositories.SucursalRepository, l *zap.SugaredLogger) SucursalService {
	return &service{r, l}
}

// Validate valida que cada campo de una Sucursal tenga un valor correcto
func (svc *service) Validate(s *entities.Sucursal) apierrors.APIError {
	s.Direccion = strings.TrimSpace(s.Direccion)

	var fields []apierrors.ErrorField
	err := validate.Struct(s)
	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			switch fieldError.Tag() {
			case "required":
				fields = append(fields, apierrors.ErrorField{Field: fieldError.Field(), Error: "Campo requerido"})
			case "min":
			case "max":
				switch fieldError.Field() {
				case "lat":
					fields = append(fields, apierrors.ErrorField{Field: fieldError.Field(), Error: "La latitud debe ser un número entre -90 y 90"})
				case "long":
					fields = append(fields, apierrors.ErrorField{Field: fieldError.Field(), Error: "La longitud debe ser un número entre -180 y 180"})
				}
			default:
				fields = append(fields, apierrors.ErrorField{Field: fieldError.Field(), Error: "Error"})
			}
		}
		return apierrors.NewValidationError("Error validando la sucursal", fields)
	}
	return nil
}

func (svc *service) GetByID(id primitive.ObjectID) (*entities.Sucursal, apierrors.APIError) {
	return svc.repo.GetByID(id)
}

func (svc *service) GetNearest(lat, lon float64) (*entities.Sucursal, apierrors.APIError) {
	location := entities.Location{
		Type:        "Point",
		Coordinates: []float64{lon, lat},
	}
	return svc.repo.GetNearest(location)
}

// Create valida una sucursal y la crea mediante su repositorio
func (svc *service) Create(s *entities.Sucursal) apierrors.APIError {
	err := svc.Validate(s)
	if err != nil {
		return err
	}
	return svc.repo.Create(s)
}
