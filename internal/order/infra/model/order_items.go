package model

import (
	shared_model "github.com/maruki00/deligo/internal/shared/model"
)

type OrderItem struct {
	shared_model.BaseModel
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Qty       int     `json:"qty"`
	UnitPrice float32 `json:"unit_price"`
}

func (_this *OrderItem) GetOrderId() int {
	return _this.OrderId
}
func (_this *OrderItem) GetProductId() int {
	return _this.ProductId
}
func (_this *OrderItem) GetQty() int {
	return _this.Qty
}
func (_this *OrderItem) GetUnitPrice() float32 {
	return _this.UnitPrice
}
