package orderrouter

import (
	"pre_order_food_resto_module/controller/ordercontroller"
	"pre_order_food_resto_module/repositories/order"
	"pre_order_food_resto_module/services/restoservice/orderservice"

	"github.com/gin-gonic/gin"
)

var (
	preorderRepo = order.OrderHandler()
	preOrderSvc  = orderservice.OrderService(preorderRepo)
	preOrderCnt  = ordercontroller.Handler(preOrderSvc)
)

func OrderRouter(engine *gin.Engine) {

	preOrder := engine.Group("api/preorder")
	{
		preOrder.POST("/create", preOrderCnt.CreatePreOrder)
		preOrder.GET("/id", preOrderCnt.GetPreOrderByID)
		preOrder.GET("/userid", preOrderCnt.GetPreOrderByUserID)
		preOrder.GET("/list", preOrderCnt.ListPreOrders)
	}
}
