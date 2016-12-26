package chat_grpctransport

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"

	endpoints "github.com/moul/chatsvc/gen/endpoints"
	pb "github.com/moul/chatsvc/gen/pb"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoints.Endpoints) pb.ChatServiceServer {
	var options []grpctransport.ServerOption
	_ = options
	return &grpcServer{

		chat: &server{
			e: endpoints.ChatEndpoint,
		},
	}
}

type grpcServer struct {
	chat streamHandler
}

func (s *grpcServer) Chat(server pb.ChatService_ChatServer) error {
	return s.chat.Do(server, nil)
}

func decodeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

type streamHandler interface {
	Do(server interface{}, req interface{}) (err error)
}

type server struct {
	e endpoints.StreamEndpoint
}

func (s server) Do(server interface{}, req interface{}) (err error) {
	if err := s.e(server, req); err != nil {
		return err
	}
	return nil
}
