package models

import (
	"time"
)

type Queryable interface {
	GetKeyValidators() map[string]func(interface{}) bool
}

// ValidateString checks if the provided value is a string.
func ValidateString(val interface{}) bool {
	switch v := val.(type) {
	case string:
		return true
	case *string:
		return v != nil
	default:
		return false
	}
}

// ValidateString checks if the provided value is a time.Time.
func ValidateTime(val interface{}) bool {
	switch v := val.(type) {
	case time.Time:
		return true
	case *time.Time:
		return v != nil
	default:
		return false
	}
}

// ValidateInt checks if the provided value is an int.
func ValidateInt(val interface{}) bool {
	switch v := val.(type) {
	case int:
		return true
	case *int:
		return v != nil
	default:
		return false
	}
}

// ValidateBool checks if the provided value is a bool.
func ValidateBool(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return true
	case *bool:
		return v != nil
	default:
		return false
	}
}
