package server

import (
	"context"
	feedv1 "feed/gen/protos/feed/v1"
	"feed/pkg/test"
	"feed/pkg/users"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeedService_CreatePost(t *testing.T) {
	postID1 := uuid.New()
	user1 := &feedv1.User{
		UserId: uuid.NewString(),
		Name:   "Joe",
	}

	tests := []struct {
		desc    string
		author *feedv1.User
		request *connect.Request[feedv1.CreatePostRequest]

		expectedResponse *connect.Response[feedv1.CreatePostResponse]
	}{
		{
			desc: "happy path: post create successfully",
			author: user1,
			request: &connect.Request[feedv1.CreatePostRequest]{
				Msg: &feedv1.CreatePostRequest{
					Id:      postID1.String(),
					Content: "hello world",
				},
			},

			expectedResponse: &connect.Response[feedv1.CreatePostResponse]{
				Msg: &feedv1.CreatePostResponse{},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			tc.request.Header().Add("X-User", tc.author.UserId)

			db := test.MakeLocalDB(t)
			err := users.CreateUser(db, tc.author)
			require.NoError(t, err)
			feedService := MakeFeedService(db)

			ctx := context.Background()
			response, err := feedService.CreatePost(ctx, tc.request)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResponse.Msg, response.Msg, "equal response")
		})
	}
}

func TestFeedService_SyncFeed(t *testing.T) {
	// db := test.MakeLocalDB(t)

	// s := MakeFeedService(db)

}
