package isbnverifier
import "strings"

func IsValidISBN(isbn string) bool {
	s := strings.ReplaceAll(isbn, "-", "")
    if len(s) != 10 {
        return false
    }
    sum := 0
    for i := 0; i < len(s); i++ {
        c := s[i]
        var digit int
        switch {
        case c >= '0' && c <= '9':
        	digit = int(c - '0')
        case c == 'X' && i == 9:
        	digit = 10
        default:
            return false
        }
        sum += digit * (10 - i)
    }
    return sum % 11 == 0
}
