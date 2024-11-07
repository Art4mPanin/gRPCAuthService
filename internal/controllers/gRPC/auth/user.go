package auth

import (
	"context"
	"github.com/Art4mPanin/gRPCAuthService/internal/data/gen/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *ServerAuth) GetMe(ctx context.Context, req *auth.GetMeRequest) (*auth.GetMeResponse, error) {
	authJwtHeader := req.Auth_JWT_Header
	user, err := s.Services.User.GetMe(authJwtHeader)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token validation failed: %v", err)
	}
	if user == nil {
		s.log.Error("userinfo is nil in getting userinfo", slog.String("jwt_auth_header", authJwtHeader))
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	return &auth.GetMeResponse{
		User: &auth.User{
			Id:             int32(user.ID),
			Username:       user.Username,
			Email:          user.Email,
			IsSuperuser:    user.IsSuperuser,
			HashedPassword: user.HashedPassword,
		},
	}, nil
}
