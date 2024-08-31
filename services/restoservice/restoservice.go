package restoservice

import (
	"errors"
	"pre_order_food_resto_module/dtos/restodto"
	"pre_order_food_resto_module/model/resto"
	"pre_order_food_resto_module/repositories/restorepo"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RestaurantService interface {
	AddRestaurant(req *restodto.AddRestaurantDTO) (*resto.Restaurant, error)
	GetRestaurants(page, limit int64) ([]*resto.Restaurant, error)
	GetRestaurant(req string) (*resto.Restaurant, error)

	AddRegistrationDetails(req *restodto.RegistrationDetailsDTO) error
	AddPaymentDetails(req restodto.PaymentDetails) error
}

type restoService struct {
	restoRepo restorepo.RestaurantRepository
}

func NewRestaurantService(repo restorepo.RestaurantRepository) RestaurantService {
	return &restoService{restoRepo: repo}
}

// AddRestaurant implements RestaurantService.
func (r *restoService) AddRestaurant(req *restodto.AddRestaurantDTO) (*resto.Restaurant, error) {
	// address dto
	address := new(resto.Address)
	address.AddressLine1 = req.Address.AddressLine1
	address.City = req.Address.City
	address.State = req.Address.State
	address.PinCode = req.Address.PinCode
	address.Latitude = req.Address.Latitude
	address.Longitude = req.Address.Longitude

	// contact dto
	contact := new(resto.Contact)
	contact.MobileNumber = req.Contact.MobileNumber
	contact.EmailId = req.Contact.EmailId

	args := resto.Restaurant{}
	args.Name = req.Name
	args.Contact = contact
	args.Address = address
	args.CuisineTypes = req.CuisineTypes
	args.OpenTime = req.OpenTime
	args.CloseTime = req.CloseTime
	args.CreatedAt = time.Now()
	args.UpdatedAt = time.Now()

	resto, err := r.restoRepo.AddRestaurant(&args)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("restaurant already registerd")
		}
		return nil, err
	}
	return resto, nil
}

func (r *restoService) GetRestaurants(page, limit int64) ([]*resto.Restaurant, error) {
	result, err := r.restoRepo.GetRestaurants(page, limit)
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
		objId, err := primitive.ObjectIDFromHex(req)
		if err != nil {
			return nil, err
		}
		result, err = r.restoRepo.GetRestaurant(objId)
		return result, err
	}

	result, err = r.restoRepo.GetRestaurant(req)
	return result, err
}

func (r *restoService) AddRegistrationDetails(req *restodto.RegistrationDetailsDTO) error {
	args := new(resto.RegistrationDetails)
	args.GstnNumber = req.GstnNumber
	args.CstnNumber = req.CstnNumber

	layout := "20060102"
	t, err := time.Parse(layout, req.EstablishedDate)

	if err != nil {
		return err
	}
	args.EstablishedDate = t

	objId, err := primitive.ObjectIDFromHex(req.RestaurantId)
	if err != nil {
		return errors.New("invalid restaurnt id")
	}

	return r.restoRepo.AddRegistrationDetails(*args, objId)
}

func (r *restoService) AddPaymentDetails(req restodto.PaymentDetails) error {
	var args resto.PaymentDetails
	args.UpiCode = req.UpiCode
	objId, err := primitive.ObjectIDFromHex(req.RestaurantId)

	if err != nil {
		return errors.New("invalid id found")
	}

	return r.restoRepo.AddPaymentDetails(args, objId)
}
