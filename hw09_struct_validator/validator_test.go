package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: Response{
				Code: 404,
				Body: "Not found!",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 999,
				Body: "Someting went wrong!",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   errors.New("not included in validation set"),
				},
			},
		},
		{
			in: Token{
				Header:    nil,
				Payload:   nil,
				Signature: nil,
			},
			expectedErr: nil,
		},
		{
			in:          App{Version: "12345"},
			expectedErr: nil,
		},
		{
			in: App{Version: "123456"},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   errors.New("invalid length value"),
				},
			},
		},
		{
			in: App{Version: "1234"},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   errors.New("invalid length value"),
				},
			},
		},
		{
			in: User{
				ID:     "111111111122222222223333333333444444",
				Name:   "Name",
				Age:    49,
				Email:  "test@example.com",
				Role:   "admin",
				Phones: []string{"01234567899"},
				meta:   nil,
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "111111111122222222223333333333444444",
				Name:   "Name",
				Age:    49,
				Email:  "broken email",
				Role:   "stuff",
				Phones: []string{"01234567890"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Email",
					Err:   errors.New("regexp not matched"),
				},
			},
		},
		{
			in: User{
				ID:     "BROKEN",
				Name:   "Name",
				Age:    99,
				Email:  "broken email",
				Role:   "not in list",
				Phones: []string{"12345678901", "ai phone"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   errors.New("invalid length value"),
				},
				ValidationError{
					Field: "Age",
					Err:   errors.New("more than input value"),
				},
				ValidationError{
					Field: "Email",
					Err:   errors.New("regexp not matched"),
				},
				ValidationError{
					Field: "Role",
					Err:   errors.New("not included in validation set"),
				},
				ValidationError{
					Field: "Phones",
					Err:   errors.New("invalid length value"),
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			resError := Validate(tt.in)
			var validationErrors ValidationErrors
			if errors.As(resError, &validationErrors) {
				require.Equal(t, tt.expectedErr, validationErrors)
				return
			}
			require.Equal(t, resError, tt.expectedErr)
		})
	}
}
