package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
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
		cli.StringFlag{
			Name:   "sender",
			Value:  "",
			EnvVar: "SENDER",
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

	var errc = make(chan error)

	sender := c.String("sender")
	if sender == "" {
		sender = fmt.Sprintf("pid%d", os.Getpid())
	}

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				errc <- nil
				return
			}
			if err != nil {
				errc <- err
				return
			}
			if message.Sender != sender {
				fmt.Printf("%s> %s\n", message.Sender, message.Message)
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "/quit" {
				errc <- nil
				return
			}
			stream.Send(&chatpb.ChatRequest{
				SetSender: sender,
				Message:   line,
			})
		}
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("quit: %v\n", <-errc)
	return nil
}
