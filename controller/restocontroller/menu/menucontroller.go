package menu

import (
	"net/http"
	"pre_order_food_resto_module/services/restoservice/menuservice"
	"pre_order_food_resto_module/utils"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                                  INTERFACE                                 */
/* -------------------------------------------------------------------------- */
type MenuController interface {
	GenerateMenuCard(ctx *gin.Context)
	GetMenuCardByRestaurant(ctx *gin.Context)
	ListMenuCard(ctx *gin.Context)

	CreateItem(ctx *gin.Context)
	UpdateItem(ctx *gin.Context)
	ListItesmsByMenuCard(ctx *gin.Context)

	AddItemPrice(ctx *gin.Context)
	RemoveItemPrice(ctx *gin.Context)

	ListCategory(ctx *gin.Context)
}

/* -------------------------------------------------------------------------- */
/*                                  RECEIVER                                  */
/* -------------------------------------------------------------------------- */
type menuController struct {
	menuSvc menuservice.MenuService
}

/* -------------------------------------------------------------------------- */
/*                                   HANDLER                                  */
/* -------------------------------------------------------------------------- */
func MenuHandler(svc menuservice.MenuService) MenuController {
	return &menuController{
		menuSvc: svc,
	}
}

/* -------------------------------------------------------------------------- */
/*                                  FUNCTIONS                                 */
/* -------------------------------------------------------------------------- */

/* ---------------------------------- Menu ---------------------------------- */
func (c *menuController) GenerateMenuCard(ctx *gin.Context) {
	restaurantID := ctx.Query("id")

	if restaurantID == "" {
		response := utils.BuildFailedResponse("restaurant id required")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := c.menuSvc.GenerateMenuCard(restaurantID)
	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := utils.BuildSuccessResponse("menu card has been generated", nil)
	ctx.JSON(http.StatusOK, response)
}
func (c *menuController) GetMenuCardByRestaurant(ctx *gin.Context) {
	restaurantID := ctx.Query("id")

	if restaurantID == "" {
		response := utils.BuildFailedResponse("restaurant id required")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	menuCard, err := c.menuSvc.GetMenuCardByRestaurant(restaurantID)
	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := utils.BuildSuccessResponse("menu card has been generated", menuCard)
	ctx.JSON(http.StatusOK, response)
}
func (c *menuController) ListMenuCard(ctx *gin.Context) {

	menuCard, err := c.menuSvc.ListMenuCard()
	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := utils.BuildSuccessResponse("menu card has been generated", menuCard)
	ctx.JSON(http.StatusOK, response)
}

func (c *menuController) ListCategory(ctx *gin.Context) {
	categories, err := c.menuSvc.LisCategory()
	if err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Category fetched successfully", categories)
	ctx.JSON(http.StatusOK, response)
}
