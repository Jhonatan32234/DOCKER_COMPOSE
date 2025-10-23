# Imagen base de Go
FROM golang:1.23

# Directorio de trabajo
WORKDIR /app

# Copiar dependencias
COPY go.mod go.sum ./
RUN go mod tidy
RUN apt-get update && apt-get install -y default-mysql-client

# Copiar todo el código fuente y el script de espera
COPY . .
COPY wait-for-mysql.sh .
RUN chmod +x wait-for-mysql.sh

# Compilar la aplicación
RUN go build -o /app/app main.go

# Exponer el puerto
EXPOSE 8000

# Ejecutar script de espera y luego la app (host debe ser el nombre del servicio mysql)
CMD ["./wait-for-mysql.sh", "mysql-jhonatan", "/app/app"]
