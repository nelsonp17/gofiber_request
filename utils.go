package request

import (
	"fmt"
	"strconv"
	"strings"
)




func GetDigitFromRule(rule string) (int, error) {
	parts := strings.Split(rule, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid rule format")
	}
	digit, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid digit: %v", err)
	}
	return digit, nil
}
