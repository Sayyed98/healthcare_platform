package main

import (
	"database/sql"
	"fmt"
	"log"

	"hms/hospital-service/grpc_client"
	"hms/hospital-service/handler"
	"hms/hospital-service/middleware"
	"hms/hospital-service/repository"
	"hms/hospital-service/service"
	"hms/hospital-service/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := utils.LoadConfig()

	db, err := initPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	authClient, err := grpc_client.NewAuthClient("localhost:9090")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewHospitalRepository(db)
	svc := service.NewHospitalService(repo, authClient)
	h := handler.NewHospitalHandler(svc)

	r := gin.Default()
	registerRoutes(r, h, middleware.AuthMiddleware(authClient))

	r.Run(":8081")
}

func initPostgres(cfg *utils.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
