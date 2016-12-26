package chat_httptransport

import (
	"encoding/json"
	context "golang.org/x/net/context"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "github.com/moul/chatsvc/gen/endpoints"
	pb "github.com/moul/chatsvc/gen/pb"
)

var _ = log.Printf
var _ = gokit_endpoint.Chain
var _ = httptransport.NewClient

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(ctx context.Context, svc pb.ChatServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {

	return nil
}
