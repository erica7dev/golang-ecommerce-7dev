package routes

import (
	"github.com/erica7dev/ecommerce-go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
}