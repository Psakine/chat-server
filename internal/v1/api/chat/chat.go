package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Psakine/chat-server/pkg/v1/chat"
	"github.com/brianvoe/gofakeit/v6"
)

// Server ...
type Server struct {
	chat.UnimplementedChatV1Server
}

// NewChatServer ...
func NewChatServer() *Server {
	return &Server{}
}

// Create ...
func (s Server) Create(_ context.Context, _ *chat.CreateRequest) (*chat.CreateResponse, error) {
	return &chat.CreateResponse{Id: gofakeit.Int64()}, nil
}

// Delete ...
func (s Server) Delete(_ context.Context, _ *chat.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// SendMessage ...
func (s Server) SendMessage(_ context.Context, _ *chat.SendMessageRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
