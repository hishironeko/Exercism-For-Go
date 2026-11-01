package darts
import "math"

func Score(x, y float64) int {
	point := math.Sqrt(x * x + y * y)
    var score int
    switch {
    case point > 10.0 :
        score = 0
    case point <= 10.0 && point > 5.0 :
        score = 1
    case point <= 5.0 && point > 1.0:
        score = 5
    case point <= 1.0 && point >= 0.0:
        score = 10
    default:
    }
    return score
}
