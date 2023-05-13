package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func StructsToString[E any](elements []E, length int) []string {
	if len(elements) == 0 {
		return []string{"The list is empty."}
	}

	var res []string
	currentString := ""
	for _, element := range elements {
		strStruct := fmt.Sprintf(
			"----------\n%s----------\n",
			structToString(reflect.ValueOf(element)),
		)
		if len(currentString)+len(strStruct) < length {
			currentString += strStruct
		} else {
			res = append(res, currentString)
			currentString = strStruct
		}
	}

	if currentString != "" {
		res = append(res, currentString)
	}
	return res
}

func structToString(val reflect.Value) string {
	var res string
	for i := 0; i < val.NumField(); i++ {
		if strVal := valToString(val.Field(i)); strVal != "" {
			res += fmt.Sprintf("%s: %s\n", val.Type().Field(i).Name, strVal)
		}
	}
	return res
}

func valToString(val reflect.Value) string {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.String:
		return val.String()
	default:
		return ""
	}
}
