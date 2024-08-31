package resto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	AddressLine1 string  `json:"address_line_1" bson:"address_line_1"`
	City         string  `json:"city" bson:"city"`
	State        string  `json:"state" bson:"state"`
	PinCode      string  `json:"pin_code" bson:"pin_code"`
	Latitude     float64 `json:"latitude" bson:"latitude"`
	Longitude    float64 `json:"longitude" bson:"longitude"`
}

type Contact struct {
	MobileNumber []string `json:"mobile_number" bson:"mobile_number"`
	EmailId      string   `json:"email_id" bson:"email_id"`
}

type RegistrationDetails struct {
	GstnNumber      string    `json:"gstn_number" bson:"gstn_number"`
	CstnNumber      string    `json:"cstn_number" bson:"ctsn_number"`
	EstablishedDate time.Time `json:"established_date" bson:"established_date"`
}

type PaymentDetails struct {
	UpiCode    string `json:"upi_code" bson:"upi_code"`
	IsVerified bool   `json:"is_verified" bson:"is_verified"`
	IsActive   bool   `json:"is_active" bson:"is_active"`
	QRCodePath string `json:"qr_code_path" bson:"qr_code_path"`
}

type Restaurant struct {
	ID                  primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name                string               `json:"name" bson:"name"`
	CuisineTypes        []string             `json:"cuisine_types" bson:"cuisine_types"`
	Address             *Address             `json:"address" bson:"address"`
	Contact             *Contact             `json:"contact" bson:"contact"`
	RegistrationDetails *RegistrationDetails `json:"registration_details" bson:"registration_details"`
	PaymentDetails      *PaymentDetails      `json:"payment_details" bson:"payment_details"`
	OpenTime            string               `json:"open_time" bson:"open_time"`
	CloseTime           string               `json:"close_time" bson:"close_time"`
	IsVerified          bool                 `json:"is_verified" bson:"is_verified"`
	CreatedAt           time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at" bson:"updated_at"`
}
