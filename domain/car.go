package domain

import (
	"context"
	"time"
)

//Car ...
type Car struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name" validate:"required"`
	Brand         string    `json:"brand" validate:"required"`
	Price         string    `json:"price"`
	Condition     string    `json:"condition"`
	Quantity      int64     `json:"quantity"`
	Description   string    `json:"description"`
	Specification string    `json:"specification"`
	ImageURL      string    `json:"imageURL"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

//CarUsecase represent the Car's usecases
type CarUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Car, string, error)
	GetByID(ctx context.Context, id int64) (Car, error)
	Update(ctx context.Context, ar *Car) error
	GetByName(ctx context.Context, title string) (Car, error)
	Store(context.Context, *Car) error
	Delete(ctx context.Context, id int64) error
}

// CarRepository represent the Car's repository contract
type CarRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Car, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Car, error)
	GetByName(ctx context.Context, name string) (Car, error)
	Update(ctx context.Context, ar *Car) error
	Store(ctx context.Context, a *Car) error
	Delete(ctx context.Context, id int64) error
}
