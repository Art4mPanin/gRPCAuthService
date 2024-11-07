package gRPC

import (
	"fmt"
	authrpc "github.com/Art4mPanin/gRPCAuthService/internal/controllers/gRPC/auth"
	userrepository "github.com/Art4mPanin/gRPCAuthService/internal/repositories/user"
	authservice "github.com/Art4mPanin/gRPCAuthService/internal/services/AuthService"
	"github.com/Art4mPanin/gRPCAuthService/internal/services/TokenService"
	"github.com/Art4mPanin/gRPCAuthService/internal/services/UserService"
	storage "github.com/Art4mPanin/gRPCAuthService/internal/storage"
	"github.com/sagikazarmark/slog-shim"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

type GRPC struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGRPC(log *slog.Logger, port int) *GRPC {
	gRPCServer := grpc.NewServer()
	db := storage.InitDB()
	registerAuthHandler(gRPCServer, db, log) // todo: db coonection

	return &GRPC{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func registerAuthHandler(server *grpc.Server, DB *gorm.DB, logger *slog.Logger) {

	repo := userrepository.NewUserRepository(DB)
	authService := authservice.NewAuthService(repo, logger)
	tokenService := TokenService.NewTokenService(repo, logger)
	userService := UserService.NewUserService(repo, logger)
	authrpc.Register(server, authService, userService, tokenService)
	// todo: other services (user, token)
	//	todo: pornhub.Register(server, pronhubService)
}

func (g *GRPC) Run() error {
	s, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	if err = g.gRPCServer.Serve(s); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	g.log.Info("gRPC server is running on port: %d", g.port)
	return nil
}
func (g *GRPC) Close() {
	g.gRPCServer.GracefulStop()
}
