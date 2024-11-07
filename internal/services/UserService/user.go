package UserService

import (
	"github.com/Art4mPanin/gRPCAuthService/internal/models"
	"github.com/Art4mPanin/gRPCAuthService/pkg/utils/jwt"
	"log/slog"
)

type userUserRepository interface {
	FindUserInDBByID(id int) (*models.User, error)
}

type UserService struct {
	//Db       *gorm.DB
	userRepo userUserRepository
	log      *slog.Logger
}

func NewUserService(repository userUserRepository, log *slog.Logger) *UserService {
	return &UserService{
		userRepo: repository,
		log:      log,
	}

}

// --- user.go
func (u *UserService) GetMe(authJwtHeader string) (*models.User, error) {
	u.log.Info("Getting user by JWT: %s", authJwtHeader)
	token, err := jwt.GetToken(authJwtHeader)
	if err != nil {
		u.log.Error("Error getting and parsing JWT token: %s", err)
		return nil, err
	}
	u.log.Info("Validating JWT token: %s", token)
	userid, err := jwt.ValidateToken(token)
	if err != nil {
		u.log.Error("Error validating JWT token: %s", err)
		return nil, err
	}
	u.log.Info("User ID from JWT token: %d", userid)
	userId := userid
	user, err := u.userRepo.FindUserInDBByID(userId)
	if err != nil {
		u.log.Error("Error finding user in DB by ID: %s", err)
		return nil, err
	}
	u.log.Info("User found by ID: %+v", user)
	return &models.User{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		IsSuperuser:    user.IsSuperuser,
		HashedPassword: user.HashedPassword,
	}, nil
}

// ---
