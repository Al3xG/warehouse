package model

import "errors"

type Warehouse struct {
	ID       int
	Name     string
	Location string
	Capacity int
	Velocity int
}

// куда вынести ошибки? т.к. некоторые повторяются у разных моделей
var (
	ErrNameRequired        = errors.New("name is required")
	ErrLocationRequired    = errors.New("location is required")
	ErrWarehouseOverloaded = errors.New("warehouse is overloaded")
	ErrCapacityRequired    = errors.New("capacity is required and need > 0")
)

// SHIFT + F6 refactor

func NewWarehouse(id int, name string, location string, capacity, velocity int) (*Warehouse, error) {
	if name == "" {
		return nil, ErrNameRequired
	}
	if capacity <= 0 {
		return nil, ErrCapacityRequired
	}
	if location == "" {
		return nil, ErrLocationRequired
	}
	if velocity > capacity {
		return nil, ErrWarehouseOverloaded
	}

	return &Warehouse{
		ID:       id,
		Name:     name,
		Location: location,
		Capacity: capacity,
		Velocity: velocity,
	}, nil
}
