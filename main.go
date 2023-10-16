package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type Users struct {
	Name     string `required:"true" maxLen:"10"`
	Age      int    `required:"true" min:"18" max:"60"`
	Username string `required:"true"`
	Password string `required:"true" minLen:"6"`
	Level    string `required:"true"`
	Email    string `required:"true"`
}

// ValidateStruct adalah fungsi untuk memvalidasi sebuah struct berdasarkan tag-nya
func ValidateStruct(s interface{}) error {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Memeriksa apakah field harus ada
		if field.Tag.Get("required") == "true" {
			// Memeriksa apakah itu string atau integer
			if value.Kind() == reflect.String {
				if value.String() == "" {
					return fmt.Errorf("%s is required", field.Name)
				}
			} else if value.Kind() == reflect.Int {
				if value.Int() == 0 {
					return fmt.Errorf("%s is required", field.Name)
				}
			}
		}

		// Memeriksa constraint panjang maksimum
		if maxLen, ok := field.Tag.Lookup("maxLen"); ok {
			if len(value.String()) > atoi(maxLen) {
				return fmt.Errorf("%s exceeds maximum length of %s", field.Name, maxLen)
			}
		}

		// Memeriksa constraint panjang minimum
		if minLen, ok := field.Tag.Lookup("minLen"); ok {
			if len(value.String()) < atoi(minLen) {
				return fmt.Errorf("%s is shorter than minimum length of %s", field.Name, minLen)
			}
		}

		// Memeriksa constraint nilai minimum (untuk integer)
		if min, ok := field.Tag.Lookup("min"); ok {
			if intVal, err := strconv.Atoi(min); err == nil {
				if value.Int() < int64(intVal) {
					return fmt.Errorf("%s is less than minimum value of %s", field.Name, min)
				}
			}
		}

		// Memeriksa constraint nilai maksimum (untuk integer)
		if max, ok := field.Tag.Lookup("max"); ok {
			if intVal, err := strconv.Atoi(max); err == nil {
				if value.Int() > int64(intVal) {
					return fmt.Errorf("%s is greater than maximum value of %s", field.Name, max)
				}
			}
		}

		// Memeriksa format email
		if field.Name == "Email" {
			emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
			if !emailRegex.MatchString(value.String()) {
				return fmt.Errorf("%s is not a valid email address", field.Name)
			}
		}
	}

	return nil
}

// atoi adalah fungsi bantu untuk mengonversi string menjadi integer
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	newUser := Users{
		Name:     "Rashford",
		Age:      25,
		Username: "marcusrashford",
		Password: "123456",
		Level:    "admin",
		Email:    "rashford@email.com",
	}

	err := ValidateStruct(newUser)
	fmt.Println(err)
}
