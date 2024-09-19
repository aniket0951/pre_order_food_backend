package menuservice

import (
	"errors"
	"pre_order_food_resto_module/connections"
	"pre_order_food_resto_module/dtos/restodto"
	"pre_order_food_resto_module/model/resto"
	repo "pre_order_food_resto_module/repositories/restorepo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  INTERFACE                                 */
/* -------------------------------------------------------------------------- */
type MenuService interface {
	/* ---------------------------------- Menu ---------------------------------- */
	GenerateMenuCard(restaurantID string) error
	GetMenuCardByRestaurant(restaurantID string) (restodto.MenuObject, error)
	GetMenuCardById(menuCardId string) (restodto.MenuObject, error)
	ListMenuCard() ([]restodto.MenuObject, error)

	/* -------------------------------- Category -------------------------------- */
	LisCategory() ([]resto.Category, error)
	GetCategoryByID(id string) (resto.Category, error)

	/* ---------------------------------- Item ---------------------------------- */
	CreateItem(req restodto.CreateItemDTO) (restodto.ItemObject, error)
	UpdateItem(ctx *gin.Context, req restodto.UpdateItemDTO) error
	ListItemsByMenuCardID(menuCardID string) ([]restodto.ItemObject, error)

	/* -------------------------------- ItemPrice ------------------------------- */
	AddItemPrice(req restodto.CreateItemPriceDTO) error
	RemoveItemPrice(req restodto.RemoveItemPriceDTO) error
}

/* -------------------------------------------------------------------------- */
/*                                  RECEIVER                                  */
/* -------------------------------------------------------------------------- */
type menuService struct {
	menuRepo repo.MenuRepository
}

/* -------------------------------------------------------------------------- */
/*                                   HANDLER                                  */
/* -------------------------------------------------------------------------- */
func Handler(repo repo.MenuRepository) MenuService {
	return &menuService{
		menuRepo: repo,
	}
}

/* ---------------------------------- Menu ---------------------------------- */
func (s *menuService) GenerateMenuCard(restaurantID string) error {
	objID, err := uuid.Parse(restaurantID)

	if err != nil {
		return errors.New("invalid restaurant id")
	}

	// check restaurant exists or not
	var restaurant resto.Restaurant
	db := connections.DB()
	err = db.Model(resto.Restaurant{}).Where("id =?", objID).Take(&restaurant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("restaurant not found")
		}
		return nil
	}

	args := resto.MenuCard{
		RestaurantID: objID,
		Restaurant:   restaurant,
	}

	return s.menuRepo.GenerateMenuCard(args)
}
func (s *menuService) GetMenuCardByRestaurant(restaurantID string) (restodto.MenuObject, error) {
	objID, err := uuid.Parse(restaurantID)

	if err != nil {
		return restodto.MenuObject{}, errors.New("invalid restaurant id")
	}

	result, err := s.menuRepo.GetMenuCardByRestaurant(objID)
	if err != nil {
		return restodto.MenuObject{}, err
	}

	menuCard := restodto.MenuObject{
		ID:           result.ID,
		RestaurantID: result.Restaurant.ID,
		Restaurant:   result.Restaurant,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}
	return menuCard, nil
}
func (s *menuService) GetMenuCardById(menuCardId string) (restodto.MenuObject, error) {
	objID, err := uuid.Parse(menuCardId)

	if err != nil {
		return restodto.MenuObject{}, errors.New("invalid restaurant id")
	}

	result, err := s.menuRepo.GetMenuCardById(objID)
	if err != nil {
		return restodto.MenuObject{}, err
	}

	menuCard := restodto.MenuObject{
		ID:           result.ID,
		RestaurantID: result.Restaurant.ID,
		Restaurant:   result.Restaurant,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}
	return menuCard, nil
}
func (s *menuService) ListMenuCard() ([]restodto.MenuObject, error) {
	result, err := s.menuRepo.ListMenuCard()
	if err != nil {
		return nil, err
	}

	menuCard := make([]restodto.MenuObject, len(result))

	for i, menu := range result {
		menuCard[i] = restodto.MenuObject{
			ID:           menu.ID,
			RestaurantID: menu.RestaurantID,
			Restaurant:   menu.Restaurant,
			CreatedAt:    menu.CreatedAt,
			UpdatedAt:    menu.UpdatedAt,
		}
	}

	return menuCard, nil
}
