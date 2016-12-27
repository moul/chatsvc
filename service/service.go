package chatsvc

import (
	"fmt"
	"io"

	nats "github.com/nats-io/go-nats"

	"github.com/moul/chatsvc/gen/pb"
)

type Service struct{}

func New() chatpb.ChatServiceServer {
	return &Service{}
}

func (s *Service) Chat(stream chatpb.ChatService_ChatServer) (err error) {
	// connect to nats server + set encoder
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	defer nc.Close()

	// forward nats messages to the stream
	c.Subscribe("general", func(m *chatpb.ChatRequest) {
		fmt.Printf("Received a message from %q: %q\n", m.SetSender, m.Message)
		stream.Send(&chatpb.ChatResponse{
			Sender:  m.SetSender,
			Message: m.Message,
			ErrMsg:  "",
		})
	})

	// forward stream messages to the nats
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		c.Publish("general", in)
	}

	return nil
}
