package parkinglot

import "errors"

type Spot struct {
	id            int64
	levelId       int64
	vehicleType   VehicleType
	isOccupied    bool
	parkedVehicle Vehicle
}

var spotAutoIncrementId int64 = 0

func NewSpot(levelId int64, vType VehicleType) *Spot {
	spotAutoIncrementId += 1
	return &Spot{
		id:            spotAutoIncrementId,
		levelId:       levelId,
		vehicleType:   vType,
		isOccupied:    false,
		parkedVehicle: nil,
	}
}

func (s *Spot) Park(v Vehicle) error {
	if s.isOccupied == false && s.vehicleType == v.getType() {
		s.parkedVehicle = v
		s.isOccupied = true
		return nil
	}
	return errors.New("Cannot park")
}

func (s *Spot) UnPark(v Vehicle) error {
	if s.isOccupied == true && s.vehicleType == v.getType() && s.parkedVehicle.getId() == v.getId() {
		s.parkedVehicle = nil
		s.isOccupied = false
		return nil
	}
	return errors.New("Cannot unpark")
}
