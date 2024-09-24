package restorepo

import (
	"errors"
	"pre_order_food_resto_module/connections"

	md "pre_order_food_resto_module/model/resto"

	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RestaurantRepository interface {
	AddRestaurant(args *md.Restaurant) error
	GetRestaurants(page, limit int64) ([]*md.Restaurant, error)
	GetRestaurant(args interface{}, isUUID bool) (*md.Restaurant, error)
	AddRestaurantAddress(args *md.Address) error
	UpdteRestaurantAddress(args *md.Address) error

	AddRestaurantContact(args md.Contact) error
	UpdateRestaurantContact(args md.Contact) error
	UpdateRestaurant(args md.Restaurant, restoId primitive.ObjectID) error

	AddRegistrationDetails(args md.RegistrationDetails) error

	AddPaymentDetails(args md.PaymentDetails) error
	UpdatePaymentDetails(args md.PaymentDetails) (md.PaymentDetails, error)
}

type restoRepo struct {
	db *gorm.DB
}

func NewRestaurantRepository() RestaurantRepository {
	return &restoRepo{
		db: connections.DB(),
	}
}

func (rr *restoRepo) AddRestaurant(args *md.Restaurant) error {
	return rr.db.Create(&args).Error
}

func (rr *restoRepo) GetRestaurants(page, limit int64) ([]*md.Restaurant, error) {
	var restaurants []*md.Restaurant
	err := rr.db.Offset(int(page)).Limit(int(limit)).Find(&restaurants)
	return restaurants, err.Error
}

func (rr *restoRepo) GetRestaurant(args interface{}, isUUID bool) (*md.Restaurant, error) {
	result := new(md.Restaurant)
	var err error
	if isUUID {
		err = rr.db.First(&result, "id=?", args).Error
	} else {
		err = rr.db.First(&result, "name=?", args).Error
	}

	return result, err
}

/* --------------------- RESTAURANT REGISTRATION DETAILS -------------------- */
func (rr *restoRepo) AddRegistrationDetails(args md.RegistrationDetails) error {
	return rr.db.Create(&args).Error
}

/* -------------------------------------------------------------------------- */
/*                             Restaurant Payment                             */
/* -------------------------------------------------------------------------- */
func (rr *restoRepo) AddPaymentDetails(args md.PaymentDetails) error {
	return rr.db.Create(&args).Error
}
func (rr *restoRepo) UpdatePaymentDetails(args md.PaymentDetails) (md.PaymentDetails, error) {
	db := rr.db.Session(&gorm.Session{})

	result := db.Where("restaurant_id =?", args.RestaurantID).Updates(&args)
	return args, result.Error
}

/* -------------------------------------------------------------------------- */
/*                             Restaurant Address                             */
/* -------------------------------------------------------------------------- */
func (rr *restoRepo) AddRestaurantAddress(args *md.Address) error {
	return rr.db.Create(&args).Error
}

func (rr *restoRepo) UpdteRestaurantAddress(args *md.Address) error {

	err := rr.db.Model(&md.Address{}).Where("restaurant_id = ?", args.RestaurantID).Updates(&args)
	if err.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return err.Error
}

/* -------------------------------------------------------------------------- */
/*                             Restaurant Conatct                             */
/* -------------------------------------------------------------------------- */
func (rr *restoRepo) AddRestaurantContact(args md.Contact) error {
	return rr.db.Create(&args).Error
}

func (rr *restoRepo) UpdateRestaurantContact(args md.Contact) error {
	db := rr.db.Session(&gorm.Session{})
	result := db.Where("restaurant_id = ?", args.RestaurantID).Updates(&args)
	return result.Error
}

func (rr *restoRepo) UpdateRestaurant(args md.Restaurant, restoId primitive.ObjectID) error {
	// filter := bson.M{
	// 	"_id": restoId,
	// }

	// update := bson.M{
	// 	"$set": bson.M{
	// 		"name":          args.Name,
	// 		"cuisine_types": args.CuisineTypes,
	// 		"open_time":     args.OpenTime,
	// 		"close_time":    args.CloseTime,
	// 		"updated_at":    time.Now(),
	// 	},
	// }

	// result, err := rr.restoCollection.UpdateOne(context.Background(), filter, update)

	// if err != nil {
	// 	return err
	// }

	// if result.ModifiedCount == 0 {
	// 	return errors.New("failed to update address")
	// }
	return nil
}
