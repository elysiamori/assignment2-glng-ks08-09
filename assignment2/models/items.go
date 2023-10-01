package models

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"item_id"`
	CreatedAt   string `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt   string `gorm:"type:timestamp" json:"updated_at"`
	ItemCode    string `gorm:"type:varchar(255)" json:"item_code" validate:"required,unique"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Quantity    uint   `gorm:"type:int" json:"quantity"`
	OrderID     uint   `gorm:"type:int" json:"order_id"`
}
