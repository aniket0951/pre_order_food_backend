package main

import (
	"pre_order_food_resto_module/routers/restoroutes"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	restoroutes.RestaurantRouter(engine)

	log.Info("App run successfully..")
	if err := engine.Run(":8282"); err != nil {
		log.Info("Failed to run the app")
	}
}
