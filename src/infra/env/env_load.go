package env

import (
	"fmt"
	"reflect"
)

func LoadEnv(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		panic("LoadEnv requires a non-nil pointer to a struct")
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		panic("LoadEnv requires a pointer to a struct")
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanSet() {
			continue
		}

		structField := typ.Field(i)

		envKey := structField.Tag.Get("field")
		if envKey == "" {
			envKey = structField.Name
		}

		env := GetEnv(envKey)

		if defaultValue := structField.Tag.Get("default"); defaultValue != "" {
			env.Default(defaultValue)
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(env.AsString())
		case reflect.Int:
			field.SetInt(int64(env.AsInt()))
		case reflect.Int64:
			field.SetInt(env.AsInt64())
		case reflect.Float64:
			field.SetFloat(env.AsFloat64())
		case reflect.Bool:
			field.SetBool(env.AsBool())
		case reflect.Slice:
			if field.Type().Elem().Kind() == reflect.String {
				field.Set(reflect.ValueOf(env.GetEnvAsStrings(",")))
			} else {
				panic(fmt.Sprintf("Unsupported slice type for field %s: %s", structField.Name, field.Type().Elem().Kind()))
			}
		default:
			panic(fmt.Sprintf("Unsupported type for field %s: %s", structField.Name, field.Kind()))
		}
	}
}
