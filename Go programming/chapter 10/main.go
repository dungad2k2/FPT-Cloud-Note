package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    const hotelName = "Gopher Paris Inn"
    const totalRooms = 134
    const firstRoomNumber = 110

    rand.Seed(time.Now().UTC().UnixNano())
    roomsOccupied := rand.Intn(totalRooms)
    roomsAvailable := totalRooms - roomsOccupied

    occupancyRate := compute_occupancy_rate(roomsOccupied, totalRooms)
    occupancyLevel := compute_occupancy_level(occupancyRate)
    fmt.Println("Hotel:", hotelName)
    fmt.Println("Number of rooms", totalRooms)
    fmt.Println("Rooms available", roomsAvailable)
    fmt.Println("                  Occupancy Level:", occupancyLevel)
    fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate)

    if roomsAvailable > 0 {
        fmt.Println("Rooms:")
        for i := 0; roomsAvailable > i; i++ {
            roomNumber := firstRoomNumber + i
            size := rand.Intn(6) + 1
            nights := rand.Intn(10) + 1
            print_the_details_of_a_specific_room(roomNumber, size, nights)	
        }
    } else {
        fmt.Println("No rooms available for tonight")
    }

}
func compute_occupancy_rate(roomsOccupied int, totalRooms int) float64 {
	return (float64(roomsOccupied) / float64(totalRooms)) * 100
}
func compute_occupancy_level(occupancyRate float64) string {
	if occupancyRate > 70 {
		return "High"
	} else if occupancyRate > 20 {
		return "Medium"
	} else {
		return "Low"
	}
}
func print_the_details_of_a_specific_room(roomNumber int, size int, nights int) {
	fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
}