package restodto

type Address struct {
	AddressLine1 string  `json:"address_line_1" validate:"required"`
	City         string  `json:"city" validate:"required"`
	State        string  `json:"state" validate:"required"`
	PinCode      string  `json:"pin_code" validate:"required"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type Contact struct {
	MobileNumber []string `json:"mobile_number" validate:"required"`
	EmailId      string   `json:"email_id" validate:"required,email"`
}

type AddRestaurantDTO struct {
	Name         string   `json:"name" validate:"required,min=4"`
	CuisineTypes []string `json:"cuisine_types" validate:"required"`
	Address      Address  `json:"address" validate:"required"`
	Contact      Contact  `json:"contact" validate:"required"`
	OpenTime     string   `json:"open_time" validate:"required"`
	CloseTime    string   `json:"close_time" validate:"required"`
}

type PaginationRequest struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type RegistrationDetailsDTO struct {
	RestaurantId    string `json:"id" validate:"required"`
	GstnNumber      string `json:"gstn_number" validate:"required"`
	CstnNumber      string `json:"cstn_number" validate:"required"`
	EstablishedDate string `json:"established_date" validate:"required"`
}

type PaymentDetails struct {
	RestaurantId string `json:"id" validate:"required"`
	UpiCode      string `json:"upi_code" validate:"required"`
}
