package resto

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Address struct {
	ID           uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID uuid.UUID `gorm:"column:restaurant_id;unique;type:uuid;"`
	AddressLine1 string    `gorm:"column:address_line_1;type:varchar(20);"`
	City         string    `gorm:"column:city;type:varchar(20)"`
	State        string    `gorm:"column:state;type:varchar(20)"`
	PinCode      string    `gorm:"column:pin_code;type:varchar(20)"`
	Latitude     float64   `gorm:"column:latitude;type:double precision"`
	Longitude    float64   `gorm:"column:longitude;type:double precision"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Contact struct {
	ID           uuid.UUID      `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID uuid.UUID      `gorm:"column:restaurant_id;unique;type:uuid;"`
	MobileNumber pq.StringArray `gorm:"column:mobile_number;type:text[];"`
	EmailId      string         `gorm:"column:email_id;type:varchar(40);"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RegistrationDetails struct {
	ID              uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID    uuid.UUID `gorm:"column:restaurant_id;unique;type:uuid;"`
	GstnNumber      string    `gorm:"column:gstn_number;type:varchar(60);"`
	CstnNumber      string    `gorm:"column:cstn_number;type:varchar(60);"`
	EstablishedDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type PaymentDetails struct {
	ID           uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID uuid.UUID `gorm:"column:restaurant_id;unique;type:uuid;"`
	UpiCode      string    `gorm:"column:upi_code;type:varchar(120);"`
	IsVerified   bool      `gorm:"column:is_verified;default:false;"`
	IsActive     bool      `gorm:"column:is_active:default:false;"`
	QRCodePath   string    `gorm:"column:qr_code_path;type:varchar(60)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Restaurant struct {
	ID           uuid.UUID      `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string         `gorm:"column:name;type:varchar(256);index"`
	CuisineTypes pq.StringArray `gorm:"column:cuisine_types;type:text[]"`
	OpenTime     string         `gorm:"column:open_time;type:varchar(20)"`
	CloseTime    string         `gorm:"column:close_time;type:varchar(20)"`
	IsVerified   bool           `gorm:"column:is_verified;default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
