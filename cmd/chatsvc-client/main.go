package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/moul/chatsvc/gen/pb"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "grpc-server",
			Value:  "127.0.0.1:9000",
			EnvVar: "GRPC_SERVER",
		},
	}
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	conn, err := grpc.Dial(
		c.String("grpc-server"),
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx := context.Background()
	client := chatpb.NewChatServiceClient(conn)

	stream, err := client.Chat(ctx)
	if err != nil {
		return err
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}

	return nil
}
