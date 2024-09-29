package orderservice

import (
	"pre_order_food_resto_module/dtos/orderdto"
	"pre_order_food_resto_module/model/ordermodel"
	"pre_order_food_resto_module/repositories/order"
	"pre_order_food_resto_module/repositories/restorepo"
	"pre_order_food_resto_module/services/restoservice/menuservice"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

/* -------------------------------------------------------------------------- */
/*                                  INTERFACE                                 */
/* -------------------------------------------------------------------------- */
type Interface interface {
	CreatePreOrder(req orderdto.CreatePreOrderReq) error
	GetPreOrderByID(preOrderID string) (ordermodel.PreOrders, error)
	GetPreOrderByUserID(userID string) (ordermodel.PreOrders, error)
	ListPreOrders() ([]ordermodel.PreOrders, error)
}

/* -------------------------------------------------------------------------- */
/*                                  RECEIVER                                  */
/* -------------------------------------------------------------------------- */
type orderService struct {
	orderRepo order.Interface
	menuSvc   menuservice.MenuService
}

/* -------------------------------------------------------------------------- */
/*                                   HANDLER                                  */
/* -------------------------------------------------------------------------- */
func OrderService(repo order.Interface) *orderService {
	itemRepo := restorepo.MenuHandler()
	menuSvc := menuservice.Handler(itemRepo)
	return &orderService{orderRepo: repo, menuSvc: menuSvc}
}

func (svc *orderService) CreatePreOrder(req orderdto.CreatePreOrderReq) error {
	resObj, err := uuid.Parse(req.RestaurantID)
	if err != nil {
		return errors.New("invalid restaurant id")
	}

	userObj, err := uuid.Parse(req.RestaurantID)
	if err != nil {
		return errors.New("invalid user id")
	}

	if len(req.ItemID) == 0 {
		return errors.New("invalid items")
	}

	// validate all items
	for _, item := range req.ItemID {
		_, err = svc.menuSvc.GetItemByID(item)
		if err != nil {
			return err
		}
	}

	if req.TotalAmount == 0 || req.PaidAmount == 0 {
		return errors.New("invalid amounts")
	}

	// parse the avl time
	layout := "200601021504"

	avlT, err := time.Parse(layout, req.AvailableTime)
	if err != nil {
		return errors.New("available time invalid")
	}

	args := ordermodel.PreOrders{
		RestaurantID:     resObj,
		UserID:           userObj,
		ItemsID:          req.ItemID,
		Note:             req.Note,
		TotalAmount:      req.TotalAmount,
		PaidAmount:       req.PaidAmount,
		SenderUPI:        req.SenderUPI,
		ReceiverUPI:      req.ReceiverUPI,
		TransactionID:    req.TransactionID,
		TransactionRefID: req.TransactionRefID,
		PaymentStatus:    req.PaymentStatus,
		AvailableTime:    avlT,
	}

	err = svc.orderRepo.CreatePreOrder(args)
	if err != nil {
		return errors.Wrap(err, "[CreatePreOrder][CreatePreOrder]")
	}
	return nil
}

func (svc *orderService) GetPreOrderByID(preOrderID string) (ordermodel.PreOrders, error) {
	objID, err := uuid.Parse(preOrderID)

	if err != nil {
		return ordermodel.PreOrders{}, errors.New("invalid pre order id")
	}

	preOrder, err := svc.orderRepo.GetPreOrderByID(objID)

	if err != nil {
		return preOrder, errors.Wrap(err, "[GetPreOrderByID][GetPreOrderByID]")
	}

	return preOrder, nil
}

func (svc *orderService) GetPreOrderByUserID(userID string) (ordermodel.PreOrders, error) {
	objID, err := uuid.Parse(userID)

	if err != nil {
		return ordermodel.PreOrders{}, errors.New("invalid pre order id")
	}

	preOrder, err := svc.orderRepo.GetPreOrderByUserID(objID)

	if err != nil {
		return preOrder, errors.Wrap(err, "[GetPreOrderByUserID][GetPreOrderByUserID]")
	}

	return preOrder, nil
}

func (svc *orderService) ListPreOrders() ([]ordermodel.PreOrders, error) {
	preOrders, err := svc.orderRepo.ListPreOrders()

	if err != nil {
		return preOrders, errors.Wrap(err, "[ListPreOrders][ListPreOrders]")
	}
	return preOrders, nil
}
