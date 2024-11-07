package password

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(dbpassword, inpassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(inpassword))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

func GeneratePassword(ctx context.Context, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}
