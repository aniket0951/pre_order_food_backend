package restorepo

import (
	"pre_order_food_resto_module/connections"

	model "pre_order_food_resto_module/model/resto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  INTERFACE                                 */
/* -------------------------------------------------------------------------- */
type MenuRepository interface {
	GenerateMenuCard(args model.MenuCard) error
	GetMenuCardByRestaurant(restaurantID uuid.UUID) (model.MenuCard, error)
	ListMenuCard() ([]model.MenuCard, error)
	GetMenuCardById(menuCardID uuid.UUID) (model.MenuCard, error)

	LisCategory() ([]model.Category, error)
	GetCategoryByID(id uuid.UUID) (model.Category, error)

	CreateItem(args model.Item) (model.Item, error)
	ListItem(offset int) ([]model.Item, error)
	GetItemByMenuCardID(menuCardID uuid.UUID, itemName string) (model.Item, error)
	GetItemByID(itemID uuid.UUID) (model.Item, error)
	UpdateItem(args model.Item) (model.Item, error)
	ListItesmsByMenuCard(menuCardID uuid.UUID) ([]model.Item, error)

	AddItemPrice(args model.ItemPrice) (model.ItemPrice, error)
	RemoveItemPrice(itemID uuid.UUID) error
}

/* -------------------------------------------------------------------------- */
/*                                  RECEIVER                                  */
/* -------------------------------------------------------------------------- */
type menuGorm struct {
	db *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                   HANDLER                                  */
/* -------------------------------------------------------------------------- */
func MenuHandler() MenuRepository {
	return &menuGorm{
		db: connections.DB(),
	}
}

/* -------------------------------------------------------------------------- */
/*                                  FUNCTIONS                                 */
/* -------------------------------------------------------------------------- */

/* -------------------------------- Menu Card ------------------------------- */
func (g *menuGorm) GenerateMenuCard(args model.MenuCard) error {
	result := g.db.Create(&args)
	return result.Error
}
func (g *menuGorm) GetMenuCardByRestaurant(restaurantID uuid.UUID) (model.MenuCard, error) {
	var menuCard model.MenuCard
	result := g.db.Preload("Restaurant").Where("restaurant_id = ?", restaurantID).Take(&menuCard)
	return menuCard, result.Error
}
func (g *menuGorm) ListMenuCard() ([]model.MenuCard, error) {
	var menuCards []model.MenuCard
	result := g.db.Preload("Restaurant").Find(&menuCards)
	return menuCards, result.Error
}
func (g *menuGorm) GetMenuCardById(menuCardID uuid.UUID) (model.MenuCard, error) {
	var menuCard model.MenuCard
	result := g.db.Preload("Restaurant").Where("id = ?", menuCardID).Take(&menuCard)
	return menuCard, result.Error
}

/* -------------------------------- Category -------------------------------- */
func (g *menuGorm) LisCategory() ([]model.Category, error) {
	db := g.db.Session(&gorm.Session{})
	var categories []model.Category
	err := db.Where("is_deleted = ?", false).Find(&categories)
	return categories, err.Error
}
func (g *menuGorm) GetCategoryByID(id uuid.UUID) (model.Category, error) {
	var category model.Category
	result := g.db.Where("id = ?", id).Take(&category)
	return category, result.Error
}

/* ---------------------------------- Item ---------------------------------- */
func (g *menuGorm) CreateItem(args model.Item) (model.Item, error) {
	db := g.db.Session(&gorm.Session{})
	return args, db.Create(&args).Error
}
func (g *menuGorm) ListItem(offset int) ([]model.Item, error) {
	db := g.db.Session(&gorm.Session{})
	var items []model.Item
	err := db.Limit(20).Offset(offset).Find(&items)
	return items, err.Error
}
func (g *menuGorm) GetItemByMenuCardID(menuCardID uuid.UUID, itemName string) (model.Item, error) {
	var item model.Item
	result := g.db.Where("menu_card_id = ? and name = ?", menuCardID, itemName).First(&item)
	return item, result.Error
}
func (g *menuGorm) UpdateItem(args model.Item) (model.Item, error) {
	result := g.db.Where("id = ?", args.ID).Updates(&args)
	return args, result.Error
}
func (g *menuGorm) GetItemByID(itemID uuid.UUID) (model.Item, error) {
	var item model.Item
	return item, g.db.Where("id = ?", itemID).First(&item).Error
}
func (g *menuGorm) ListItesmsByMenuCard(menuCardID uuid.UUID) ([]model.Item, error) {
	var items []model.Item
	result := g.db.Where("menu_card_id = ?", menuCardID).Preload("Category").Preload("Prices").Find(&items)
	return items, result.Error
}

/* ------------------------------- Item Price ------------------------------- */
func (g *menuGorm) AddItemPrice(args model.ItemPrice) (model.ItemPrice, error) {
	args.ID = uuid.New()
	return args, g.db.Create(&args).Error
}
func (g *menuGorm) RemoveItemPrice(itemID uuid.UUID) error {
	return g.db.Where("item_id = ?", itemID).Delete(&model.ItemPrice{}).Error
}
