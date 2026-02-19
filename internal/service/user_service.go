package service

import (
	"context"
	"errors"

	"github.com/crisywini/go-products-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	queries *repository.Queries
}

func NewUserService(queries *repository.Queries) *UserService {
	return &UserService{queries: queries}
}

func (s *UserService) Register(ctx context.Context, name, email, password string) (repository.User, error) {
	_, err := s.queries.GetUserByEmail(ctx, email)
	if err == nil {
		return repository.User{}, errors.New("Property: email already exists!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return repository.User{}, errors.New("Error hashing the password!")
	}

	user, err := s.queries.CreateUser(ctx, repository.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return repository.User{}, errors.New("Error creating a new user!")
	}
	return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (repository.User, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return repository.User{}, errors.New("Invalid Credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return repository.User{}, errors.New("Invalid Credentials")
	}
	return user, nil
}
