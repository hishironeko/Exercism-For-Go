package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than 0")
	}
	var steps int
	for a := n; a != 1; steps++ {
		if a % 2 == 0 {
			a = a / 2
		} else {
			a = 3 * a + 1
		}
	}
	return steps, nil
}