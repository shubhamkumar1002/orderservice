package repository

import (
	"orderService/model"
	"orderService/pubsub"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (ord *OrderRepository) Create(orderCreateDTO *model.OrderCreateDTO) (*model.Order, error) {
	orderId := uuid.New()

	newOrder := &model.Order{
		ID:          orderId,
		UserID:      orderCreateDTO.UserID,
		ProductID:   orderCreateDTO.ProductID,
		Quantity:    orderCreateDTO.Quantity,
		TotalAmount: orderCreateDTO.TotalAmount,
		OrderStatus: model.OrderPlaced,
		CreatedAt:   time.Now(),
	}

	err := ord.DB.Create(&newOrder).Error
	if err != nil {

		return nil, err
	}

	paymentEvent := model.PaymentEvent{
		OrderID:       newOrder.ID,
		PaymentStatus: string(model.Pending),
		TotalAmount:   newOrder.TotalAmount,
	}
	pubsub.SubmitMessage(paymentEvent)
	return newOrder, nil
}

func (ord *OrderRepository) GetOrderByID(id uuid.UUID) (*model.Order, error) {
	var order model.Order
	if err := ord.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (ord *OrderRepository) GetOrders() ([]model.Order, error) {
	var orders []model.Order
	err := ord.DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (ord *OrderRepository) UpdateStatus(id uuid.UUID, orderstatus string, paymentstatus string) (*model.Order, error) {
	var updateOrder model.Order
	if orderstatus != "" {
		err := ord.DB.Model(&updateOrder).Where("id = ?", id).Update("order_status", orderstatus).Update("updated_at", time.Now()).Error
		if err != nil {
			return nil, err
		}
	}

	if paymentstatus != "" {
		paymentEvent := model.PaymentEvent{
			OrderID:       id,
			PaymentStatus: paymentstatus,
		}
		pubsub.SubmitMessage(paymentEvent)
	}

	return &updateOrder, nil
}
