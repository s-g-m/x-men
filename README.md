# X-Men
Proyecto para detectar si un humano es mutante basándose en su secuencia de ADN.

## Producción 

Url de produccion https://mercadolibre-2021.uk.r.appspot.com/

## Como ejecutar el proyecto

Crear archivo `.env` que va en la raíz del proyecto, es necesario para ejecutar el proyecto en local o con docker:

```
ENVIRONMENT=Local
APP_NAME=x-men
APP_PORT=:3000
DB_URI=******
DB_NAME=Xmen
REDIS_URI=******
API_KEY=key de prueba
```

Para la base de datos mongo montada en docker puedes utilizar `mongodb://host.docker.internal:27017/` o `mongodb://172.17.0.1:27017/`
pero si la base de datos esta instalada localmente, utiliza `mongodb://localhost:27017/`

Para la base de datos redis montada en docker puedes utilizar `host.docker.internal:6379` o `172.17.0.1:6379`
pero si la base de datos esta instalada localmente, utiliza `localhost:6379`

### Local

- Instalar [Go 1.17+](https://golang.org/doc/install)
- Instalar [MongoDB](https://docs.mongodb.com/manual/installation/) o ejecutar comando para crear contenedor de Mongo [ver](#docker)
- Instalar [Redis](https://redis.io/topics/quickstart) o ejecutar comando para crear contenedor de Redis [ver](#docker)
- Crear archivo `.env` en la raíz del proyecto y reemplazar con los valores correspondientes
- Ejecutar comando `go mod tidy`
- Ejecutar comando `go run .`
- Puedes usar Air para recargar la aplicación Go en vivo [Air] (https://github.com/cosmtrek/air)

La variable APP_PORT indica el puerto que expone la aplicación.

### Docker

- Instalar [Docker](https://www.docker.com/)
- Crear un volumen docker para la base de datos con el comando `docker volume create xmen-volumen`
- Crear archivo `.env` en la raíz del proyecto y reemplazar con los valores correspondientes
- Ejecuta el siguiente comando para crear contenedor de Mongo `docker run --rm -it -p 27017:27017 --name xmen-db -v xmen-volumen:/data/db mongo:4` 
- Ejecuta el siguiente comando para crear contenedor de Redis `docker run --rm -it -p 6379:6379 --name redis-db redis`
- Ejecuta el siguiente comando para crear imagen del proyecto`docker build -f Dev.DockerFile -t x-men .`
- Ejecuta el siguiente comando para correr la imagen`docker run --rm -it -p 3000:3000 -v "$(pwd):/app" --env-file ./.env x-men`

La aplicación correrá sobre el puesto 3000.

## Uso

Para todos los endpoint se debe enviar la cabecera `x-api-key` con la información de la variable de entorno `API_KEY`. Excepto por el endpoint de `/health`

### Mutant

El endpoint expuesto es POST `/x-men/mutant`

Ejemplo Request:
```json
{
    "dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}
```

Ejemplo Response status 200:
```json
{
    "message": "is mutant"
}
```

Ejemplo Response status 403:
```json
{
    "message": "no is mutant"
}
```

Ejemplo Response status otros:
```json
{
    "message": "error description"
}
```


### Stats

El endpoint expuesto es GET `/x-men/stats`

Ejemplo Response status 200:
```json
{
    "count_mutant_dna":40, 
    "count_human_dna":100, 
    "ratio":0.4
}
```

Ejemplo Response status otros:
```json
{
    "message": "error description"
}
```