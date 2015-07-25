package exceptions

import (
	"fmt"
)

type Exception struct {
	Group       string
	Message     string
	Description string
}

func (err Exception) Error() string {
	return fmt.Sprintf("%s: %s %s", err.Group, err.Message, err.Description)
}
