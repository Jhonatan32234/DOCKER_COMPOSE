package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"

    "github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No se pudo cargar el archivo .env, usando variables del entorno")
    }

    dsn := os.Getenv("DB_DSN")
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    DB = database
}
