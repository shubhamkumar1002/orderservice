package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"orderService/model"
	"orderService/service"
)

type OrderController struct {
	Service service.OrderService
}

// @Summary Create Order
// @Description Create an Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body model.OrderCreateDTO true "Order details"
// @Success 201 {object} model.Order
// @Router /order/ [post]
func (oc *OrderController) CreateOrder(ctx iris.Context) {
	var order model.OrderCreateDTO
	if err := ctx.ReadJSON(&order); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	newOrder, err := oc.Service.Create(&order)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to create order"})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(newOrder)
}

// @Summary Get Order by OrderID
// @Description Get a Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "order ID"
// @Success 200 {object} model.Order
// @Router /order/{id}  [get]
func (oc *OrderController) GetOrderByID(ctx iris.Context) {
	idParam := ctx.Params().Get("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid order ID"})
		return
	}

	order, err := oc.Service.GetOrderById(id)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": fmt.Sprintf("Order not found with OrderID: %s", id)})
		return
	}

	ctx.JSON(order)
}

// @Summary Get All Orders
// @Description Get an array of Orders
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {array} model.Order
// @Router /orders [get]
func (oc *OrderController) GetOrders(ctx iris.Context) {
	result, err := oc.Service.GetOrders()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Internal server error"})
	}

	if result == nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"message": "No orders found"})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Orders retrieved successfully", "orders": result})
}

// @Summary Update Order status by orderid
// @Description update order status
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "order ID"
// @Param order body model.OrderUpdateDTO true "Order update details"
// @Success 200 {object} string
// @Router /order/{id} [patch]
func (oc *OrderController) UpdateOrderStatus(ctx iris.Context) {
	var orderupdate model.OrderUpdateDTO
	idParam := ctx.Params().Get("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid order ID"})
		return
	}

	if err := ctx.ReadJSON(&orderupdate); err != nil || (orderupdate.OrderStatus == "" && orderupdate.PaymentStatus == "") {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid status"})
		return
	}

	if _, err := oc.Service.UpdateOrderStatus(id, orderupdate.OrderStatus, orderupdate.PaymentStatus); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to update status"})
		return
	}

	ctx.JSON(iris.Map{"message": "Order status updated"})
}
