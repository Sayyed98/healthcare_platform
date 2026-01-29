package grpc_client

import (
	"context"
	pb "hms/proto/auth"
	"time"

	"google.golang.org/grpc"
)

type AuthClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(addr string) (*AuthClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &AuthClient{client: pb.NewAuthServiceClient(conn)}, nil
}

func (a *AuthClient) ValidateSession(ctx context.Context, sessionID string) (*pb.ValidateSessionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return a.client.ValidateSession(ctx, &pb.ValidateSessionRequest{
		SessionId: sessionID,
	})
}
