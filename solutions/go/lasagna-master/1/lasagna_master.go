package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
    if minutes == 0 {
        return len(layers) * 2
    }
    return len(layers) * minutes
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    n := 0
    s := 0.0
    for a := 0; a < len(layers); a++ {
        switch {
        case layers[a] == "noodles":
            n++
        case layers[a] == "sauce":
            s++
        default:
        }
    }
    noodles = n * 50
    sauce = s * 0.2
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendslist []string, mylist []string) {
    mylist[len(mylist) - 1] = friendslist[len(friendslist) - 1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(oldquantities []float64, people int) []float64 {
    newquantities := make([]float64, len(oldquantities))
    for a := 0; a < len(newquantities); a++ {
        newquantities[a] = oldquantities[a] * float64(people) / 2
    }
    return newquantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
