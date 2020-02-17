package searchuser

import (
    "github.com/sirupsen/logrus"
    service "github.com/tapokshot/microslab/internal/app/grpc"
    "github.com/tapokshot/microslab/internal/app/searchuser/controller"
    "google.golang.org/grpc"
    "net"
)

type SearchUserServer struct {
    config *Config
    logger *logrus.Logger
}

func New(config *Config) *SearchUserServer {
    return &SearchUserServer{
        config: config,
        logger: logrus.New(),
    }
}

func (server *SearchUserServer) Launch() error {
    if err := server.configureLogger(); err != nil {
        return err
    }
    listener, err := net.Listen(server.config.Network, server.config.Port)
    if err != nil {
        return err
    }
    grpcServer := grpc.NewServer()
    service.RegisterSearchUserServer(grpcServer, &controller.SearchUserController{})
    if err := grpcServer.Serve(listener); err != nil {
        return err
    }
    return nil
}

func (server *SearchUserServer) configureLogger() error {
    level, err := logrus.ParseLevel(server.config.LogLevel)
    if err != nil {
        return err
    }
    server.logger.SetLevel(level)
    return nil
}
