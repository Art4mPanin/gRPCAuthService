package auth

import (
	"context"
	"github.com/Art4mPanin/gRPCAuthService/internal/data/gen/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *ServerAuth) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	// where's the fucking error here ????
	accessToken, refreshToken, user2, err := s.Services.Auth.Login(ctx, username, password)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to login user")
	}
	if user2 == nil {
		s.log.Error("user is nil after login attempt", slog.String("username", username))
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	//TODO: ERROR HANDLE
	return &auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: &auth.User{
			Id:             int32(user2.ID),
			Username:       user2.Username, // user2 potentially could be nil, need error handler
			Email:          user2.Email,
			IsSuperuser:    user2.IsSuperuser,
			HashedPassword: user2.HashedPassword,
		},
	}, nil
}
func (s *ServerAuth) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	email := req.GetEmail()
	accessToken, refreshToken, user2, err := s.Services.Auth.Register(ctx, username, password, email)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to register user")
	}
	//TODO: ERROR HANDLING
	if user2 == nil || accessToken == "" || refreshToken == "" {
		s.log.Error("user is nil or tokens are empty after register attempt", slog.String("username", username))
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	return &auth.RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: &auth.User{
			Id:             int32(user2.ID),
			Username:       user2.Username,
			Email:          user2.Email,
			IsSuperuser:    user2.IsSuperuser,
			HashedPassword: user2.HashedPassword,
		},
	}, nil
}
