package restocontroller

import (
	"net/http"
	"pre_order_food_resto_module/dtos/restodto"
	"pre_order_food_resto_module/services/restoservice"
	"pre_order_food_resto_module/utils"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type RestaurantController interface {
	AddRestaurant(ctx *gin.Context)
	GetRestaurants(ctx *gin.Context)
	GetRestaurant(ctx *gin.Context)
	AddRestaurantAddress(ctx *gin.Context)
	UpdateRestaurantAddress(ctx *gin.Context)

	AddRestaurantContact(ctx *gin.Context)
	UpdateRestaurantContact(ctx *gin.Context)
	UpdateRestaurant(ctx *gin.Context)

	AddRegistrationDetails(ctx *gin.Context)
	AddPaymentDetails(ctx *gin.Context)
}

type restoController struct {
	restoService restoservice.RestaurantService
}

func NewRestaurantController(service restoservice.RestaurantService) RestaurantController {
	return &restoController{restoService: service}
}

// AddRestaurant implements RestaurantController.
func (r *restoController) AddRestaurant(ctx *gin.Context) {
	var req restodto.AddRestaurantDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	validate := validator.New()

	// Validate the struct
	err := validate.Struct(req)
	if err != nil {

		response := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = r.restoService.AddRestaurant(&req)

	if err != nil {
		log.Info("Error from service : ", err)
		response := utils.BuildFailedResponse(err.Error())
		log.Info("Error receive from response builder : ", response)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildSuccessResponse("Restaurant has been added successfully", nil)

	ctx.JSON(http.StatusCreated, response)
}

func (r *restoController) GetRestaurants(ctx *gin.Context) {
	var req restodto.PaginationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := r.restoService.GetRestaurants(req.Page, req.Limit)

	if err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildSuccessResponse("Data fetched success", result)
	ctx.JSON(http.StatusOK, res)
}
func (r *restoController) GetRestaurant(ctx *gin.Context) {
	req := ctx.Param("tag")

	result, err := r.restoService.GetRestaurant(req)

	if err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildSuccessResponse("data fetched success", result)
	ctx.JSON(http.StatusOK, res)
}

func (r *restoController) AddRegistrationDetails(ctx *gin.Context) {
	req := new(restodto.RegistrationDetailsDTO)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	if err := r.restoService.AddRegistrationDetails(req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildSuccessResponse("restaurant registration details added successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (r *restoController) AddPaymentDetails(ctx *gin.Context) {
	var req restodto.PaymentDetails

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	v := validator.New()

	if err := v.Struct(req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	if err := r.restoService.AddPaymentDetails(req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := utils.BuildSuccessResponse("payment details has been addedd", nil)
	ctx.JSON(http.StatusOK, res)
}

func (r *restoController) AddRestaurantAddress(ctx *gin.Context) {
	req := new(restodto.Address)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := r.restoService.AddRestaurantAddress(req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildSuccessResponse("restaurant address update successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (r *restoController) UpdateRestaurantAddress(ctx *gin.Context) {
	req := new(restodto.Address)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := r.restoService.UpdteRestaurantAddress(req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildSuccessResponse("restaurant address update successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

/* -------------------------------------------------------------------------- */
/*                             Restaurant Contact                             */
/* -------------------------------------------------------------------------- */
func (r *restoController) AddRestaurantContact(ctx *gin.Context) {
	req := new(restodto.Contact)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := r.restoService.AddRestaurantContact(req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildSuccessResponse("restaurant contact added successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
func (r *restoController) UpdateRestaurantContact(ctx *gin.Context) {
	var req restodto.Contact

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := r.restoService.UpdateRestaurantContact(&req)
	if err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildSuccessResponse("restaurant payment update successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (r *restoController) UpdateRestaurant(ctx *gin.Context) {
	req := new(restodto.AddRestaurantDTO)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := r.restoService.UpdateRestaurant(req); err != nil {
		res := utils.BuildFailedResponse(err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildSuccessResponse("restaurant update successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
