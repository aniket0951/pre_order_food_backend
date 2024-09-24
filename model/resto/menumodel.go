package resto

import (
	"time"

	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                                  Menu Card                                 */
/* -------------------------------------------------------------------------- */
type MenuCard struct {
	ID           uuid.UUID  `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID uuid.UUID  `gorm:"column:restaurant_id;unique;not null"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
	IsActive     bool       `gorm:"column:is_active;default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

/* -------------------------------------------------------------------------- */
/*                  Category represents the categories table                  */
/* -------------------------------------------------------------------------- */
type Category struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	Type        string // e.g., "Appetizer", "Dessert"
	IsDeleted   bool   `gorm:"column:is_deleted;deafult:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

/* -------------------------------------------------------------------------- */
/*                       Item represents the items table                      */
/* -------------------------------------------------------------------------- */
type Item struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"column:name;not null"`
	Description string
	CategoryID  uuid.UUID `gorm:"column:category_id;not null"`
	MenuCardID  uuid.UUID `gorm:"column:menu_card_id;not null"`
	DietaryType string    // e.g., "Vegetarian", "Non-Vegetarian"
	IsAvailable bool      `gorm:"default:true"`
	ImageURL    string
	Category    Category    `gorm:"foreignKey:CategoryID"`
	MenuCard    MenuCard    `gorm:"foreignKey:MenuCardID"`
	Prices      []ItemPrice `gorm:"foreignKey:ItemID"`
	IsDeleted   bool        `gorm:"default:false;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

/* -------------------------------------------------------------------------- */
/*                                 Item Prices                                */
/* -------------------------------------------------------------------------- */
type ItemPrice struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	ItemID    uuid.UUID `gorm:"not null"`
	Size      string    // e.g., "Small", "Large"
	Price     float64   `gorm:"not null"`
	Item      Item      `gorm:"foreignKey:ItemID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
