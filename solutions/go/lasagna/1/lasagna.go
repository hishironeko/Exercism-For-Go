package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40
var time1, time2, time3 int
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	time1 := OvenTime - actualMinutesInOven
    return time1
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	time2 := numberOfLayers * 2
    return time2
}
// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	time3 := numberOfLayers * 2 + actualMinutesInOven
    return time3
}
