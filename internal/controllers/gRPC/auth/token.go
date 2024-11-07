package auth

import (
	"context"
	"github.com/Art4mPanin/gRPCAuthService/internal/data/gen/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAuth) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	oldrefreshToken := req.RefreshToken
	newAccessToken, newRefreshToken, err := s.Services.Token.RefreshToken(oldrefreshToken)
	if err != nil || newAccessToken == "" || newRefreshToken == "" {
		return nil, status.Errorf(codes.Unauthenticated, "refresh token validation failed: %v", err)
	}
	return &auth.RefreshResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
func (s *ServerAuth) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	authjwtHeader := req.Auth_JWT_Header
	err := s.Services.Token.Validate(authjwtHeader)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token validation failed: %v", err)
	}
	return &auth.ValidateResponse{Valid: true}, nil
}
