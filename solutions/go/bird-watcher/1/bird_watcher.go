package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for day := 0; day < len(birdsPerDay); day++ {
        sum += birdsPerDay[day]
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    daystart := (week - 1) * 7
    dayend := daystart + 7
    sum := 0
    for day := daystart; day < dayend; day++ {
        sum = sum + birdsPerDay[day]
    }
    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for day := 0; day < len(birdsPerDay); day += 2 {
        birdsPerDay[day]++
    }
    return birdsPerDay
}
