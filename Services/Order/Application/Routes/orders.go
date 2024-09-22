package order_routes

import (
	order_application_services "delivery/Services/Order/Application/Services"
	order_usergateway_controllers "delivery/Services/Order/UserGateway/Controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var OrderRouter = func(router *gin.Engine, db *gorm.DB) {
	controller := order_usergateway_controllers.NewOrderController()
	order := router.Group("/order/")
	_ = order.POST("/create", nil)
	_ = order.POST("/cancel", nil)
	_ = order.POST("/history/{id}", nil)
	_ = order.POST("/pickup-confirm/{id}", nil)
	_ = router.POST("/chat", nil)
}
