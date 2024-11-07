package auth

import (
	"context"
	"github.com/Art4mPanin/gRPCAuthService/internal/data/gen/auth"
	"github.com/Art4mPanin/gRPCAuthService/internal/models"
	"google.golang.org/grpc"
	"log/slog"
)

type grpcAuthService interface {
	Login(ctx context.Context, username, password string) (string, string, *models.User, error)
	Register(ctx context.Context, username, password, email string) (string, string, *models.User, error)
}
type grpcUserService interface {
	GetMe(authJwtHeader string) (*models.User, error)
}
type grpcTokenService interface {
	Validate(token string) error
	RefreshToken(token string) (atoken, rtoken string, err error)
}
type services struct {
	Auth  grpcAuthService
	User  grpcUserService
	Token grpcTokenService
}

type ServerAuth struct {
	auth.UnimplementedAuthServer

	Services services
	log      *slog.Logger
}

// controller can be: struct with method / func login(ctx echo.context) err {}
// http: controller -> service -> repository -> db
// http adaper way: controller -> adapter -> userService -> userRepo
//                                        |-> tokenService -> tokenRepo
// http service injection way: controller -> authservice -> userRepo
//                                             |-> tokenService
// here we have injected token service as dependency in authservice without using a adaper

// grpc: no controller, but server method. server is ServerAuth here, that should implement auth.UnimplementedAuthServer

// grpc: server (listeing connections from net.listen()) -> struct with functions -> func (example: login) -> service -> repository -> db

// Regiser: add struct with functions to server

// rpc: remote protocoll caller
// rpc is a remote call function with result

func Register(gRPC *grpc.Server, authService grpcAuthService, userService grpcUserService, tokenService grpcTokenService) {
	// grpc is a sever
	// ServerAuth is a struct with functions

	auth.RegisterAuthServer(gRPC, &ServerAuth{Services: services{
		Auth:  authService,
		User:  userService,
		Token: tokenService,
	}})
}

//TODO: LOGGER
