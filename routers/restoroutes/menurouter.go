package restoroutes

import (
	"pre_order_food_resto_module/controller/restocontroller/menu"
	"pre_order_food_resto_module/repositories/restorepo"
	"pre_order_food_resto_module/services/restoservice/menuservice"

	"github.com/gin-gonic/gin"
)

var (
	menuRepo       = restorepo.MenuHandler()
	menuSvc        = menuservice.Handler(menuRepo)
	menuController = menu.MenuHandler(menuSvc)
)

func MenuRouter(router *gin.Engine) {
	/* ---------------------------------- Menu ---------------------------------- */
	menu := router.Group("api/menu")
	{
		menu.POST("/generate", menuController.GenerateMenuCard)
		menu.GET("/restaurant", menuController.GetMenuCardByRestaurant)
		menu.GET("/list", menuController.ListMenuCard)
	}

	/* -------------------------------- category -------------------------------- */
	category := router.Group("api")
	{
		category.GET("/category", menuController.ListCategory)
	}

	/* ---------------------------------- Item ---------------------------------- */
	item := router.Group("api/item")
	{
		item.POST("/create", menuController.CreateItem)
		item.POST("/update", menuController.UpdateItem)
		item.GET("/list", menuController.ListItesmsByMenuCard)
	}

	/* ------------------------------- Item Price ------------------------------- */
	itemPrice := router.Group("api/item")
	{
		itemPrice.POST("/create-price", menuController.AddItemPrice)
		itemPrice.DELETE("/remove-price", menuController.RemoveItemPrice)
	}
}
