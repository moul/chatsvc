package chat_endpoints

import (
	"fmt"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/moul/chatsvc/gen/pb"
	context "golang.org/x/net/context"
)

var _ = endpoint.Chain
var _ = fmt.Errorf
var _ = context.Background

type StreamEndpoint func(server interface{}, req interface{}) (err error)

type Endpoints struct {
	ChatEndpoint StreamEndpoint
}

func (e *Endpoints) Chat(server pb.ChatService_ChatServer) error {
	return fmt.Errorf("not implemented")
}

func MakeChatEndpoint(svc pb.ChatServiceServer) StreamEndpoint {
	return func(server interface{}, request interface{}) error {

		return svc.Chat(server.(pb.ChatService_ChatServer))

	}
}

func MakeEndpoints(svc pb.ChatServiceServer) Endpoints {
	return Endpoints{

		ChatEndpoint: MakeChatEndpoint(svc),
	}
}
