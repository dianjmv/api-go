# Proyecto de Go con MongoDB

Este proyecto es un ejemplo de una API Backend-for-Frontend (BFF) desarrollada con Go y MongoDB. La API se encarga de integrar datos de dos colecciones (`client` y `client-portfolio`) y proporciona endpoints para acceder a estos datos.

## Estructura del Proyecto

El proyecto sigue una estructura de Clean Architecture para mantener el código organizado y modular. Aquí está la estructura del proyecto:
```
go-api/
│
├── cmd/
│ └── main.go
├── internal/
│ ├── app/
│ │ └── app.go
│ ├── client/
│ │ ├── delivery/
│ │ │ └── http/
│ │ │ └── handler.go
│ │ ├── repository/
│ │ │ └── mongo/
│ │ │ └── client_repository.go
│ │ └── usecase/
│ │ └── client_usecase.go
│ └── clientportfolio/
│ ├── delivery/
│ │ └── http/
│ │ └── handler.go
│ ├── repository/
│ │ └── mongo/
│ │ └── clientportfolio_repository.go
│ └── usecase/
│ └── clientportfolio_usecase.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── init-mongo.js
└── wait-for-it.sh
```


### Descripción de los Directorios y Archivos

- **cmd/**: Contiene el punto de entrada principal de la aplicación.
    - `main.go`: Inicia la aplicación llamando al paquete `app`.

- **internal/**: Contiene la lógica interna de la aplicación organizada por módulos.
    - **app/**: Contiene la configuración de la aplicación.
        - `app.go`: Inicializa y configura la aplicación, incluyendo los repositorios, casos de uso y handlers.
    - **client/**: Contiene la lógica relacionada con la entidad `Client`.
        - **delivery/http/**: Contiene los handlers HTTP para los endpoints relacionados con `Client`.
            - `handler.go`: Define los endpoints y sus handlers.
        - **repository/mongo/**: Contiene la implementación del repositorio de MongoDB para `Client`.
            - `client_repository.go`: Implementa las operaciones de MongoDB para `Client`.
        - **usecase/**: Contiene los casos de uso para `Client`.
            - `client_usecase.go`: Define la lógica de negocio para `Client`.
    - **clientportfolio/**: Contiene la lógica relacionada con la entidad `ClientPortfolio`.
        - **delivery/http/**: Contiene los handlers HTTP para los endpoints relacionados con `ClientPortfolio`.
            - `handler.go`: Define los endpoints y sus handlers.
        - **repository/mongo/**: Contiene la implementación del repositorio de MongoDB para `ClientPortfolio`.
            - `clientportfolio_repository.go`: Implementa las operaciones de MongoDB para `ClientPortfolio`.
        - **usecase/**: Contiene los casos de uso para `ClientPortfolio`.
            - `clientportfolio_usecase.go`: Define la lógica de negocio para `ClientPortfolio`.

- `Dockerfile`: Define la configuración del contenedor Docker para la aplicación Go.
- `docker-compose.yml`: Define los servicios Docker, incluyendo MongoDB y la aplicación Go.
- `go.mod` y `go.sum`: Archivos de gestión de dependencias de Go.
- `init-mongo.js`: Script de inicialización de MongoDB para poblar las colecciones `client` y `client-portfolio`.
- `wait-for-it.sh`: Script para esperar a que MongoDB esté listo antes de iniciar la aplicación Go.

## Configuración y Ejecución

### Prerrequisitos

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Instrucciones para Ejecutar el Proyecto

1. **Clonar el Repositorio**

   ```sh
   git clone https://github.com/tu-usuario/go-api.git
   cd go-api
Construir y Levantar los Contenedores

docker-compose up --build
Verificar la Inicialización de los Datos

Puedes verificar que los datos se han insertado correctamente en MongoDB utilizando la consola de MongoDB o cualquier cliente MongoDB como MongoDB Compass.

docker exec -it mongo bash
mongo
use clientDB
db.client.find().pretty()

use clientPortfolioDB
db["client-portfolio"].find().pretty()
Probar los Endpoints

Puedes usar curl o cualquier herramienta como Postman para probar los endpoints de la API.

curl http://localhost:8080/client?clientId=9999999999
curl http://localhost:8080/client-portfolio?customerCode=0001005922
Licencia

Este proyecto está licenciado bajo la Licencia MIT. Ver el archivo LICENSE para más detalles.


Este `README.md` proporciona una descripción clara de la estructura del proyecto, cómo configurarlo y ejecutarlo, y cómo probar los endpoints. Si necesitas más detalles o ajustes específicos, no dudes en decírmelo.