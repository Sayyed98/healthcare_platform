package grpc_Server

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	start := time.Now()

	// Call actual RPC
	resp, err := handler(ctx, req)

	duration := time.Since(start)

	if err != nil {
		st, _ := status.FromError(err)
		log.Printf(
			"[gRPC] method=%s status=%s duration=%s error=%v",
			info.FullMethod,
			st.Code(),
			duration,
			st.Message(),
		)
		return nil, err
	}

	log.Printf(
		"[gRPC] method=%s status=OK duration=%s",
		info.FullMethod,
		duration,
	)

	return resp, nil
}
