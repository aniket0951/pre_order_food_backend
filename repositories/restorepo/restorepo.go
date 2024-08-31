package restorepo

import (
	"context"
	"errors"
	"pre_order_food_resto_module/config"
	"time"

	md "pre_order_food_resto_module/model/resto"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	restoConnection = config.GetCollection("restaurant")
)

type RestaurantRepository interface {
	AddRestaurant(args *md.Restaurant) (*md.Restaurant, error)
	GetRestaurants(page, limit int64) ([]*md.Restaurant, error)
	GetRestaurant(args interface{}) (*md.Restaurant, error)

	AddRegistrationDetails(args md.RegistrationDetails, restoId primitive.ObjectID) error
	AddPaymentDetails(args md.PaymentDetails, restoId primitive.ObjectID) error
}

type restoRepo struct {
	restoCollection *mongo.Collection
}

func NewRestaurantRepository() RestaurantRepository {
	if err := config.CreateUniqueIndex(restoConnection, "name"); err != nil {
		log.Error("Error while creating index : ", err)
	}
	return &restoRepo{
		restoCollection: restoConnection,
	}
}

func (rr *restoRepo) AddRestaurant(args *md.Restaurant) (*md.Restaurant, error) {

	obj, err := rr.restoCollection.InsertOne(context.Background(), &args)
	if err != nil {
		return nil, err
	}
	if _, ok := obj.InsertedID.(primitive.ObjectID); ok {
		args.ID = obj.InsertedID.(primitive.ObjectID)
	}
	return args, nil
}

func (rr *restoRepo) GetRestaurants(page, limit int64) ([]*md.Restaurant, error) {
	opts := options.Find()
	opts.SetSkip(page)
	opts.SetLimit(limit)

	cursor, err := rr.restoCollection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var result []*md.Restaurant
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (rr *restoRepo) GetRestaurant(args interface{}) (*md.Restaurant, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"_id": args},
			{"name": args},
		},
	}
	result := new(md.Restaurant)
	err := rr.restoCollection.FindOne(context.Background(), filter).Decode(&result)
	return result, err
}

// RESTAURANT REGISTRATION DETAILS
func (rr *restoRepo) AddRegistrationDetails(args md.RegistrationDetails, restoId primitive.ObjectID) error {
	filter := bson.M{
		"_id": restoId,
	}

	update := bson.M{
		"$set": bson.M{
			"registration_details": args,
			"updated_at":           time.Now(),
		},
	}

	result, err := rr.restoCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		return errors.New("failed to add restaurant registration details")
	}
	return nil
}

func (rr *restoRepo) AddPaymentDetails(args md.PaymentDetails, restoId primitive.ObjectID) error {
	filter := bson.M{
		"_id": restoId,
	}

	update := bson.M{
		"$set": bson.M{
			"payment_details": args,
			"updated_at":      time.Now(),
		},
	}
	log.Info(filter, update)
	result, err := rr.restoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		return errors.New("failed to add payment details")
	}
	return nil
}
