package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/swaggo/http-swagger"
	"orderService/config"
	"orderService/controller"
	_ "orderService/docs"
	"orderService/repository"
	"orderService/service"
)

// @title Order Service API
// @version 1.0
// @description This is a Simple application for Orders
// @BasePath /
func main() {
	app := iris.New()
	db, error := config.ConnectToDB()
	if error != nil {
		fmt.Printf("Connection Lost")
		return
	}

	repo := &repository.OrderRepository{DB: db}
	service := &service.OrderService{Repo: repo}
	orderHandler := &controller.OrderController{Service: *service}

	app.Get("/swagger/{any:path}", iris.FromStd(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)))
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, World from order service")
	})
	app.Post("/order", orderHandler.CreateOrder)
	app.Get("/orders", orderHandler.GetOrders)
	app.Get("/order/:id", orderHandler.GetOrderByID)
	app.Patch("/order/:id", orderHandler.UpdateOrderStatus)

	app.Listen(":8080")
}
