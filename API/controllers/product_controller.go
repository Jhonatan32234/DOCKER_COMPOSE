package controllers

import (
    "API/entities"
    "API/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    products := services.GetAllProducts()
    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var product entities.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    created := services.CreateProduct(product)
    c.JSON(http.StatusCreated, created)
}

func GetProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    product, err := services.GetProductByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var product entities.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product.ID = uint(id)
    services.UpdateProduct(product)
    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    services.DeleteProduct(uint(id))
    c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}
