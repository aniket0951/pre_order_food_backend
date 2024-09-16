package restoservice

import (
	"errors"
	"pre_order_food_resto_module/dtos/restodto"
	"pre_order_food_resto_module/model/resto"
	"pre_order_food_resto_module/repositories/restorepo"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type RestaurantService interface {
	AddRestaurant(req *restodto.AddRestaurantDTO) error
	GetRestaurants(page, limit int64) ([]*resto.Restaurant, error)
	GetRestaurant(req string) (*resto.Restaurant, error)
	AddRestaurantAddress(req *restodto.Address) error
	UpdteRestaurantAddress(req *restodto.Address) error
	AddRestaurantContact(req *restodto.Contact) error
	UpdateRestaurant(req *restodto.AddRestaurantDTO) error

	UpdateRestaurantContact(req *restodto.Contact) error

	AddRegistrationDetails(req *restodto.RegistrationDetailsDTO) error

	AddPaymentDetails(req restodto.PaymentDetails) error
	UpdatePaymentDetails(req restodto.PaymentDetails) (resto.PaymentDetails, error)
}

type restoService struct {
	restoRepo restorepo.RestaurantRepository
}

func NewRestaurantService(repo restorepo.RestaurantRepository) RestaurantService {
	return &restoService{restoRepo: repo}
}

/* -------------------------------------------------------------------------- */
/*                                 Restaurant                                 */
/* -------------------------------------------------------------------------- */
func (r *restoService) AddRestaurant(req *restodto.AddRestaurantDTO) error {

	args := resto.Restaurant{}
	args.Name = req.Name
	args.CuisineTypes = req.CuisineTypes
	args.OpenTime = req.OpenTime
	args.CloseTime = req.CloseTime
	args.IsVerified = false
	args.CreatedAt = time.Now()
	args.UpdatedAt = time.Now()

	err := r.restoRepo.AddRestaurant(&args)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("restaurant already registerd")
		}
		return err
	}
	return nil
}
func (r *restoService) UpdateRestaurant(req *restodto.AddRestaurantDTO) error {
	args := resto.Restaurant{
		Name:         req.Name,
		CuisineTypes: req.CuisineTypes,
		OpenTime:     req.OpenTime,
		CloseTime:    req.CloseTime,
	}

	objId, err := primitive.ObjectIDFromHex(req.RestaurantId)

	if err != nil {
		return errors.New("invalid id found")
	}

	return r.restoRepo.UpdateRestaurant(args, objId)
}
func (r *restoService) GetRestaurants(page, limit int64) ([]*resto.Restaurant, error) {

	offset := (page - 1) * limit
	result, err := r.restoRepo.GetRestaurants(offset, limit)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("data not found")
	}
	return result, nil
}
func (r *restoService) GetRestaurant(req string) (*resto.Restaurant, error) {
	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	var result *resto.Restaurant
	var err error
	t := strings.ReplaceAll(req, " ", "")
	if !re.MatchString(t) {
		objId, err := uuid.Parse(req)
		if err != nil {
			return nil, errors.New("invalid id")
		}
		result, err = r.restoRepo.GetRestaurant(objId, true)
		if result == nil || err != nil {
			return nil, errors.New("restaurant not found")
		}

		return result, nil
	}

	result, err = r.restoRepo.GetRestaurant(req, false)
	if result == nil || err != nil {
		return nil, errors.New("restaurant not found")
	}
	return result, nil
}

/* -------------------------------------------------------------------------- */
/*                             Restaurant Address                             */
/* -------------------------------------------------------------------------- */
func (r *restoService) AddRestaurantAddress(req *restodto.Address) error {
	args := resto.Address{
		AddressLine1: req.AddressLine1,
		State:        req.State,
		City:         req.City,
		PinCode:      req.PinCode,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		UpdatedAt:    time.Now(),
	}

	objId, err := uuid.Parse(req.RestaurantId)

	if err != nil {
		return errors.New("invalid id found")
	}

	/* ------------------------ RESTAURANT EXISTS OR NOT ------------------------ */
	restaurant, err := r.restoRepo.GetRestaurant(objId, true)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("restaurant not found")
		}
		return err
	}

	if restaurant == nil {
		return errors.New("restaurant not found")
	}

	args.RestaurantID = objId
	err = r.restoRepo.AddRestaurantAddress(&args)
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		return errors.New("address alreay exists")
	}
	return err
}
func (r *restoService) UpdteRestaurantAddress(req *restodto.Address) error {
	args := resto.Address{
		AddressLine1: req.AddressLine1,
		State:        req.State,
		City:         req.City,
		PinCode:      req.PinCode,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		UpdatedAt:    time.Now(),
	}

	objId, err := uuid.Parse(req.RestaurantId)
	if err != nil {
		return errors.New("invalid id found")
	}

	args.RestaurantID = objId
	err = r.restoRepo.UpdteRestaurantAddress(&args)
	return err
}

/* -------------------------------------------------------------------------- */
/*                              RestaurantContact                             */
/* -------------------------------------------------------------------------- */
func (r *restoService) AddRestaurantContact(req *restodto.Contact) error {
	// check restaurant exists or not
	restaurant, err := r.GetRestaurant(req.RestaurantId)
	if err != nil {
		return err
	}

	args := resto.Contact{
		RestaurantID: restaurant.ID,
		MobileNumber: req.MobileNumber,
		EmailId:      req.EmailId,
	}

	err = r.restoRepo.AddRestaurantContact(args)

	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		return errors.New("contact already found")
	}

	return nil
}

func (r *restoService) UpdateRestaurantContact(req *restodto.Contact) error {
	// parse uuid
	objId, err := uuid.Parse(req.RestaurantId)
	if err != nil {
		return errors.New("invalid restaurant id")
	}

	args := resto.Contact{
		RestaurantID: objId,
		MobileNumber: req.MobileNumber,
		EmailId:      req.EmailId,
	}

	return r.restoRepo.UpdateRestaurantContact(args)
}

/* -------------------------------------------------------------------------- */
/*                       Restaurant RegistrationDetails                       */
/* -------------------------------------------------------------------------- */

func (r *restoService) AddRegistrationDetails(req *restodto.RegistrationDetailsDTO) error {
	args := new(resto.RegistrationDetails)
	args.GstnNumber = req.GstnNumber
	args.CstnNumber = req.CstnNumber

	layout := "20060102"
	established_date, err := time.Parse(layout, req.EstablishedDate)

	if err != nil {
		return err
	}

	/* ------------------------------- VALID UUID ------------------------------- */
	objId, err := uuid.Parse(req.RestaurantId)
	if err != nil {
		return errors.New("invalid restaurnt id")
	}

	/* --------------------- CHECK RESTAURANT EXISTS OR NOT --------------------- */
	restaurant, err := r.restoRepo.GetRestaurant(objId, true)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("restaurant not found")
		}
		return err
	}

	if restaurant == nil {
		return errors.New("restaurant not found")
	}

	args.EstablishedDate = established_date
	args.RestaurantID = objId

	err = r.restoRepo.AddRegistrationDetails(*args)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("registration details already exists")
		}
	}

	return err
}

/* -------------------------------------------------------------------------- */
/*                             Restaurant Payment                             */
/* -------------------------------------------------------------------------- */
func (r *restoService) AddPaymentDetails(req restodto.PaymentDetails) error {
	var args resto.PaymentDetails
	args.UpiCode = req.UpiCode

	objId, err := uuid.Parse(req.RestaurantId)

	if err != nil {
		return errors.New("invalid id found")
	}

	/* ------------------------ RESTAURANT EXISTS OR NOT ------------------------ */
	restaurnat, err := r.restoRepo.GetRestaurant(objId, true)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("restaurant not found")
		}
		return err
	}
	if restaurnat == nil {
		return errors.New("restaurant not found")
	}

	args.RestaurantID = objId
	err = r.restoRepo.AddPaymentDetails(args)
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		return errors.New("payment details already exists")
	}

	return err
}
func (r *restoService) UpdatePaymentDetails(req restodto.PaymentDetails) (resto.PaymentDetails, error) {

	objId, err := uuid.Parse(req.RestaurantId)

	if err != nil {
		return resto.PaymentDetails{}, errors.New("invalid id found")
	}

	/* ------------------------ RESTAURANT EXISTS OR NOT ------------------------ */
	restaurnat, err := r.restoRepo.GetRestaurant(objId, true)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resto.PaymentDetails{}, errors.New("restaurant not found")
		}
		return resto.PaymentDetails{}, err
	}
	if restaurnat == nil {
		return resto.PaymentDetails{}, errors.New("restaurant not found")
	}
	var args resto.PaymentDetails
	args.UpiCode = req.UpiCode
	args.RestaurantID = objId

	payment, err := r.restoRepo.UpdatePaymentDetails(args)
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		return payment, nil
	}

	return payment, err
}
