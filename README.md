# API Sucursales

  
## Instalación

El siguiente comando instala todas las dependencias recursivamente (en caso de querer correr el servicio de manera local).

```bash
go get ./...
```


## Ejecución

La aplicación correrá en http://localhost:8080

```bash
docker-compose build
docker-compose up
```


## Tests

El siguiente comando corre todos los tests de la aplicación (MongoDB debe estar instalado localmente).

```bash
go test ./...
```


## Documentación

La documentacion swagger se puede encontrar en http://localhost:8080/docs


## Funcionalidad:

- Crear Sucursal.
- Buscar Sucursal por ID.
- Buscar Sucursal más cercana a un punto.


### Crear Sucursal
- Endpoint: /sucursales
- Método: POST

Ejemplo de request body:

```jsonc
{
	"direccion": "Pres. Juan Domingo Perón 4739", //string
	"latitud": -34.7625601, //float
	"longitud": -58.2192142, //float
}
```

Ejemplo de response (status code 201):

```jsonc
{
	"id": "5edb503e14e865fa8ac0cf4b", //string
	"direccion": "Pres. Juan Domingo Perón 4739", //string
	"latitud": -34.7625601, //float
	"longitud": -58.2192142, //float
}
```


### Buscar Sucursal por coordenadas
- Endpoint: /sucursales?lat={}&lon={}
- Método: GET

Ejemplo de request:

```
GET /sucursales?lat=-34.7625601&lon=-58.2192142
```

Ejemplo de response (status code 200):

```jsonc
{
	"id": "5edb503e14e865fa8ac0cf4b", //string
	"direccion": "Pres. Juan Domingo Perón 4739", //string
	"latitud": -34.7625601, //float
	"longitud": -58.2192142, //float
}
```


### Buscar Sucursal por ID
- Endpoint: /sucursales/:idSucursal
- Método: GET

Ejemplo de request:
```
GET /sucursales/5edb503e14e865fa8ac0cf4b
```

Ejemplo de response (status code 200):

```jsonc
{
	"id": "5edb503e14e865fa8ac0cf4b", //string
	"direccion": "Pres. Juan Domingo Perón 4739", //string
	"latitud": -34.7625601, //float
	"longitud": -58.2192142, //float
}
```


## Licencia

[MIT](https://choosealicense.com/licenses/mit/)
