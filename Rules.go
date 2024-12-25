package request

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	ErrInvalidPassword = errors.New("the password must has at least 8 of lenght and contain one special and numeric character")
	ErrInvalidPhone    = errors.New("invalid phone number")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidDni      = errors.New("invalid DNI or NIF")
	ErrRequiredField   = errors.New("this field is required")
	ErrMinLength       = errors.New("the field does not meet the minimum length requirement")
	ErrMaxLength       = errors.New("the field exceeds the maximum length requirement")
	ErrInvalidPrice    = errors.New("the price format is invalid")
	ErrNotUnique       = errors.New("the value is not unique in the database")
	ErrInvalidUrl      = errors.New("invalid URL")
	ErrInvalidInteger  = errors.New("invalid integer")
	ErrInvalidFloat    = errors.New("invalid float")
	ErrInvalidDate     = errors.New("invalid date format")
	ErrInvalidDatetime = errors.New("invalid datetime format")
	ErrInvalidTime     = errors.New("invalid time format")
)

// RuleRequired checks if a string is not empty
func RuleRequired(value string) error {
	if value == "" {
		return ErrRequiredField
	}
	return nil
}

// RuleMin checks if a string meets the minimum length requirement
func RuleMin(value string, minLength int) error {
	if len(value) < minLength {
		return fmt.Errorf("%w: minimum length is %d", ErrMinLength, minLength)
	}
	return nil
}

// RuleMax checks if a string does not exceed the maximum length requirement
func RuleMax(value string, maxLength int) error {
	if len(value) > maxLength {
		return fmt.Errorf("%w: maximum length is %d", ErrMaxLength, maxLength)
	}
	return nil
}

// RulePrice checks if a string is a valid price format
func RulePrice(value string) error {
	validPrice, err := regexp.MatchString(`^\d+(\.\d{1,2})?$`, value)
	if err != nil {
		return fmt.Errorf("price pattern: %v", err)
	}
	if !validPrice {
		return ErrInvalidPrice
	}
	return nil
}

// RuleUnique checks if a value is unique in the specified table and column
func RuleUnique(db *sql.DB, tableName, columnName, value string) error {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", tableName, columnName)
	var count int
	err := db.QueryRow(query, value).Scan(&count)
	if err != nil {
		return fmt.Errorf("database query error: %v", err)
	}
	if count > 0 {
		return ErrNotUnique
	}
	return nil
}

// RuleEmail validate if email is not empty or have a bad format.
func RuleEmail(email string) error {
	if email == "" {
		return ErrInvalidEmail
	}

	validEmail, err := regexp.MatchString(`^([a-zA-Z0-9])+([a-zA-Z0-9\._-])*@([a-zA-Z0-9_-])+(\.)+([a-zA-Z0-9_-]+)$`, email)
	if err != nil {
		return fmt.Errorf("email pattern: %v", err)
	}

	if !validEmail {
		return ErrInvalidEmail
	}

	return nil
}

func RulePhone(phone string) error {
	if phone == "" {
		return ErrInvalidPhone
	}

	validPhone, err := regexp.MatchString(`^\+[1-9]\d{1,14}$`, phone)
	if err != nil {
		return fmt.Errorf("phone pattern: %v", err)
	}

	if !validPhone {
		return ErrInvalidPhone
	}

	return nil
}

// Cabe destacar que esto lo saque de https://es.wikipedia.org/wiki/N%C3%BAmero_de_identificaci%C3%B3n_fiscal
func RuleDni(dni string) error {
	//Si el largo del NIF es diferente a 9, acaba el método.
	if len(dni) != 9 {
		return ErrInvalidDni
	}

	secuenciaLetrasDni := "TRWAGMYFPDXBNJZSQVHLCKE"
	dni = strings.ToUpper(dni)

	matchFirstLetter := false
	matchLastLetter := false
	// letterCount := 0
	// fmt.Println(string(nif[0]), nif[0], nif[0] < 48 || nif[0] > 57)
	if dni[0] < 48 || dni[0] > 57 {
		for _, c := range secuenciaLetrasDni {
			if dni[0] == byte(c) {
				matchFirstLetter = true
				break
			}
		}
		if !matchFirstLetter {
			return ErrInvalidDni
		}
	}

	if dni[8] < 48 || dni[8] > 57 {
		for _, c := range secuenciaLetrasDni {
			if dni[8] == byte(c) {
				matchLastLetter = true
				break
			}
		}
		if !matchLastLetter {
			return ErrInvalidDni
		}
	}

	// if letterCount == 0 {
	// }

	for _, c := range dni[1:8] {
		if c < 48 || c > 57 {
			return ErrInvalidDni
		}
	}

	return nil

}

func RulePassword(password string) error {
	hasNumber := false
	hasSpecial := false
	hasLetter := false

	if len(password) < 8 {
		return ErrInvalidPassword
	}

	for _, c := range password {
		switch {
		case unicode.IsLetter(c):
			hasLetter = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if hasNumber && hasSpecial && hasLetter {
		return nil
	}

	return ErrInvalidPassword
}

// Cabe destacar que esto lo saque de https://es.wikipedia.org/wiki/N%C3%BAmero_de_identificaci%C3%B3n_fiscal
func RuleNifFormat(nif string) (bool, error) {

	//Si el largo del NIF es diferente a 9, acaba el método.
	if len(nif) != 9 {
		return false, ErrInvalidDni
	}

	secuenciaLetrasNIF := "TRWAGMYFPDXBNJZSQVHLCKE"
	nif = strings.ToUpper(nif)

	matchFirstLetter := false
	matchLastLetter := false
	// fmt.Println(string(nif[0]), nif[0], nif[0] < 48 || nif[0] > 57)
	if nif[0] < 48 || nif[0] > 57 {
		for _, c := range secuenciaLetrasNIF {
			if nif[0] == byte(c) {
				matchFirstLetter = true
				break
			}
		}
		if !matchFirstLetter {
			return false, ErrInvalidDni
		}
	}

	if nif[8] < 48 || nif[8] > 57 {
		for _, c := range secuenciaLetrasNIF {
			if nif[8] == byte(c) {
				matchLastLetter = true
				break
			}
		}
		if !matchLastLetter {
			return false, ErrInvalidDni
		}
	}

	return true, nil
}

// RuleUrl checks if a string is a valid URL
func RuleUrl(value string) error {
	validUrl, err := regexp.MatchString(`^(http|https):\/\/[^\s$.?#].[^\s]*$|^www\.[^\s$.?#].[^\s]*$`, value)
	if err != nil {
		return fmt.Errorf("URL pattern: %v", err)
	}
	if !validUrl {
		return ErrInvalidUrl
	}
	return nil
}

// RuleInteger checks if a string is a valid integer
func RuleInteger(value string) error {
	if _, err := strconv.Atoi(value); err != nil {
		return ErrInvalidInteger
	}
	return nil
}

// RuleFloat checks if a string is a valid float
func RuleFloat(value string) error {
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		return ErrInvalidFloat
	}
	return nil
}

// RuleDate checks if a string is a valid date in the format YYYY-MM-DD
func RuleDate(value string) error {
	if _, err := time.Parse("2006-01-02", value); err != nil {
		return ErrInvalidDate
	}
	return nil
}

// RuleDatetime checks if a string is a valid datetime in the format YYYY-MM-DD HH:MM:SS
func RuleDatetime(value string) error {
	if _, err := time.Parse("2006-01-02 15:04:05", value); err != nil {
		return ErrInvalidDatetime
	}
	return nil
}

// RuleTime checks if a string is a valid time in the format HH:MM:SS
func RuleTime(value string) error {
	if _, err := time.Parse("15:04:05", value); err != nil {
		return ErrInvalidTime
	}
	return nil
}
