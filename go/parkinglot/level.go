package parkinglot

type Level struct {
	id    int64
	lotId int64
	spots []*Spot
}

var levelAutoIncrementId int64 = 0

func NewLevel(lotId int64) *Level {
	levelAutoIncrementId += 1
	return &Level{
		id:    levelAutoIncrementId,
		lotId: lotId,
		spots: make([]*Spot, 0),
	}
}

func (l *Level) ParkVehicle(v Vehicle) error {
	var err error = nil
	for j := 0; j < len(l.spots); j++ {
		err = l.spots[j].Park(v)
		if err == nil {
			return nil
		}
	}
	return err
}

func (l *Level) UnParkVehicle(v Vehicle) error {
	var err error = nil
	for j := 0; j < len(l.spots); j++ {
		err = l.spots[j].UnPark(v)
		if err == nil {
			return nil
		}
	}
	return err
}
