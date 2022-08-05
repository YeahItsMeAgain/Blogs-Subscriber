package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func StructsToString[E comparable](elements []E) string {
	var res string

	for _, element := range elements {
		res += "----------\n"
		val := reflect.ValueOf(element)
		for i := 0; i < val.NumField(); i++ {
			strVal, err := valToString(val.Field(i))
			if err != nil {
				continue
			}
			res += fmt.Sprintf(
				"%s: %s\n",
				val.Type().Field(i).Name, strVal,
			)
		}
		res += "----------\n"
	}
	return res
}

func valToString(val reflect.Value) (string, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10), nil
	case reflect.String:
		return val.String(), nil
	}
	return "", errors.New("not primitive")
}
