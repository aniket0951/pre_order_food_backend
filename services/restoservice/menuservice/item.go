package menuservice

import (
	"errors"
	"log"
	"pre_order_food_resto_module/dtos/restodto"
	"pre_order_food_resto_module/model/resto"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *menuService) CreateItem(req restodto.CreateItemDTO) (restodto.ItemObject, error) {

	dietaryType := strings.ToLower(req.DietaryType)
	if dietaryType != "veg" && dietaryType != "non_veg" {
		return restodto.ItemObject{}, errors.New("invalid dietaryType, should be VEG, NON_VEG")
	}

	/* ---------------------- check category exitst or not ---------------------- */
	category, err := s.GetCategoryByID(req.CategoryID)
	if err != nil {
		return restodto.ItemObject{}, err
	}

	/* ---------------------- check menu card exists or not --------------------- */
	menuCard, err := s.GetMenuCardById(req.MenuCardId)
	if err != nil {
		return restodto.ItemObject{}, err
	}

	/* -------------------- check item already exists or not -------------------- */
	_, err = s.GetItemByMenuCardID(req.MenuCardId, req.Name)
	log.Println("check item in menu card : ", err, errors.Is(err, gorm.ErrRecordNotFound))
	if err == nil {
		return restodto.ItemObject{}, errors.New("item already exists")
	}

	args := resto.Item{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  category.ID,
		MenuCardID:  menuCard.ID,
		Category:    category,
		MenuCard:    resto.MenuCard(menuCard),
		DietaryType: req.DietaryType,
	}

	items, err := s.menuRepo.CreateItem(args)
	if err != nil {
		return restodto.ItemObject{}, err
	}

	item := restodto.ItemObject{
		ID:          items.ID,
		Name:        items.Name,
		Description: items.Description,
		CategoryID:  items.CategoryID,
		MenuCardID:  items.MenuCardID,
		DietaryType: items.DietaryType,
		IsAvailable: items.IsAvailable,
		// Category:    items.Category,
		// MenuCard:    items.MenuCard,
		CreatedAt: items.CreatedAt,
		UpdatedAt: items.UpdatedAt,
	}

	return item, nil
}
func (s *menuService) GetItemByMenuCardID(menuCardID string, itemName string) (*restodto.ItemObject, error) {
	objID, err := uuid.Parse(menuCardID)
	if err != nil {
		return &restodto.ItemObject{}, errors.New("invalid menu card")
	}

	item, err := s.menuRepo.GetItemByMenuCardID(objID, itemName)
	if err != nil {
		return &restodto.ItemObject{}, err
	}

	itemPrice := make([]restodto.ItemPriceObject, len(item.Prices))

	for i, price := range item.Prices {
		itemPrice[i] = restodto.ItemPriceObject{
			ID:    price.ID,
			Size:  price.Size,
			Price: price.Price,
		}
	}

	itemObj := restodto.ItemObject{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		CategoryID:  item.CategoryID,
		MenuCardID:  item.MenuCardID,
		DietaryType: item.DietaryType,
		IsAvailable: item.IsAvailable,
		Prices:      itemPrice,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
	log.Println("Returning Item Object  : ", itemObj)
	return &itemObj, nil
}
func (s *menuService) UpdateItem(ctx *gin.Context, req restodto.UpdateItemDTO) error {
	objID, err := uuid.Parse(req.ID)
	if err != nil {
		return errors.New("invalid item id")
	}

	/* ------------------------ check item exists or not ------------------------ */
	item, err := s.menuRepo.GetItemByID(objID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("item not found")
		}
		return err
	}
	if item.ID != objID {
		return errors.New("item not found")
	}

	/* -------------------- check DietaryType is VEG,NON_VEG -------------------- */
	dietaryType := strings.ToLower(req.DietaryType)
	if dietaryType != "veg" && dietaryType != "non_veg" {
		return errors.New("invalid DietaryType, it should be VEG, NON_VEG")
	}

	item.Name = req.Name
	item.Description = req.Description
	item.DietaryType = req.DietaryType

	_, err = s.menuRepo.UpdateItem(item)
	return err
}
func (s *menuService) ListItemsByMenuCardID(menuCardID string) ([]restodto.ItemObject, error) {
	objID, err := uuid.Parse(menuCardID)

	if err != nil {
		return nil, err
	}

	result, err := s.menuRepo.ListItesmsByMenuCard(objID)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("items not found for this menu card")
	}

	itemObj := make([]restodto.ItemObject, len(result))

	for i, item := range result {
		itemPrice := make([]restodto.ItemPriceObject, len(item.Prices))

		for i, price := range item.Prices {
			itemPrice[i] = restodto.ItemPriceObject{
				ID:    price.ID,
				Size:  price.Size,
				Price: price.Price,
			}
		}

		itemObj[i] = restodto.ItemObject{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			CategoryID:  item.CategoryID,
			MenuCardID:  item.MenuCardID,
			DietaryType: item.DietaryType,
			IsAvailable: item.IsAvailable,
			Category:    item.Category,
			// MenuCard:    item.MenuCard,s
			Prices:    itemPrice,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
	}
	return itemObj, nil
}

/* -------------------------------------------------------------------------- */
/*                                  ItemPrice                                 */
/* -------------------------------------------------------------------------- */
func (s *menuService) AddItemPrice(req restodto.CreateItemPriceDTO) error {
	objId, err := uuid.Parse(req.ItemID)

	if err != nil {
		return errors.New("invalid item id")
	}

	size := strings.ToLower(req.Size)
	if size != "small" && size != "medium" && size != "large" {
		return errors.New("invalid item size")
	}

	/* ------------------------ check item exists or not ------------------------ */
	item, err := s.menuRepo.GetItemByID(objId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("item not found")
		}
		return err
	}

	/* -------------------------- create a price first -------------------------- */
	args := resto.ItemPrice{
		ItemID: objId,
		Size:   req.Size,
		Price:  req.Price,
	}
	itemPrice, err := s.menuRepo.AddItemPrice(args)
	if err != nil {
		return err
	}

	/* ------------------------ add item price into items ----------------------- */
	item.Prices = append(item.Prices, itemPrice)
	_, err = s.menuRepo.UpdateItem(item)
	if err != nil {
		// remove created price
		if err := s.menuRepo.RemoveItemPrice(objId, itemPrice.ID); err != nil {
			return err
		}
		return err
	}
	return nil
}
func (s *menuService) RemoveItemPrice(req restodto.RemoveItemPriceDTO) error {
	itemObjID, err := uuid.Parse(req.ItemID)
	if err != nil {
		return errors.New("invalid id")
	}
	// ID == itemPriceID
	itemPriceID, err := uuid.Parse(req.ItemPriceID)
	if err != nil {
		return errors.New("invalid id")
	}

	return s.menuRepo.RemoveItemPrice(itemObjID, itemPriceID)
}
