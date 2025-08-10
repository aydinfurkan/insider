package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Env struct {
	Key   string
	Value any
}

func GetEnv(key string) *Env {
	lazyInit()

	if value, exists := os.LookupEnv(key); exists {
		return &Env{
			Key:   key,
			Value: value,
		}
	}

	return &Env{
		Key:   key,
		Value: nil,
	}
}

func (e *Env) Default(v interface{}) *Env {

	if e.Value == nil {
		e.Value = v
	}

	return e
}

func (e *Env) AsString() string {
	if e.Value == nil {
		panic(fmt.Sprintf("Environment variable %s not found on env or default value is not set", e.Key))
	}
	if v, ok := e.Value.(string); ok {
		return v
	}
	panic(fmt.Sprintf("Environment variable %s is not a string", e.Key))
}

func (e *Env) AsInt() int {
	if e.Value == nil {
		panic(fmt.Sprintf("Environment variable %s not found on env or default value is not set", e.Key))
	}

	if v, ok := e.Value.(int); ok {
		return v
	}

	if v, ok := e.Value.(string); ok {
		if intValue, err := strconv.Atoi(v); err == nil {
			return intValue
		}
	}

	panic(fmt.Sprintf("Environment variable %s is not an int", e.Key))
}

func (e *Env) AsInt64() int64 {
	if e.Value == nil {
		panic(fmt.Sprintf("Environment variable %s not found on env or default value is not set", e.Key))
	}

	if v, ok := e.Value.(int64); ok {
		return v
	}

	if v, ok := e.Value.(string); ok {
		if intValue, err := strconv.ParseInt(v, 10, 64); err == nil {
			return intValue
		}
	}

	panic(fmt.Sprintf("Environment variable %s is not an int64", e.Key))
}

func (e *Env) AsFloat64() float64 {
	if e.Value == nil {
		panic(fmt.Sprintf("Environment variable %s not found on env or default value is not set", e.Key))
	}

	if v, ok := e.Value.(float64); ok {
		return v
	}

	if v, ok := e.Value.(string); ok {
		if floatValue, err := strconv.ParseFloat(v, 64); err == nil {
			return floatValue
		}
	}

	panic(fmt.Sprintf("Environment variable %s is not a float64", e.Key))
}

func (e *Env) AsBool() bool {
	if e.Value == nil {
		panic(fmt.Sprintf("Environment variable %s not found on env or default value is not set", e.Key))
	}

	if v, ok := e.Value.(bool); ok {
		return v
	}

	if v, ok := e.Value.(string); ok {
		if boolValue, err := strconv.ParseBool(v); err == nil {
			return boolValue
		}
	}

	panic(fmt.Sprintf("Environment variable %s is not a boolean", e.Key))
}

func (e *Env) AsBytes() []byte {
	if e.Value == nil {
		panic(fmt.Sprintf("Environment variable %s not found on env or default value is not set", e.Key))
	}

	if v, ok := e.Value.(string); ok {
		return []byte(v)
	}

	panic(fmt.Sprintf("Environment variable %s cannot be converted to bytes", e.Key))
}

func (e *Env) GetEnvAsStrings(separator string) []string {
	if e.Value == nil {
		panic(fmt.Sprintf("Environment variable %s not found on env or default value is not set", e.Key))
	}

	if v, ok := e.Value.(string); ok {
		if v == "" {
			return []string{}
		}
		slc := strings.Split(v, separator)
		for i := range slc {
			slc[i] = strings.TrimSpace(slc[i])
		}
		return slc
	}

	panic(fmt.Sprintf("Environment variable %s cannot be converted to string slice", e.Key))
}
