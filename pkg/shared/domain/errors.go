package domain

import (
	"fmt"
)

var (
	ResourceNotFound = func(resource string) error {
		return fmt.Errorf("'%s' resource not found", resource)
	}
	ResourcesEmpty   = fmt.Errorf("resource is empty")
	ResourcesInvalid = fmt.Errorf("resource is invalid")
)
