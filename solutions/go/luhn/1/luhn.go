package luhn

import "strings"

func Valid(id string) bool {
	trueid := strings.ReplaceAll(id, " ", "")
    if len(trueid) <= 1 {
        return false
    }
    for _, r := range trueid {
        if r > '9' || r < '0' {
            return false
        }
    }
    sum := 0
    double := false
    for i := len(trueid) - 1; i >= 0; i-- {
        digit := int(trueid[i] - '0')
        if double {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
        double = !double
    }
    return sum % 10 == 0
}
