package env

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// FillStructFromEnv Fill struct fields from enviroment variables, using tag group "env". Work only with string, bool and int
func FillStructFromEnv(pointerToStruct any) error {
	structValue := reflect.ValueOf(pointerToStruct)
	if structValue.Kind() != reflect.Ptr || structValue.Elem().Kind() != reflect.Struct {
		return errors.New("passed value is not a pointer to a structure")
	}

	structType := structValue.Elem().Type()
	for i := 0; i < structType.NumField(); i++ {
		field := structValue.Elem().Field(i)
		if !field.IsValid() || !field.CanSet() {
			return fmt.Errorf("incorrect field %sm IsValid %v, CanSet %v", structType.Field(i).Name, field.IsValid(), field.CanSet())
		}
		switch field.Kind() {
		case reflect.Int:
			value, _ := strconv.ParseInt(os.Getenv(getTagValue(structType.Field(i), "env")), 10, 64)
			field.SetInt(value)
		case reflect.String:
			field.SetString(os.Getenv(getTagValue(structType.Field(i), "env")))
		case reflect.Bool:
			value, _ := strconv.ParseBool(os.Getenv(getTagValue(structType.Field(i), "env")))
			field.SetBool(value)
		default:
			return fmt.Errorf("incorrect field kind %s is %v", structType.Field(i).Name, field.Kind())
		}
	}

	return nil
}

func FillStructFromEnvFatal(pointerToStruct any) {
	err := FillStructFromEnv(pointerToStruct)
	if err != nil {
		log.Fatal(err)
	}
}

func getTagValue(field reflect.StructField, tagType string) string {
	tag := field.Tag.Get(tagType)
	if tag == "" {
		return ""
	}

	tagValue := strings.Split(tag, ",")[0]
	if tagValue == "" {
		log.Fatalf("incorrect tag format %s for field %s", tagType, field.Name)
	}

	return tagValue
}
