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

func FormattedArrayJson(to interface{}) []string {
	switch v := to.(type) {
	case string:
		fmt.Println("Caso 1 y 2: string o string separado por coma: ", v)
		// Caso 1 y 2: string o string separado por coma
		return strings.Split(v, ",")
	case []string:
		// Caso 3: un tipo desconocido que se imprime en consola asi [email1 email2]
		fmt.Println("Caso 3: un tipo desconocido que se imprime en consola asi [email1 email2]. ", v)
		return v
	default:
		// Caso por defecto: convertir a string y dividir por espacios
		fmt.Println("Caso por defecto: convertir a string y dividir por espacios: ", v)

		v2 := fmt.Sprint(v)
		v2 = strings.Replace(v2, "[", "", -1)
		v2 = strings.Replace(v2, "]", "", -1)

		return strings.Fields(v2)
	}
}
