package TokenService

import (
	"github.com/Art4mPanin/gRPCAuthService/internal/models"
	"github.com/Art4mPanin/gRPCAuthService/pkg/utils/jwt"
	"log/slog"
	"strings"
)

type tokenUserRepository interface {
	FindUserInDBByID(id int) (*models.User, error)
}

type TokenService struct {
	userRepo tokenUserRepository
	log      *slog.Logger
}

func NewTokenService(repository tokenUserRepository, log *slog.Logger) *TokenService {
	return &TokenService{
		userRepo: repository,
		log:      log,
	}
}

func (s *TokenService) Validate(token string) error {
	s.log.Info("Attempting to validate token")
	s.log.Info("Token: " + token)
	token = strings.TrimPrefix(token, "Bearer ")
	jwttoken, err := jwt.GetToken(token)
	if err != nil {
		s.log.Warn("Error getting and parsing token: %v", err)
		return err
	}
	s.log.Info("Got and parsed the token")
	userid, err := jwt.ValidateToken(jwttoken)
	if err != nil {
		s.log.Warn("Error validating the token: %v", err)
		return err
	}
	s.log.Info("Validated the token")
	_, err = s.userRepo.FindUserInDBByID(userid)
	if err != nil {
		s.log.Warn("User with the following token not found in the database")
		return err
	}
	s.log.Info("User with the following token found in the database")
	return nil
}
func (s *TokenService) RefreshToken(token string) (string, string, error) {
	s.log.Info("Attempting to refresh token")
	jwtToken, err := jwt.GetToken(token)
	if err != nil {
		s.log.Warn("Error getting and parsing token: %v", err)
		return "", "", err
	}
	s.log.Info("Got and parsed the token")
	userid, err := jwt.ValidateToken(jwtToken)
	if err != nil {
		s.log.Warn("Error validating the token: %v", err)
		return "", "", err
	}
	s.log.Info("Validated the token")
	user, err := s.userRepo.FindUserInDBByID(userid)
	if err != nil {
		s.log.Warn("User with the following token not found in the database")
		return "", "", err
	}
	s.log.Info("User with the following token found in the database")
	accessToken, refreshToken, err := jwt.CreateToken(userid, user.IsSuperuser)
	if err != nil {
		s.log.Warn("Failed to create tokens: %v", err)
		return "", "", err
	}
	s.log.Info("Created new tokens")
	return accessToken, refreshToken, nil
}
