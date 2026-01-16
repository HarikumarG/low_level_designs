package parkinglot

import (
	"fmt"
)

type Operation int

const (
	ENTRY Operation = 1
	EXIT  Operation = 0
)

type ParkingLot struct {
	id     int64
	levels []*Level
}

var lotAutoIncrementId int64 = 0

func NewParkingLot() *ParkingLot {
	lotAutoIncrementId += 1
	return &ParkingLot{
		id:     lotAutoIncrementId,
		levels: make([]*Level, 0),
	}
}

func (l *ParkingLot) PrintStatus() {
	fmt.Printf("Parking Lot - %d\n\n", l.id)
	for i := 0; i < len(l.levels); i++ {
		fmt.Printf("Level ID - %d\n", l.levels[i].id)
		for j := 0; j < len(l.levels[i].spots); j++ {
			fmt.Printf("Spot - %d, %d, %v\n", l.levels[i].spots[j].id, l.levels[i].spots[j].vehicleType, l.levels[i].spots[j].isOccupied)
		}
		fmt.Printf("\n")
	}
}

func (l *ParkingLot) Entry(v Vehicle) error {
	err := l.parkVehicle(v)
	if err != nil {
		return err
	}
	return nil
}

func (l *ParkingLot) Exit(v Vehicle) error {
	err := l.unParkVehicle(v)
	if err != nil {
		return err
	}
	return nil
}

func (l *ParkingLot) parkVehicle(v Vehicle) error {
	var err error = nil
	for i := 0; i < len(l.levels); i++ {
		err = l.levels[i].ParkVehicle(v)
		if err == nil {
			return nil
		}
	}
	return err
}

func (l *ParkingLot) unParkVehicle(v Vehicle) error {
	var err error = nil
	for i := 0; i < len(l.levels); i++ {
		err = l.levels[i].UnParkVehicle(v)
		if err == nil {
			return nil
		}
	}
	return nil
}
