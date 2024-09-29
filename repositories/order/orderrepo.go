package order

import (
	"pre_order_food_resto_module/connections"
	"pre_order_food_resto_module/model/ordermodel"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  INTERFACE                                 */
/* -------------------------------------------------------------------------- */
type Interface interface {
	CreatePreOrder(preOrder ordermodel.PreOrders) error
	GetPreOrderByID(preOrderID uuid.UUID) (ordermodel.PreOrders, error)
	GetPreOrderByUserID(userID uuid.UUID) (ordermodel.PreOrders, error)
	ListPreOrders() ([]ordermodel.PreOrders, error)
}

/* -------------------------------------------------------------------------- */
/*                                  RECEIVER                                  */
/* -------------------------------------------------------------------------- */
type orderRepo struct {
	DB *gorm.DB
}

/* -------------------------------------------------------------------------- */
/*                                   HANDLER                                  */
/* -------------------------------------------------------------------------- */
func OrderHandler() *orderRepo {
	return &orderRepo{DB: connections.DB()}
}

func (repo *orderRepo) CreatePreOrder(preOrder ordermodel.PreOrders) error {
	db := repo.DB.Session(&gorm.Session{})

	result := db.Create(&preOrder)
	if result.Error != nil {
		return errors.Wrap(result.Error, "[CreatePreOrder][Create]")
	}
	return nil
}

func (repo *orderRepo) GetPreOrderByID(preOrderID uuid.UUID) (ordermodel.PreOrders, error) {
	var preOrder ordermodel.PreOrders

	result := repo.DB.Where("id = ?", preOrderID).Take(&preOrder)

	if result.Error != nil {
		return preOrder, errors.Wrap(result.Error, "[GetPreOrder][Take]")
	}

	return preOrder, nil
}

func (repo *orderRepo) GetPreOrderByUserID(userID uuid.UUID) (ordermodel.PreOrders, error) {
	var preOrder ordermodel.PreOrders

	result := repo.DB.Where("user_id = ?", userID).Take(&preOrder)
	if result.Error != nil {
		return preOrder, errors.Wrap(result.Error, "[GetPreOrderByUserID][Take]")
	}

	return preOrder, nil
}

func (repo *orderRepo) ListPreOrders() ([]ordermodel.PreOrders, error) {
	var preOrders []ordermodel.PreOrders
	result := repo.DB.Find(&preOrders)

	if result.Error != nil {
		return preOrders, errors.Wrap(result.Error, "[ListPreOrders][Find]")
	}

	return preOrders, nil
}
