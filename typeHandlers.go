package crunchyTools

import (
	"strconv"
)

// ConvertIntBoolToString will take a int uint or a boolean and return a string
func ConvertIntBoolToString(value interface{}) string {
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 10)
	case bool:
		{
			if value == true {
				return "true"
			}
			return "false"
		}
	default:
		return ""
	}
}
