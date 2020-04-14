package config

import (
	"os"
	"reflect"
	"strings"
)

const TagSeparator = "_"
const TagNameTOML = "toml"
const TagNameEnv = "env"

func indirect(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}
	if v.IsNil() {
		v.Set(reflect.New(v.Type().Elem()))
	}
	return indirect(reflect.Indirect(v))
}

// This function loads environment variables for config parameters with the env tag
// The environment variable's syntax for a complex struct is as given below:
// A.B.C.D is <toml-a>_<toml-b>_<toml-c>_<env-d>
// where toml-x is the toml tag for x and env-x is the env tag for x
func LoadEnvironmentVariables(value interface{}) error {
	rv := reflect.ValueOf(value)
	return parse("", "", indirect(rv))
}

// This method is called recursively in order to load environment variables for any config attribute that has a env tag
func parse(tomlTag string, envTag string, rv reflect.Value) error {
	switch rv.Kind() {

	case reflect.Ptr:
		elem := reflect.New(rv.Type().Elem())

		err := parse(tomlTag, envTag, reflect.Indirect(elem.Elem()))
		if err != nil {
			return err
		}

		rv.Set(elem)

		return nil

	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			child := rv.Field(i)

			// In case of struct fields, get the toml and env tags to be propagated
			fieldEnvTag := rv.Type().Field(i).Tag.Get(TagNameEnv)
			fieldTomlTag := rv.Type().Field(i).Tag.Get(TagNameTOML)

			// Construct the toml tag prefix to send to parse()
			// Checking for tag as empty, as we do not want to recurse into untagged structs
			if !(IsEmpty(fieldTomlTag) && IsEmpty(fieldEnvTag)) {
				structPrefix := ""
				if !IsEmpty(tomlTag) {
					structPrefix = tomlTag + TagSeparator
				}
				structPrefix = structPrefix + fieldTomlTag
				err := parse(structPrefix, fieldEnvTag, child)
				if err != nil {
					return err
				}
			}
		}

		return nil

	case reflect.String:
		// If the env tag is not empty, we need to load the value from an environment variable
		if !IsEmpty(envTag) {
			tag := strings.ToUpper(tomlTag + envTag)

			value := os.Getenv(tag)
			rv.SetString(value)
		}

		return nil
	}

	return nil
}

func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}

	reflectVal := reflect.ValueOf(val)

	switch reflectVal.Kind() {
	case reflect.Int:
		return val.(int) == 0

	case reflect.Int64:
		return val.(int64) == 0

	case reflect.String:
		return strings.TrimSpace(val.(string)) == ""

	case reflect.Map:
		fallthrough
	case reflect.Slice:
		return reflectVal.IsNil() || reflectVal.Len() == 0
	}

	return false
}
