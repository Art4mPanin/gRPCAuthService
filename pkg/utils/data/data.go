package data

import "fmt"

func ValidateLoginRequest(username, password string) error {
	if len(username) < 6 {
		return fmt.Errorf("username must be at least 6 characters long")
	}
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	return nil
}
