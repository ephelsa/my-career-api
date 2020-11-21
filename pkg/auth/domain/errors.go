package domain

import (
	fmt "fmt"
)

var (
	InvalidAuth = func(email string) error {
		return fmt.Errorf("%s invalid password or email", email)
	}
	UnRegisteredUser = func(email string) error {
		return fmt.Errorf("%s exists", email)
	}
	PasswordWithoutMinLen = func(minLen int) error {
		return fmt.Errorf("password must contain %d characters min", minLen)
	}
)
