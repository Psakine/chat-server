package config

import (
	"net"
)

const (
	grpcHostEnvName = "0.0.0.0"
	grpcPortEnvName = "50052"
)

// GRPCConfig describes config interface
type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

// NewGRPCConfig creates config struct
func NewGRPCConfig() (GRPCConfig, error) {
	return &grpcConfig{
		grpcHostEnvName,
		grpcPortEnvName,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
