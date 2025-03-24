package finautils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func ValidateIntRef(intRef string) bool {
	base := intRef[4:]
	checkRef, err := CreateIntRef(base)
	if err != nil {
		return false
	}
	return intRef == checkRef
}

func CreateIntRef(base string) (string, error) {
	newBase := base + "RF00"
	newBaseStr := strings.Builder{}
	for i := 0; i < len(newBase); i++ {
		rn := rune(newBase[i])
		fmt.Println("Rune:", rn)
		if !unicode.IsDigit(rn) {
			newBaseStr.WriteString(converRuneToNumStr(rune(newBase[i])))
		} else {
			newBaseStr.WriteRune(rune(newBase[i]))
		}
	}
	iVal, err := strconv.Atoi(newBaseStr.String())
	if err != nil {
		return "", err
	}
	cs := (98 - (iVal % 97))

	intRef := fmt.Sprintf("RF%2d%s", cs, base)
	return intRef, nil
}

func converRuneToNumStr(val rune) string {
	return strconv.Itoa(int(val - 55))
}
