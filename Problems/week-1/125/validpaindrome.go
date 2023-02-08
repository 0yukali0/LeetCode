import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
    f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
    input := strings.ToLower(strings.Join(strings.FieldsFunc(s, f), ""))
    if len(input) == 0 {
        return true
    }

    for i, j := 0, len(input) - 1; i <= j; {
        if input[i] != input[j] {
            return false
        }
        i++
        j--
    }
    return true
}
