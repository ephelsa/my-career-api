package data

import (
	"fmt"
)

var (
	InvalidAuthError = func(email string) error {
		return fmt.Errorf("%s, email or password are invalid", email)
	}
	UserRegisteredError = func(email string) error {
		return fmt.Errorf("%s exists", email)
	}
	UserRegisteredWithoutConfirmError = func(email string) error {
		return fmt.Errorf("%s must confirm the registry", email)
	}
	UserNotRegisteredError = func(email string) error {
		return fmt.Errorf("%s doesn't exist", email)
	}
	PasswordWithoutMinLenError = func(minLen int) error {
		return fmt.Errorf("password must contain %d characters min", minLen)
	}
)
