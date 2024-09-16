package restodto

import (
	"pre_order_food_resto_module/model/resto"
	"time"

	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                                    Menu                                    */
/* -------------------------------------------------------------------------- */
type MenuObject struct {
	ID           uuid.UUID        `json:"id"`
	RestaurantID uuid.UUID        `json:"restaurant_id,omitempty"`
	Restaurant   resto.Restaurant `json:"restaurant"`
	IsActive     bool             `json:"is_active"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

/* -------------------------------------------------------------------------- */
/*                                CreateItemDTO                               */
/* -------------------------------------------------------------------------- */
type CreateItemDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id" validate:"required"`
	MenuCardId  string `json:"menu_card_id" validate:"required"`
	DietaryType string `json:"dietary_type" validate:"required,oneof=VEG NON-VEG"`
}

/* -------------------------------------------------------------------------- */
/*                                UpdateItemDTO                               */
/* -------------------------------------------------------------------------- */
type UpdateItemDTO struct {
	ID          string `json:"item_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	DietaryType string `json:"dietary_type" validate:"required,oneof=VEG NON-VEG"`
}

/* -------------------------------------------------------------------------- */
/*                             CreateItemPriceDTO                             */
/* -------------------------------------------------------------------------- */
type CreateItemPriceDTO struct {
	ItemID string  `json:"item_id" validate:"required"`
	Size   string  `json:"size" validate:"required,oneof=SMALL MEDIUM LARGE"`
	Price  float64 `json:"price" validate:"required"`
}

type ItemObject struct {
	ID          uuid.UUID         `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	CategoryID  uuid.UUID         `json:"category_id"`
	MenuCardID  uuid.UUID         `json:"menu_card_id"`
	DietaryType string            `json:"dietary_type"`
	IsAvailable bool              `json:"is_available"`
	Category    interface{}       `json:"category"`
	MenuCard    interface{}       `json:"menu_card"`
	Prices      []ItemPriceObject `json:"prices"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ItemPriceObject struct {
	ID    uuid.UUID `json:"id"`
	Size  string    `json:"size"`
	Price float64   `json:"price"`
}
