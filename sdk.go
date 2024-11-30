package sdk

// display functions

import "fmt"

func Display(value interface{}) string {
	switch v := value.(type) {
	case []interface{}:
		return fmt.Sprintf("%v", v)
	case map[string]interface{}:
		return ""
	case string:
		return v
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// check if value is a map
func IsMap(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}

// check if value is a slice
func IsSlice(value interface{}) bool {
	_, ok := value.([]interface{})
	return ok
}
