package menuservice

import (
	"errors"
	"pre_order_food_resto_module/model/resto"

	"github.com/google/uuid"
)

func (s *menuService) LisCategory() ([]resto.Category, error) {
	result, err := s.menuRepo.LisCategory()
	if err != nil {
		return result, err
	}

	if len(result) == 0 {
		return result, errors.New("category not found")
	}
	return result, nil
}

func (s *menuService) GetCategoryByID(id string) (resto.Category, error) {
	objID, err := uuid.Parse(id)
	if err != nil {
		return resto.Category{}, errors.New("Invalid category id")
	}
	return s.menuRepo.GetCategoryByID(objID)
}
