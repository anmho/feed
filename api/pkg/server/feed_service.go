package server

import (
	"context"
	"database/sql"
	"errors"
	feedv1 "feed/gen/protos/feed/v1"
	"feed/gen/protos/feed/v1/feedv1connect"
	"log/slog"
	"time"

	"connectrpc.com/connect"
)

var _ feedv1connect.FeedServiceHandler = (*FeedService)(nil)

// scoped logger should have info about the request populated (who, what, when, why, where etc)
var logger = slog.Default()

type FeedService struct {
	db   *sql.DB
	opts options
}

type FanoutMode int

const (
	WriteFanout FanoutMode = iota
	ReadFanout
)

type options struct {
	fanoutMode FanoutMode
}

type Option func() func(o *options)
type optFunc func(o *options)

func WithWriteFanout() optFunc {
	return func(o *options) {
		o.fanoutMode = WriteFanout
	}
}
func WithReadFanout() optFunc {
	return func(o *options) {
		o.fanoutMode = WriteFanout
	}
}

func MakeFeedService(db *sql.DB, opts ...optFunc) *FeedService {
	var o options

	for _, opt := range opts {
		opt(&o)
	}

	return &FeedService{
		db:   db,
		opts: o,
	}
}

func (s *FeedService) SyncFeed(
	ctx context.Context,
	req *connect.Request[feedv1.SyncFeedRequest],
) (*connect.Response[feedv1.SyncFeedResponse], error) {
	userID := req.Header().Get("X-User")

	switch s.opts.fanoutMode {
	case WriteFanout:
		// Read the records in the feed table for this user
		posts := make([]*feedv1.Post, 0)

		query := `
		SELECT post_id, content
		FROM feeds f JOIN posts p ON f.post_id = p.post_id
		WHERE user_id = $1
		;
		`
		// we need to index user_id and post_id
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		rows, err := s.db.QueryContext(ctx, query, userID)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var p feedv1.Post
			rows.Scan(&p.Id, &p.Content)
			posts = append(posts, &p)
		}

		resp := &connect.Response[feedv1.SyncFeedResponse]{
			Msg: &feedv1.SyncFeedResponse{
				Items: posts,
			},
		}
		return resp, nil
	case ReadFanout:
		// Compute which users this user is a follower of and gather the posts for each of those users.
	default:
		logger.Error("unknown fanout option", "fanoutMode", s.opts.fanoutMode)
	}

	return nil, nil
}

func (s *FeedService) CreatePost(
	ctx context.Context,
	req *connect.Request[feedv1.CreatePostRequest],
) (*connect.Response[feedv1.CreatePostResponse], error) {
	if req.Msg == nil {
		return nil, errors.New("nil msg")
	}

	var resp *connect.Response[feedv1.CreatePostResponse]

	userID := req.Header().Get("X-User")



	// what happens when a cancel occurs?
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	query := `
	INSERT INTO posts (
		id, content, author_id
	) VALUES (
		$1, $2, $3
	)
	;
	`

	_, err := s.db.ExecContext(ctx, query, 
		req.Msg.Id, 
		req.Msg.Content,
		userID,
	)
	if err != nil {
		return nil, err
	}
	resp = &connect.Response[feedv1.CreatePostResponse]{
		Msg: &feedv1.CreatePostResponse{},
	}
	switch s.opts.fanoutMode {
	case WriteFanout:
		// Fanout the post to all of this user's followers
	case ReadFanout:
		// Do nothing. Compute fanout on the fly
	}
	return resp, nil

}

func (s *FeedService) FollowUser(
	ctx context.Context,
	request *connect.Request[feedv1.FollowUserRequest],
) (*connect.Response[feedv1.FollowUserResponse], error) {
	// get the claims from the context since we parsed that and added it

	query := `
	INSERT INTO followers (
		from_user, to_user
	) VALUES (
	 	$1, $2
	)
	;
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, request.Msg.UserId)
	if err != nil {
		return nil, err
	}

	resp := &connect.Response[feedv1.FollowUserResponse]{
		Msg: &feedv1.FollowUserResponse{},
	}

	switch s.opts.fanoutMode {
	case WriteFanout:
		// to_user just gained another follower.
		// We need to fanout to_user's posts to from_user's feed.
		// We should do this asynchronously.
	case ReadFanout:
		// Do nothing. We calculate the feed for each user on the fly.
	}

	return resp, nil
}
