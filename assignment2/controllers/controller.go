package controllers

import (
	"net/http"

	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/models"
	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/repository"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderRepo repository.OrderRepositoryImpl
}

func ValidateItems(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// GET ALL {GET}
func (oc *OrderController) GetAll(c *gin.Context) {
	orders, err := oc.OrderRepo.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mendapatkan data order",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

// GET BY ID {GET}
func (oc *OrderController) GetById(c *gin.Context) {
	id := c.Param("id")

	order, err := oc.OrderRepo.FindById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Data tersebut tidak ditemukan",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"order": order,
	})
}

// Create {POST}
func (oc *OrderController) Add(c *gin.Context) {
	var newOrder models.Order
	c.BindJSON(&newOrder)

	createdOrder, err := oc.OrderRepo.Add(&newOrder)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menambahkan data order",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"order": createdOrder,
	})
}

// UPDATE {PUT}
func (oc *OrderController) Update(c *gin.Context) {
	id := c.Param("id")

	var updateOrder models.Order
	c.BindJSON(&updateOrder)

	existingOrder, err := oc.OrderRepo.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mencari data",
		})
		return
	}

	existingOrder.OrderedAt = updateOrder.OrderedAt
	existingOrder.CustomerName = updateOrder.CustomerName

	for _, item := range updateOrder.Items {
		existingOrder.Items = append(existingOrder.Items, item)
	}

	updatedOrder, err := oc.OrderRepo.Update(existingOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengupdate data",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"code":  200,
		"order": updatedOrder,
	})
}

// DELETE {DELETE}
func (oc *OrderController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := oc.OrderRepo.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menghapus data order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data order berhasil dihapus",
	})
}
