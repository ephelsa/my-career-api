package data

import (
	"fmt"
)

var (
	InvalidAuth = func(email string) error {
		return fmt.Errorf("%s invalid password or email", email)
	}
	UserRegistered = func(email string) error {
		return fmt.Errorf("%s exists", email)
	}
	UserRegisteredWithoutConfirm = func(email string) error {
		return fmt.Errorf("%s must confirm the registry", email)
	}
	PasswordWithoutMinLen = func(minLen int) error {
		return fmt.Errorf("password must contain %d characters min", minLen)
	}
)
