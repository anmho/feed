package server

import (
	"context"
	"fmt"
	"log"

	greetv1 "feed/gen/protos/greet/v1"
	"feed/gen/protos/greet/v1/greetv1connect"

	"connectrpc.com/connect"
)

var _ greetv1connect.GreetServiceHandler = (*GreetService)(nil)

type GreetService struct {
}

func NewGreetService() *GreetService {
	greeter := &GreetService{}
	return greeter
}

func (s *GreetService) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
