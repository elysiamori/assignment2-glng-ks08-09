package repository

import "github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/models"

type OrderRepository interface {
	FindAll() ([]models.Orders, error)
	FindById(id int64) (*models.Order, error)
	Add(order *models.Order) (*models.Order, error)
	Update(order *models.Order) (*models.Order, error)
	Delete(id int64) error
}
