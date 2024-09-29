package orderdto

/* -------------------------------------------------------------------------- */
/*                              CreatePreOrderReq                             */
/* -------------------------------------------------------------------------- */
type CreatePreOrderReq struct {
	RestaurantID     string   `json:"restaurant_id" validator:"required"`
	UserID           string   `json:"user_id" validator:"user_id"`
	ItemID           []string `json:"item_id"`
	Note             string   `json:"note"`
	TotalAmount      int64    `json:"total_amount" validator:"required"`
	PaidAmount       int64    `json:"paid_amount" validator:"required"`
	SenderUPI        string   `json:"sender_upi" validator:"required"`
	ReceiverUPI      string   `json:"receiver_upi" validator:"required"`
	TransactionID    string   `json:"transaction_id" validator:"required"`
	TransactionRefID string   `json:"transaction_ref_id" validator:"required"`
	PaymentStatus    string   `json:"payment_status" validator:"required"`
	AvailableTime    string   `json:"available_time" validator:"required"`
}
