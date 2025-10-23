package services

import (
    "API/config"
    "API/entities"
)

func GetAllProducts() []entities.Product {
    var products []entities.Product
    config.DB.Find(&products)
    return products
}

func CreateProduct(product entities.Product) entities.Product {
    config.DB.Create(&product)
    return product
}

func GetProductByID(id uint) (entities.Product, error) {
    var product entities.Product
    result := config.DB.First(&product, id)
    return product, result.Error
}

func UpdateProduct(product entities.Product) {
    config.DB.Save(&product)
}

func DeleteProduct(id uint) {
    config.DB.Delete(&entities.Product{}, id)
}
