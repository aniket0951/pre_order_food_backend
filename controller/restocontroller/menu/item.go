package menu

import (
	"net/http"
	"pre_order_food_resto_module/dtos/restodto"
	"pre_order_food_resto_module/utils"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                                  Function                                  */
/* -------------------------------------------------------------------------- */
func (c *menuController) CreateItem(ctx *gin.Context) {
	var req restodto.CreateItemDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	item, err := c.menuSvc.CreateItem(req)
	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := utils.BuildSuccessResponse("item has beed added in the menu list", item)
	ctx.JSON(http.StatusOK, response)
}
func (c *menuController) UpdateItem(ctx *gin.Context) {
	var req restodto.UpdateItemDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := c.menuSvc.UpdateItem(ctx, req)
	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := utils.BuildSuccessResponse("item has been updated", nil)
	ctx.JSON(http.StatusOK, response)
}
func (c *menuController) ListItesmsByMenuCard(ctx *gin.Context) {
	menuCardID := ctx.Query("id")

	if menuCardID == "" {
		response := utils.BuildFailedResponse("menu card id required")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	itemObj, err := c.menuSvc.ListItemsByMenuCardID(menuCardID)
	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := utils.BuildSuccessResponse("items fetched successfully", itemObj)
	ctx.JSON(http.StatusOK, response)
}

/* -------------------------------------------------------------------------- */
/*                                 Item Price                                 */
/* -------------------------------------------------------------------------- */
func (c *menuController) AddItemPrice(ctx *gin.Context) {
	var req restodto.CreateItemPriceDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := c.menuSvc.AddItemPrice(req)
	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildSuccessResponse("item price has been added", nil)
	ctx.JSON(http.StatusOK, response)
}
func (c *menuController) RemoveItemPrice(ctx *gin.Context) {
	var req restodto.RemoveItemPriceDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := c.menuSvc.RemoveItemPrice(req)

	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := utils.BuildSuccessResponse("Item price has been removed", nil)
	ctx.JSON(http.StatusOK, response)
}
