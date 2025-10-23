package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetNombre(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "nombre": "Jhonatan de Jesus Zuñiga Castro",
    })
}
