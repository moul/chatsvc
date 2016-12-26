package chat_clientgrpc

import (
	jwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	endpoints "github.com/moul/chatsvc/gen/endpoints"
	pb "github.com/moul/chatsvc/gen/pb"
)

func New(conn *grpc.ClientConn, logger log.Logger) pb.ChatServiceServer {

	return &endpoints.Endpoints{}
}
