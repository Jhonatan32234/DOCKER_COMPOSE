package main

import (
    "API/config"
    "API/entities"
    "API/routes"
)

func main() {
    config.Connect()

    // Migración automática de la entidad Product
    config.DB.AutoMigrate(&entities.Product{})

    r := routes.SetupRouter()
    r.Run(":8000")
}
