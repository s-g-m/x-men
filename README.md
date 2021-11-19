# X-Mens
Proyecto para detectar si un humano es mutante basándose en su secuencia de ADN.

## Como ejecutar el proyecto

Crear archivo `.env` que va en la raíz del proyecto, es necesario para ejecutar el proyecto en local o con docker:

```
ENVIRONMENT=Local
APP_NAME=X-Mens
APP_PORT=:3000
ENVIRONMENT=Local
DB_URI=******
DB_NAME=******
API_KEY="key de prueba"
```

### Local

- Instalar [Go 1.16+] (https://golang.org/doc/install)
- Crear archivo `.env` en la raíz del proyecto y reemplazar con los valores correspondientes
- Ejecutar comando go run .

La variable APP_PORT indica el puerto que expone la aplicación.

### Docker

- Instalar [Docker] (https://www.docker.com/)
- Crear archivo `.env` en la raíz del proyecto y reemplazar con los valores correspondientes
- Ejecuta el comando `docker build -f Dev.DockerFile -t x-mens-app .`
- Ejecuta el comando `docker run --rm -it -p 3000:3000 -v "$(pwd):/app" --env-file ./.env x-mens-app`

La aplicación correrá sobre el puesto 3000.

## Uso

Para todos los endpoint se debe enviar la cabecera `x-api-key` con la información de la variable de entorno `API_KEY`. Excepto por el endpoint de `/health`

### Mutant

El enpoint expuesto es POST `/x-mens/mutant`

Ejemplo Request:
```json
{
    "dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}
```

### Stats

El enpoint expuesto es GET `/x-mens/stats`

Ejemplo Response:
```json
{
    "count_mutant_dna":40, 
    "count_human_dna":100, 
    "ratio":0.4
}
```
