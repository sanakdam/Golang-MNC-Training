package server

import (
	"Golang-MNC-Training/Kode-Assignment-2/server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	orders *controllers.OrderController
}

func NewRouter(orders *controllers.OrderController) *Router {
	return &Router{orders: orders}
}

func (r *Router) Start(port string) {
	router := gin.Default()

	router.GET("/orders", r.orders.GetOrders)
	router.POST("/orders", r.orders.CreateOrder)
	router.PUT("/orders/:id", r.orders.UpdateOrder)
	router.DELETE("/orders/:id", r.orders.DeleteOrder)

	router.Run(port)
}
