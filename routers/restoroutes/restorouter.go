package restoroutes

import (
	"pre_order_food_resto_module/controller/restocontroller"
	"pre_order_food_resto_module/repositories/restorepo"
	"pre_order_food_resto_module/services/restoservice"

	"github.com/gin-gonic/gin"
)

var (
	restoRepo       = restorepo.NewRestaurantRepository()
	restoService    = restoservice.NewRestaurantService(restoRepo)
	restoController = restocontroller.NewRestaurantController(restoService)
)

func RestaurantRouter(engine *gin.Engine) {

	resto := engine.Group("/api")
	{
		resto.POST("/restaurant", restoController.AddRestaurant)
		resto.GET("/getrestaurants", restoController.GetRestaurants)
		resto.GET("/getrestaurant/:tag", restoController.GetRestaurant)

		resto.PUT("/restaurant", restoController.UpdateRestaurant)
		resto.PUT("/address", restoController.UpdateRestaurantAddress)
		resto.PUT("/contact", restoController.UpdateRestaurantContact)

	}

	resto_registration := engine.Group("/api/resto")
	{
		resto_registration.POST("/registration", restoController.AddRegistrationDetails)
		resto_registration.POST("/payment", restoController.AddPaymentDetails)
	}
}
