package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	Cards := []int{2, 6, 9}
    return Cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if len(slice) > index && index >= 0 {
        return slice[index]
    }
    return -1
}

// SetItem writes an item to a slice at a given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice) > index && index >= 0 {
	    slice[index] = value
        return slice
	}
    newslice := append(slice, value)
    return newslice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    slice1 := append(values, slice...)
    return slice1
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if len(slice) > index && index >= 0 {
	    slice1 := slice[:index]
    	slice2 := slice[index + 1:]
    	newslice := append(slice1, slice2...)
    	return newslice
	}
    return slice
}
