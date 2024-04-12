package env

import (
	"fmt"
	"testing"
)

func TestFillStructFromEnv(t *testing.T) {
	stringEnv := "test string env"
	intEnv := 1
	boolEnv := true

	t.Setenv("STRING_ENV", stringEnv)
	t.Setenv("INT_ENV", fmt.Sprintf("%d", intEnv))
	t.Setenv("BOOL_ENV", fmt.Sprintf("%v", boolEnv))

	type TestStructSuccess struct {
		StringField string `env:"STRING_ENV"`
		IntField    int    `env:"INT_ENV"`
		BoolField   bool   `env:"BOOL_ENV"`
	}

	var testStructSuccess TestStructSuccess

	err := FillStructFromEnv(&testStructSuccess)
	if err != nil {
		t.Fatalf("FillStructFromEnv error = %s", err.Error())
	}

	if testStructSuccess.StringField != stringEnv {
		t.Fatalf("test string env fail, StringField = %v", testStructSuccess.StringField)
	}

	if testStructSuccess.IntField != intEnv {
		t.Fatalf("test string env fail, IntField = %v", testStructSuccess.StringField)
	}

	if testStructSuccess.BoolField != boolEnv {
		t.Fatalf("test string env fail, BoolField = %v", testStructSuccess.StringField)
	}

	type TestStructFail struct {
		StringField    string  `env:"STRING_ENV"`
		IntField       int     `env:"INT_ENV"`
		BoolField      bool    `env:"BOOL_ENV"`
		IncorrectField float64 `env:"INCORRECT_ENV"`
	}

	var testStructFail TestStructFail

	err = FillStructFromEnv(&testStructFail)
	if err == nil {
		t.Fatalf("FillStructFromEnv without error on testStructFail")
	}
}
