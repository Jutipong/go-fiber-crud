package utils

import (
	"encoding/json"
	"fiber-crud/pkg/enum"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

//## ตัวอย่างการใช้งาน => util.JsonSerialize(payload)
func JsonSerialize(payload interface{}) string {
	b, err := json.Marshal(&payload)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

//## ตัวอย่างการใช้งาน => util.JsonDeserialize(str, &result)
func JsonDeserialize(str string, st interface{}) {
	json.Unmarshal([]byte(str), &st)
}

func ConvertFloat64ToJsonNumber(f float64) json.Number {
	s := fmt.Sprintf("%.2f", f)
	return json.Number(s)
}

func Contains(slice *[]string, item string) bool {
	set := make(map[string]struct{}, len(*slice))
	for _, s := range *slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func StringIsEmpty(input string) bool {
	return utf8.RuneCountInString(strings.TrimSpace(input)) == 0
}

func StringEmpty_SetDefault(input string, valueDefault string) string {
	if utf8.RuneCountInString(strings.TrimSpace(input)) == 0 {
		return valueDefault
	} else {
		return input
	}
}

func GetFiberErr(err error) (int, string) {
	if err != nil {
		e := err.(*fiber.Error)
		return e.Code, e.Message
	} else {
		return enum.Ok, ""
	}
}

func GetGookitError(v *validate.Validation) (errs []string) {
	for fieldName := range v.Errors.All() {
		errs = append(errs, fieldName)
	}
	return errs
}
