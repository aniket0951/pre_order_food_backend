package ordercontroller

import (
	"net/http"
	"pre_order_food_resto_module/dtos/orderdto"
	"pre_order_food_resto_module/services/restoservice/orderservice"
	"pre_order_food_resto_module/utils"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                                  INTERFACE                                 */
/* -------------------------------------------------------------------------- */
type Interface interface {
	CreatePreOrder(ctx *gin.Context)
	GetPreOrderByID(ctx *gin.Context)
	GetPreOrderByUserID(ctx *gin.Context)
	ListPreOrders(ctx *gin.Context)
}

/* -------------------------------------------------------------------------- */
/*                                   HANDLER                                  */
/* -------------------------------------------------------------------------- */
type orderController struct {
	orderSvc orderservice.Interface
}

/* -------------------------------------------------------------------------- */
/*                                  RECEIVER                                  */
/* -------------------------------------------------------------------------- */
func Handler(ordersvc orderservice.Interface) *orderController {
	return &orderController{
		orderSvc: ordersvc,
	}
}

func (c *orderController) CreatePreOrder(ctx *gin.Context) {
	var req orderdto.CreatePreOrderReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.orderSvc.CreatePreOrder(req)
	if err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := utils.BuildSuccessResponse("PreOrder has been created ", nil)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) GetPreOrderByID(ctx *gin.Context) {
	req := ctx.Query("preorder_id")

	if req == "" {
		res := utils.BuildFailedResponse("pre order id can not be empty")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	preOrder, err := c.orderSvc.GetPreOrderByID(req)
	if err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := utils.BuildSuccessResponse("PreOrder fetch successfully ", preOrder)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) GetPreOrderByUserID(ctx *gin.Context) {
	req := ctx.Query("user_id")

	if req == "" {
		res := utils.BuildFailedResponse("user id can not be empty")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	preOrder, err := c.orderSvc.GetPreOrderByUserID(req)
	if err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := utils.BuildSuccessResponse("PreOrder fetch successfully ", preOrder)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) ListPreOrders(ctx *gin.Context) {

	preOrder, err := c.orderSvc.ListPreOrders()
	if err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := utils.BuildSuccessResponse("PreOrder fetch successfully ", preOrder)
	ctx.JSON(http.StatusOK, response)
}
