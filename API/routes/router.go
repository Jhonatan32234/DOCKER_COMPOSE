package routes

import (
	"API/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
)


func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusOK) // Responde inmediatamente a OPTIONS con 200
            return
        }

        c.Next()
    }
}


func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.RedirectTrailingSlash = false
    
    // Aplicar middleware CORS antes de declarar rutas
    r.Use(CORSMiddleware())

    product := r.Group("/products")
    {
        product.GET("/", controllers.GetProducts)
        product.POST("/", controllers.CreateProduct)
        product.GET("/:id", controllers.GetProduct)
        product.PUT("/:id", controllers.UpdateProduct)
        product.DELETE("/:id", controllers.DeleteProduct)
    }
    r.GET("/nombre", controllers.GetNombre)

    return r
}