package models

type Order struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	CustomerName string `gorm:"type:varchar(255)" json:"customer_name"`
	OrderedAt    string `gorm:"type:timestamp" json:"ordered_at"`
	CreatedAt    string `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt    string `gorm:"type:timestamp" json:"updated_at"`
	Items        []Item `gorm:"foreignKey:OrderID" json:"items"`
}

// List Order
type Orders []Order

// Constructor
func NewOrder() *Order {
	return &Order{}
}
