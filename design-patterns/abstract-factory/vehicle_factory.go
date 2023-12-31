package main

import (
	f "design-pattern/design-patterns/abstract-factory/factory"
	c "design-pattern/design-patterns/abstract-factory/factory/car"
	m "design-pattern/design-patterns/abstract-factory/factory/motobike"
	"errors"
	"fmt"
)

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

type IVehicleFactory interface {
	GetVehicle(v int) (f.IVehicle, error)
}

func GetFactory(f int) (IVehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(c.Factory), nil
	case MotorbikeFactoryType:
		return new(m.Factory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d invalid.", f))
	}
}

func PrintVehicle(v f.IVehicle) {
	fmt.Printf("Seats: %d \n", v.GetSeats())
	fmt.Printf("Wheels: %d \n", v.GetWheels())
}
