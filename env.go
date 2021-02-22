package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrNotStructPointer = errors.New("it is not pointer to struct")
	ErrNotSupportedType = errors.New("this field is not support type")

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

type Option struct {
	EnvPrefix string
}
type Options []Option

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

func doParse(ref reflect.Value, options ...Option) error {
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		refField := ref.Field(i)
		refStructField := refType.Field(i)

		if sf, ok := isStruct(refField, refStructField); ok {
			if err := doParse(sf, append(options, Option{
				EnvPrefix: getFieldKey(refStructField),
			})...); err != nil {
				return err
			}
		}

		value, err := getValueFromEnv(refStructField, options)
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

func getValueFromEnv(field reflect.StructField, options Options) (string, error) {
	key := strings.Join(append(options.EnvPrefixes(), getFieldKey(field)), "_")
	val := os.Getenv(key)
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

func getFieldKey(structField reflect.StructField) string {
	if key := structField.Tag.Get("env_key"); key != "" {
		return key
	}
	return toSnakeCase(structField.Name)
}

func toSnakeCase(input string) string {
	n := strings.Builder{}
	for i, v := range []byte(input) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if vIsLow {
			v += 'A'
			v -= 'a'
		}
		if vIsCap && i > 0 {
			n.WriteByte('_')
		}
		n.WriteByte(v)
	}
	return n.String()
}

func isStruct(field reflect.Value, structField reflect.StructField) (reflect.Value, bool) {
	switch structField.Type.Kind() {
	case reflect.Ptr:
		if structField.Type.Elem().Kind() != reflect.Struct {
			return field, false
		}
		if field.IsNil() {
			field.Set(reflect.New(structField.Type.Elem()))
		}
		return field.Elem(), true
	case reflect.Struct:
		return field, true
	default:
		return field, false
	}
}

func (ops Options) EnvPrefixes() []string {
	var keyElms []string
	for i := range ops {
		op := ops[i]
		if op.EnvPrefix == "" {
			continue
		}
		keyElms = append(keyElms, op.EnvPrefix)
	}
	return keyElms
}
