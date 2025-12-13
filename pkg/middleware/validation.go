package middleware

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(s interface{}) []ValidationError {
	var errors []ValidationError
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("validate")

		if tag == "" {
			continue
		}

		jsonTag := field.Tag.Get("json")
		fieldName := strings.Split(jsonTag, ",")[0]
		if fieldName == "" {
			fieldName = field.Name
		}

		errors = append(errors, validateField(fieldName, value, tag)...)
	}

	return errors
}

func validateField(fieldName string, value reflect.Value, tag string) []ValidationError {
	var errors []ValidationError
	rules := strings.Split(tag, ",")

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)
		if rule == "omitempty" && isEmpty(value) {
			continue
		}

		switch {
		case rule == "required":
			if isEmpty(value) {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("%s is required", fieldName),
				})
			}
		case strings.HasPrefix(rule, "min="):
			var min int
			fmt.Sscanf(rule, "min=%d", &min)
			if value.Kind() == reflect.String && len(value.String()) < min {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("%s must be at least %d characters", fieldName, min),
				})
			} else if value.Kind() == reflect.Slice && value.Len() < min {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("%s must have at least %d items", fieldName, min),
				})
			}
		case strings.HasPrefix(rule, "max="):
			var max int
			fmt.Sscanf(rule, "max=%d", &max)
			if value.Kind() == reflect.String && len(value.String()) > max {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("%s must be at most %d characters", fieldName, max),
				})
			}
		case rule == "email":
			if value.Kind() == reflect.String && !isValidEmail(value.String()) {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("%s must be a valid email", fieldName),
				})
			}
		case strings.HasPrefix(rule, "oneof="):
			options := strings.Split(strings.TrimPrefix(rule, "oneof="), " ")
			if value.Kind() == reflect.String && !contains(options, value.String()) {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("%s must be one of: %s", fieldName, strings.Join(options, ", ")),
				})
			}
		case strings.HasPrefix(rule, "gte="):
			var min int
			fmt.Sscanf(rule, "gte=%d", &min)
			if value.Kind() == reflect.Int && value.Int() < int64(min) {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("%s must be greater than or equal to %d", fieldName, min),
				})
			}
		case strings.HasPrefix(rule, "lte="):
			var max int
			fmt.Sscanf(rule, "lte=%d", &max)
			if value.Kind() == reflect.Int && value.Int() > int64(max) {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("%s must be less than or equal to %d", fieldName, max),
				})
			}
		}
	}

	return errors
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Slice, reflect.Array:
		return v.Len() == 0
	case reflect.Ptr:
		return v.IsNil()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	default:
		return false
	}
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ValidateJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

