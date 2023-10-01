package routers

import (
	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoute(orderController *controllers.OrderController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/getorders", orderController.GetAll)
	api.GET("/orders/:id", orderController.GetById)
	api.POST("/addorders", orderController.Add)
	api.PUT("/uporders/:id", orderController.Update)
	api.DELETE("/delorders/:id", orderController.Delete)

	return r
}
