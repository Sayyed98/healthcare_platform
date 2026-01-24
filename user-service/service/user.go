package service

import (
	"errors"
	"hms/user-service/model"
	"hms/user-service/repository"
	"hms/user-service/utils"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	repo  *repository.UserRepository
	redis *redis.Client
}

func NewUserService(repo *repository.UserRepository, redis *redis.Client) *UserService {
	return &UserService{repo: repo, redis: redis}
}

func (s *UserService) Register(req model.RegisterRequest) (*model.UserResponse, error) {
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: hash,
		Role:         "PATIENT",
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *UserService) Login(req model.LoginRequest) (string, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(user.PasswordHash, req.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	sessionID := uuid.New().String()
	if err := utils.CreateSession(s.redis, sessionID, user.ID); err != nil {
		return "", err
	}

	return sessionID, nil
}
