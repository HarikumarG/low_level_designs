package parkinglot

import "errors"

type VehicleType int

type Vehicle interface {
	getType() VehicleType
	getId() string
}

const (
	CAR        VehicleType = 0
	MOTORCYCLE VehicleType = 1
	TRUCK      VehicleType = 2
)

func GetVehicle(input map[string]interface{}) (Vehicle, error) {
	switch input["vType"] {
	case CAR:
		{
			return &Car{
				vType:     CAR,
				vehicleId: input["id"].(string),
			}, nil
		}
	case MOTORCYCLE:
		{
			return &MotorCyle{
				vType:     MOTORCYCLE,
				vehicleId: input["id"].(string),
			}, nil
		}
	case TRUCK:
		{
			return &Truck{
				vType:     TRUCK,
				vehicleId: input["id"].(string),
			}, nil
		}
	}
	return nil, errors.New("Invalid vehicle type")
}

type Car struct {
	vType     VehicleType
	vehicleId string
}

func (c *Car) getType() VehicleType {
	return c.vType
}

func (c *Car) getId() string {
	return c.vehicleId
}

type MotorCyle struct {
	vType     VehicleType
	vehicleId string
}

func (c *MotorCyle) getType() VehicleType {
	return c.vType
}

func (c *MotorCyle) getId() string {
	return c.vehicleId
}

type Truck struct {
	vType     VehicleType
	vehicleId string
}

func (c *Truck) getType() VehicleType {
	return c.vType
}

func (c *Truck) getId() string {
	return c.vehicleId
}
