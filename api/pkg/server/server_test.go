package server

import (
	"context"
	"testing"

	greetv1 "feed/gen/protos/greet/v1"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
)

func TestServer_Greet(t *testing.T) {
	tests := []struct {
		desc             string
		request          *connect.Request[greetv1.GreetRequest]
		expectedResponse *greetv1.GreetResponse
	}{
		{
			desc: "happy path: greets caller",
			request: &connect.Request[greetv1.GreetRequest]{
				Msg: &greetv1.GreetRequest{
					Name: "Joe",
				},
			},
			expectedResponse: &greetv1.GreetResponse{
				Greeting: "Hello, Joe!",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			greetService := NewGreetService()

			resp, err := greetService.Greet(context.Background(), tc.request)

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tc.expectedResponse, resp.Msg)
		})
	}
}
