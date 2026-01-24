package main

import (
	"database/sql"
	"fmt"
	"hms/user-service/repository"
	"hms/user-service/service"
	"hms/user-service/utils"

	"hms/user-service/handler"

	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	registerRoutes(router, userHandler)

	log.Println("User service running on :8080")
	router.Run(":8080")

}
