package service

import (
	"github.com/google/uuid"
	"orderService/model"
	"orderService/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func (os *OrderService) Create(ocd *model.OrderCreateDTO) (*model.Order, error) {
	return os.Repo.Create(ocd)
}
func (os *OrderService) GetOrderById(id uuid.UUID) (*model.Order, error) {
	return os.Repo.GetOrderByID(id)
}

func (os *OrderService) GetOrders() ([]model.Order, error) {
	return os.Repo.GetOrders()
}

func (os *OrderService) UpdateOrderStatus(id uuid.UUID, orderstatus string, paymentstatus string) (*model.Order, error) {
	return os.Repo.UpdateStatus(id, orderstatus, paymentstatus)
}
