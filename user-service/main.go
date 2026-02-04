package main

import (
	"database/sql"
	"fmt"
	"hms/user-service/grpc_Server"
	"hms/user-service/handler"
	"hms/user-service/middleware"
	"hms/user-service/repository"
	"hms/user-service/service"
	"hms/user-service/utils"
	"net"

	pb "hms/proto/auth"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func initPostgres(cfg *utils.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	cfg := utils.LoadConfig()

	db, err := initPostgres(cfg)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	redisClient := utils.InitRedis(cfg)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, redisClient)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.Use(middleware.GinLogger())
	registerRoutes(router, userHandler, redisClient)

	// http server

	go func() {
		log.Println("User service running on :8080")
		router.Run(":8080")
	}()

	// grpc server
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}
	// grpcServer := grpc.NewServer()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_Server.LoggingInterceptor),
	)
	authServer := grpc_Server.NewAuthServer(redisClient)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	log.Println("gRPC server running on :9090")
	go grpcServer.Serve(lis)

	select {}
}
