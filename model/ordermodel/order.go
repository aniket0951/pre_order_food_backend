package ordermodel

import (
	"time"

	"github.com/google/uuid"
)

type PreOrders struct {
	ID               uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	RestaurantID     uuid.UUID `gorm:"not null"`
	UserID           uuid.UUID `gorm:"not null"`
	ItemsID          []string  `gorm:"not null;type:varchar(45)[]"`
	Note             string    `gorm:"column:note;type:varchar(100)"`
	TotalAmount      int64     `gorm:"column:total_amount;not null;type:integer"`
	PaidAmount       int64     `gorm:"column:paid_amount;not null;type:integer"`
	SenderUPI        string    `gorm:"column:sender_upi;not null;type:varchar(20)"`
	ReceiverUPI      string    `gorm:"column:receiver_upi;not null;type:varchar(20)"`
	TransactionID    string    `gorm:"column:transaction_id;not null;type:varchar(100)"`
	TransactionRefID string    `gorm:"column:transaction_ref_id;not null;type:varchar(100)"`
	PaymentStatus    string    `gorm:"column:payment_status;not null;type:varchar(40)"`
	AvailableTime    time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
