package chatsvc

import (
	"fmt"

	"github.com/moul/chatsvc/gen/pb"
)

type Service struct{}

func New() chatpb.ChatServiceServer {
	return &Service{}
}

func (s *Service) Chat(server chatpb.ChatService_ChatServer) (err error) {
	return fmt.Errorf("not implemented")
}
