package model

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string
type PaymentStatus string

const (
	Pending       PaymentStatus = "PENDING"
	Paid          PaymentStatus = "PAID"
	Cacelled      PaymentStatus = "CANCELLED"
	OrderPlaced   OrderStatus   = "ORDER PLACED"
	Shipped       OrderStatus   = "SHIPPED"
	Delivered     OrderStatus   = "DELIVERED"
	OrderCanceled OrderStatus   = "ORDER CANCELLED"
)

type Order struct {
	ID          uuid.UUID   `gorm:"type:char(255);primaryKey"`
	UserID      uint        `gorm:"not null"`
	ProductID   string      `gorm:"not null"`
	Quantity    int         `gorm:"not null;default:1"`
	TotalAmount float64     `gorm:"not null"` // price * quantity
	OrderStatus OrderStatus `gorm:"type:varchar(20);default:'ORDER PLACED'"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
}

type OrderCreateDTO struct {
	UserID      uint    `gorm:"not null"`
	ProductID   string  `gorm:"not null"`
	Quantity    int     `gorm:"not null;default:1"`
	TotalAmount float64 `gorm:"not null"`
}
