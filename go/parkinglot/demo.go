package parkinglot

import "fmt"

var input = []map[string]interface{}{
	{
		"id":        "car1",
		"vType":     CAR,
		"operation": ENTRY,
	},
	{
		"id":        "motorcycle1",
		"vType":     MOTORCYCLE,
		"operation": ENTRY,
	},
	{
		"id":        "truck1",
		"vType":     TRUCK,
		"operation": ENTRY,
	},
	{
		"id":        "car2",
		"vType":     CAR,
		"operation": ENTRY,
	},
	{
		"id":        "motorcycle2",
		"vType":     MOTORCYCLE,
		"operation": ENTRY,
	},
	{
		"id":        "truck2",
		"vType":     TRUCK,
		"operation": ENTRY,
	},
	{
		"id":        "car3",
		"vType":     CAR,
		"operation": ENTRY,
	},
	{
		"id":        "motorcycle2",
		"vType":     MOTORCYCLE,
		"operation": EXIT,
	},
	{
		"id":        "truck3",
		"vType":     TRUCK,
		"operation": ENTRY,
	},
	{
		"id":        "motorcycle3",
		"vType":     MOTORCYCLE,
		"operation": ENTRY,
	},
}

func ParkingLotDemo() {
	lot := NewParkingLot()
	level1 := NewLevel(lot.id)
	level1.spots = append(level1.spots, NewSpot(level1.id, CAR), NewSpot(level1.id, MOTORCYCLE), NewSpot(level1.id, TRUCK))
	level2 := NewLevel(lot.id)
	level2.spots = append(level2.spots, NewSpot(level2.id, CAR), NewSpot(level2.id, MOTORCYCLE), NewSpot(level2.id, TRUCK))
	lot.levels = append(lot.levels, level1, level2)

	lot.PrintStatus()

	for i := 0; i < len(input); i++ {
		v, err := GetVehicle(input[i])
		if err != nil {
			fmt.Printf("Invalid vehicle type\n")
			continue
		}
		fmt.Printf("Executing operation %v : %v : %v\n", v.getId(), v.getType(), input[i]["operation"])
		switch input[i]["operation"] {
		case ENTRY:
			{
				err := lot.Entry(v)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		case EXIT:
			{
				err := lot.Exit(v)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
		lot.PrintStatus()
	}
}
