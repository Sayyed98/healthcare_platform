package grpc_Server

import (
	"context"
	"errors"
	pb "hms/proto/auth"
	"hms/user-service/utils"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	rdb *redis.Client
}

func NewAuthServer(rdb *redis.Client) *AuthServer {
	return &AuthServer{rdb: rdb}
}

func (s *AuthServer) ValidateSession(ctx context.Context, req *pb.ValidateSessionRequest) (*pb.ValidateSessionResponse,
	error) {

	userID, err := utils.GetSession(s.rdb, req.SessionId)
	if err != nil {
		return nil, errors.New("invalid session")
	}

	userid, _ := strconv.ParseInt(userID, 10, 64)
	return &pb.ValidateSessionResponse{UserId: userid,
		Role: "PATIENT",
	}, nil
}
