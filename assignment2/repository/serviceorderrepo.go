package repository

import (
	"errors"

	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/models"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{DB: db}
}

func (repo *OrderRepositoryImpl) FindAll() ([]models.Order, error) {
	var orders []models.Order
	result := repo.DB.Preload("Items").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (repo *OrderRepositoryImpl) FindById(id string) (*models.Order, error) {
	var order models.Order
	result := repo.DB.Preload("Items").Where("id = ?", id).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *OrderRepositoryImpl) ItemCode(itemCode string) bool {
	var item models.Item
	result := repo.DB.Where("item_code = ?", itemCode).First(&item)
	if result.Error != nil {
		return false
	}

	return true
}

func (repo *OrderRepositoryImpl) Add(order *models.Order) (*models.Order, error) {

	for _, item := range order.Items {
		if repo.ItemCode(item.ItemCode) {
			return nil, errors.New("")
		}
	}

	result := repo.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (repo *OrderRepositoryImpl) Update(order *models.Order) (*models.Order, error) {

	result := repo.DB.Updates(order)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, item := range order.Items {
		result := repo.DB.Updates(&item)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return order, nil
}

func (repo *OrderRepositoryImpl) Delete(id string) error {
	result := repo.DB.Delete(&models.Order{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
