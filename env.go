package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrNotStructPointer = errors.New("it is not pointer to struct")
	ErrNotSupportedType = errors.New("this field is not support type")

	matchFirstCap = regexp.MustCompile("([A-Z])([A-Z][a-z])")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

	convertFuncMap = map[reflect.Kind]func(v string) (interface{}, error){
		reflect.String: func(v string) (interface{}, error) {
			return v, nil
		},
		reflect.Int: func(v string) (interface{}, error) {
			return strconv.ParseInt(v, 10, 0)
		},
		reflect.Int8: func(v string) (interface{}, error) {
			return strconv.ParseInt(v, 10, 8)
		},
		reflect.Int16: func(v string) (interface{}, error) {
			return strconv.ParseInt(v, 10, 16)
		},
		reflect.Int32: func(v string) (interface{}, error) {
			return strconv.ParseInt(v, 10, 32)
		},
		reflect.Int64: func(v string) (interface{}, error) {
			return strconv.ParseInt(v, 10, 64)
		},
		reflect.Uint: func(v string) (interface{}, error) {
			return strconv.ParseUint(v, 10, 0)
		},
		reflect.Uint8: func(v string) (interface{}, error) {
			return strconv.ParseUint(v, 10, 8)
		},
		reflect.Uint16: func(v string) (interface{}, error) {
			return strconv.ParseUint(v, 10, 16)
		},
		reflect.Uint32: func(v string) (interface{}, error) {
			return strconv.ParseUint(v, 10, 32)
		},
		reflect.Uint64: func(v string) (interface{}, error) {
			return strconv.ParseUint(v, 10, 64)
		},
		reflect.Float32: func(v string) (interface{}, error) {
			return strconv.ParseFloat(v, 32)
		},
		reflect.Float64: func(v string) (interface{}, error) {
			return strconv.ParseFloat(v, 64)
		},
	}
)

func Parse(v interface{}) error {
	ref := reflect.ValueOf(v)
	if ref.Kind() != reflect.Ptr {
		return ErrNotStructPointer
	}
	if ref.Elem().Kind() != reflect.Struct {
		return ErrNotStructPointer
	}

	return doParse(ref.Elem())
}

func doParse(ref reflect.Value) error {
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		refField := ref.Field(i)
		refStructField := refType.Field(i)
		value, err := getValueFromEnv(refStructField)
		if err != nil {
			return err
		}
		if value == "" {
			continue
		}

		if err := setValue(refField, refStructField, value); err != nil {
			return err
		}
	}
	return nil
}

func getValueFromEnv(field reflect.StructField) (string, error) {
	val := os.Getenv(toSnakeCase(field.Name))
	return val, nil
}

func setValue(field reflect.Value, structField reflect.StructField, value string) error {
	conv, ok := convertFuncMap[structField.Type.Kind()]
	if !ok {
		return fmt.Errorf("%w [%v]", ErrNotSupportedType, structField.Type.Name())
	}

	val, err := conv(value)
	if err != nil {
		return err
	}

	field.Set(reflect.ValueOf(val).Convert(structField.Type))
	return nil
}

func toSnakeCase(input string) string {
	output := matchFirstCap.ReplaceAllString(input, "${1}_${2}")
	output = matchAllCap.ReplaceAllString(output, "${1}_${2}")
	output = strings.ReplaceAll(output, "-", "_")
	return strings.ToUpper(output)
}
