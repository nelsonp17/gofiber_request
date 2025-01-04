package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Fields map[string]string
	Form   map[string]interface{}
	Errors map[string]string
}
type Interface interface {
	Get(key string) interface{}
	GetString(key string) string
	GetInt(key string) int
	GetInt64(key string) int64
	GetFloat(key string) float64
	GetFloat32(key string) float32
	GetBool(key string) bool
	_GetFields() []string
	_GetRulesField(field string) []string
	Validated() bool
	Start(c *fiber.Ctx) (Request, error)
	GetArray(key string) []string
}

func (r Request) GetString(key string) string {
	valor := r.Form[key]
	//fmt.Println("Tipo es:", reflect.TypeOf(valor))

	if valor == nil {
		return ""
	}

	if reflect.TypeOf(valor).String() == "string" {
		return valor.(string)
	}
	if reflect.TypeOf(valor).String() == "float64" {
		return fmt.Sprintf("%v", valor)
	}
	if reflect.TypeOf(valor).String() == "int" {
		return fmt.Sprintf("%v", valor)
	}
	if reflect.TypeOf(valor).String() == "int64" {
		return fmt.Sprintf("%v", valor)
	}
	if reflect.TypeOf(valor).String() == "bool" {
		return fmt.Sprintf("%v", valor)
	}

	return fmt.Sprintf("%v", valor)
}

func (r Request) GetInt(key string) int {
	valor := r.Form[key]
	//fmt.Println("Tipo es:", reflect.TypeOf(valor))

	if valor == nil {
		return 0
	}

	switch v := valor.(type) {
	case string:
		if intValue, err := strconv.Atoi(v); err == nil {
			return intValue
		}
	case float64:
		return int(v)
	case float32:
		return int(v)
	case int:
		return v
	case int64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	}

	return 0
}

func (r Request) GetInt64(key string) int64 {
	valor := r.Form[key]
	//fmt.Println("Tipo es:", reflect.TypeOf(valor))

	if valor == nil {
		return 0
	}

	switch v := valor.(type) {
	case string:
		if intValue, err := strconv.ParseInt(v, 10, 64); err == nil {
			return intValue
		}
	case float64:
		return int64(v)
	case float32:
		return int64(v)
	case int:
		return int64(v)
	case int64:
		return v
	case bool:
		if v {
			return 1
		}
		return 0
	}

	return 0
}

func (r Request) GetFloat(key string) float64 {
	valor := r.Form[key]
	//fmt.Println("Tipo es:", reflect.TypeOf(valor))

	if valor == nil {
		return 0
	}

	switch v := valor.(type) {
	case string:
		if floatValue, err := strconv.ParseFloat(v, 64); err == nil {
			return floatValue
		}
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	}

	return 0
}

func (r Request) GetFloat32(key string) float32 {
	valor := r.Form[key]
	//fmt.Println("Tipo es:", reflect.TypeOf(valor))

	if valor == nil {
		return 0
	}

	switch v := valor.(type) {
	case string:
		if floatValue, err := strconv.ParseFloat(v, 32); err == nil {
			return float32(floatValue)
		}
	case float64:
		return float32(v)
	case float32:
		return v
	case int:
		return float32(v)
	case int64:
		return float32(v)
	case bool:
		if v {
			return 1
		}
		return 0
	}

	return 0
}

func (r Request) GetBool(key string) bool {
	valor := r.Form[key]
	//fmt.Println("Tipo es:", reflect.TypeOf(valor))

	if valor == nil {
		return false
	}

	switch v := valor.(type) {
	case string:
		boolValue, err := strconv.ParseBool(v)
		if err == nil {
			return boolValue
		}
	case bool:
		return v
	case int:
		return v != 0
	case int64:
		return v != 0
	case float64:
		return v != 0
	}

	return false
}

func (r Request) Get(key string) interface{} {
	valor := r.Form[key]
	//fmt.Println("Tipo es:", reflect.TypeOf(valor))

	if valor == nil {
		return ""
	}
	return valor
}

func (r Request) GetArray(key string) []string {
	valor := r.Form[key]
	//fmt.Println("Tipo es:", reflect.TypeOf(valor))

	if valor == nil {
		return nil
	}

	return FormattedArrayJson(valor)
}

func (r *Request) Validated() bool {
	// r.Form[field] = strings.TrimSpace(r.Form[field].(string))
	if r.Fields == nil || len(r.Fields) == 0 {
		return true
	}
	//fmt.Println("Form", r.Form)
	fieldsErrors := make(map[string]string)

	for field, rulesString := range r.Fields {
		if rulesString == "" {
			continue
		}
		rules := r._GetRulesField(field)

		for _, rule := range rules {

			value := r.GetString(field)

			parts := strings.Split(rule, ":")
			ruleOrigin := rule
			if len(parts) == 2 {
				// min o max
				rule = parts[0]
			}

			var err error
			switch rule {
			case "required":
				err = RuleRequired(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "email":
				err = RuleEmail(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "dni":
				err = RuleDni(value)
				if err != nil {
					fmt.Println("Error en dni", field, err)
					fieldsErrors[field] = err.Error()
				}
			case "phone":
				err = RulePhone(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "url":
				err = RuleUrl(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "date":
				err = RuleDate(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "datetime":
				err = RuleDatetime(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "time":
				err = RuleTime(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "integer":
				err = RuleInteger(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "float":
				err = RuleFloat(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "password":
				err = RulePassword(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "min":
				minLength, err := GetDigitFromRule(ruleOrigin)
				if err == nil {
					err = RuleMin(value, minLength)
				}
				if err != nil {
					fmt.Println("Error en min", field, err)
					fieldsErrors[field] = err.Error()
				}
			case "max":
				maxLength, err := GetDigitFromRule(ruleOrigin)
				if err == nil {
					err = RuleMax(value, maxLength)
				}
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			case "price":
				err = RulePrice(value)
				if err != nil {
					fieldsErrors[field] = err.Error()
				}
			}
		}
	}

	if len(fieldsErrors) > 0 {
		r.Errors = fieldsErrors
		return false
	}
	return true

}

func (r *Request) Start(c *fiber.Ctx) error {
	data := r._GetFields()
	dataResponse := make(map[string]interface{}, len(data))
	body := c.Request().Body()

	// Attempt to decode as JSON
	if json.Valid(body) {
		// Convertir []byte a map[string]interface{}
		err := json.Unmarshal(body, &dataResponse)
		if err != nil {
			//fmt.Println("Error al convertir JSON:", err)
			return errors.New("error al convertir JSON")
		}

		r.Form = dataResponse
		//fmt.Println("DataResponse 1", dataResponse)
		//fmt.Println("r.Form 1", r.Form)
		return nil
	}

	// Si no es JSON, intentamos como x-www-form-urlencoded
	for _, param := range data {
		dataResponse[param] = c.FormValue(param)
	}

	r.Form = dataResponse
	//fmt.Println("DataResponse", dataResponse)
	//fmt.Println("r.Form", r.Form)
	return nil
}

func (r Request) _GetFields() []string {
	_fields := make([]string, 0, len(r.Fields))
	for key := range r.Fields {
		_fields = append(_fields, key)
	}
	return _fields
}
func (r Request) _GetRulesField(field string) []string {
	if r.Fields[field] != "" {
		return strings.Split(r.Fields[field], "|")
	}
	return nil
}
