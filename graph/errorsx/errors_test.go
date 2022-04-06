package errorsx

import (
	"fmt"
	"testing"
)

type inputError struct {
	message      string
	missingField string
}

func (i inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField
}

func validate(name, gender string) error {
	if name == "" {
		return &inputError{message: "Name is mandatory", missingField: "name"}
	}
	if gender == "" {
		return &inputError{message: "Gender is mandatory", missingField: "gender"}
	}
	return nil
}

func TestErrors(t *testing.T) {
	err := validate("", "")
	if err != nil {
		if err, ok := err.(*inputError); ok {
			_, _ = fmt.Println(err)
			_, _ = fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
	}
}
