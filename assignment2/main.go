package main

import (
	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/controllers"
	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/database"
	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/repository"
	"github.com/elysiamori/assignment2-valdy-ramadhan/assignment2/routers"
)

func main() {

	db, err := database.DBConnect()
	if err != nil {
		panic("Gagal menyambungkan ke database")
	}

	orderRepo := repository.NewOrderRepository(db)
	orderController := controllers.OrderController{OrderRepo: *orderRepo}

	router := routers.SetupRoute(&orderController)

	router.Run(":3000")

}
