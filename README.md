Objetivo

Diseñar e implementar una arquitectura de microservicios con Docker Compose que integre:

 - Un frontend en React Vite

 - Una API REST en Go

 - Una base de datos MySQL con persistencia

Servicios incluidos
Servicio	Descripción
mysql-jhonatan	Base de datos MySQL con credenciales seguras y volumen persistente
api-jhonatan	API REST en Go con arquitectura MVC, conectada a MySQL y migración automática
frontend-jhonatan	Aplicación React Vite que consume la API y muestra productos + nombre del desarrollador

Estructura del proyecto
Código

├── docker-compose.yml
├── API/
│   ├── .env
│   └── Dockerfile
├── front/
│   ├── .env
│   ├── Dockerfile
│   └── src...

Comandos para levantar el entorno

# 1. Clonar el repositorio
git clone [URL_DEL_REPOSITORIO]
cd mi_repositorio

# 2. Construir y levantar los contenedores
docker-compose up --build

# 3. Verificar servicios
# Frontend: http://localhost:3000
# API: http://localhost:8000/products
# API: http://localhost:8000/nombre

Pruebas de persistencia

# Detener los contenedores
docker-compose down o simplemente CRTL+C para detener el proceso

# Volver a levantar (los datos deben persistir)
docker-compose up
Comunicación entre servicios
El frontend se comunica con la API usando la variable VITE_API_URL definida en .env

La API se conecta a MySQL usando DB_DSN desde su .env

Todos los servicios están en la red interna internal_net y se comunican por nombre de contenedor

Requisitos cumplidos
Dockerfile personalizado para cada servicio 
Variables de entorno en .env 
Volumen persistente para MySQL 
Redes internas y comunicación por nombre 
Endpoint /nombre que retorna tu nombre completo 
Nombre visible en el frontend 
Pruebas de persistencia realizadas