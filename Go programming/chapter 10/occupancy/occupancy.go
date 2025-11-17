package occupancy
func calculateOccupancy(occupancy float32) string{
    // if occupancy > 70 {
    //     return "High"
    // } else if occupancy > 20 {
    //     return "Medium"
    // }else{
    //     return "Low"
    // }
    switch {
    case occupancy > 70:
        return "High"
    case occupancy > 20:
        return "Medium"
    default:
        return "Low"
    }
}
func calculateOccupancyRate(roomsOccupied int, totalRooms int ) float32 {
    occupancyRate := (float32(roomsOccupied) / float32(totalRooms)) * 100
    return occupancyRate
}