// Package hw09structvalidator -- HW09 otus.
package hw09structvalidator

import (
	"errors"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// ValidationError structure.
type ValidationError struct {
	Field string
	Err   error
}

// ValidationErrors slice of structs.
type ValidationErrors []ValidationError

// Error handler.
func (v ValidationErrors) Error() string {
	var builder strings.Builder
	for _, e := range v {
		builder.WriteString("field: " + e.Field + " - " + e.Err.Error())
	}
	return builder.String()
}

// Validate public structure.
func Validate(v interface{}) error {
	var vErrors ValidationErrors
	s := reflect.ValueOf(v)

	if s.Kind() != reflect.Struct {
		return ValidationErrors{
			ValidationError{
				Field: "",
				Err:   errors.New("not a structure"),
			},
		}
	}

	t := s.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			log.Println("no one tag for key 'validate'")
			continue
		}
		fv := s.Field(i)

		if !fv.CanInterface() {
			continue
		}

		errs := validateValue(tag, field.Name, fv)
		if len(errs) > 0 {
			vErrors = append(vErrors, errs...)
		}
	}

	if len(vErrors) > 0 {
		return vErrors
	}
	return nil
}

func validateValue(tag, field string, v reflect.Value) ValidationErrors {
	//nolint:exhaustive
	switch v.Kind() {
	case reflect.String:
		var errs ValidationErrors
		validators := stringValidators(tag, field, v.String())
		for _, validator := range validators {
			err := validator.validate()
			if err != nil {
				log.Println(err.Err.Error())
				errs = append(errs, *err)
			}
		}
		return errs
	case reflect.Int:
		var errs ValidationErrors
		validators := intValidators(tag, field, int(v.Int()))
		for _, validator := range validators {
			err := validator.Validate()
			if err != nil {
				errs = append(errs, *err)
			}
		}
		return errs
	case reflect.Slice:
		var errs ValidationErrors
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			errs = append(errs, validateValue(tag, field, elem)...)
		}
		return errs
	default:
		return nil
	}
}

func stringValidators(tag, field, value string) []stringValidator {
	validatorsRaw := strings.Split(tag, "|")
	validators := make([]stringValidator, 0)

	for _, valRaw := range validatorsRaw {
		val := strings.Split(valRaw, ":")
		if len(val) != 2 {
			continue
		}
		sVal := stringValidator{
			name:      val[0],
			condition: val[1],
			field:     field,
			value:     value,
		}
		validators = append(validators, sVal)
	}
	return validators
}

type stringValidator struct {
	name      string
	condition string
	field     string
	value     string
}

func (t stringValidator) validate() *ValidationError {
	switch t.name {
	case "len":
		cond, err := strconv.Atoi(t.condition)
		if err != nil {
			return nil
		}
		l := len(t.value)
		if l != cond {
			return &ValidationError{
				Field: t.field,
				Err:   errors.New("invalid length value"),
			}
		}
	case "regexp":
		err := t.regExp()
		if err != nil {
			return err
		}
	case "in":
		set := strings.Split(t.condition, ",")
		for _, e := range set {
			if t.value == e {
				return nil
			}
		}
		return &ValidationError{
			Field: t.field,
			Err:   errors.New("not included in validation set"),
		}
	}
	return nil
}

func (t stringValidator) regExp() *ValidationError {
	matched, err := regexp.MatchString(t.condition, t.value)
	if err != nil {
		return &ValidationError{Field: t.field, Err: err}
	}
	if !matched {
		return &ValidationError{
			Field: t.field,
			Err:   errors.New("regexp not matched"),
		}
	}
	return nil
}

func intValidators(tag, field string, value int) []intValidator {
	validatorsRaw := strings.Split(tag, "|")
	validators := make([]intValidator, 0)

	for _, valRaw := range validatorsRaw {
		val := strings.Split(valRaw, ":")
		if len(val) != 2 {
			continue
		}
		sVal := intValidator{
			name:      val[0],
			condition: val[1],
			field:     field,
			value:     value,
		}
		validators = append(validators, sVal)
	}
	return validators
}

type intValidator struct {
	name      string
	condition string
	field     string
	value     int
}

func (t intValidator) Validate() *ValidationError {
	switch t.name {
	case "min":
		cond, err := strconv.Atoi(t.condition)
		if err != nil {
			return nil
		}
		if t.value < cond {
			return &ValidationError{
				Field: t.field,
				Err:   errors.New("less than input value"),
			}
		}
	case "max":
		cond, err := strconv.Atoi(t.condition)
		if err != nil {
			return nil
		}
		if t.value > cond {
			return &ValidationError{
				Field: t.field,
				Err:   errors.New("more than input value"),
			}
		}
	case "in":
		err := t.in()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t intValidator) in() *ValidationError {
	set := strings.Split(t.condition, ",")
	for _, val := range set {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return nil
		}
		if t.value == intVal {
			return nil
		}
	}
	return &ValidationError{
		Field: t.field,
		Err:   errors.New("not included in validation set"),
	}
}
