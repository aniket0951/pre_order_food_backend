package main

import (
	"pre_order_food_resto_module/connections"
	"pre_order_food_resto_module/routers/restoroutes"
	"pre_order_food_resto_module/utils"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig()
	connections.Connect()
	engine := gin.New()

	restoroutes.RestaurantRouter(engine)
	restoroutes.MenuRouter(engine)

	log.Info("App run successfully..")
	if err := engine.Run(":8282"); err != nil {
		log.Info("Failed to run the app")
	}
}
