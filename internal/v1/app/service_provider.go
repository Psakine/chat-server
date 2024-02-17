package app

import (
	"log"

	"github.com/Psakine/chat-server/internal/v1/api/chat"
	"github.com/Psakine/chat-server/internal/v1/config"
)

// ServiceProvider ...
type ServiceProvider struct {
	grpcConfig config.GRPCConfig
	chatServer *chat.Server
}

// NewServiceProvider ...
func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

// GRPCConfig ...
func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// ChatServer ...
func (s *ServiceProvider) ChatServer() *chat.Server {
	if s.chatServer == nil {
		s.chatServer = chat.NewChatServer()
	}

	return s.chatServer
}
