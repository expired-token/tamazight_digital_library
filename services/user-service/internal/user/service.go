package user

import (
	"context"
	"errors"
	"fmt"
	"services/user-service/internal/model"

	"gorm.io/gorm"
)

// Define custom exceptions
type UserNotFoundError struct {
	ID       uint
	Username string
}

func (e *UserNotFoundError) Error() string {
	switch {
	case e.ID != 0:
		return fmt.Sprintf("User with ID '%q' is not found", e.ID)
	case e.Username != "":
		return fmt.Sprintf("User with username '%q' is not found", e.Username)
	default:
		return "User not found"
	}
}

type UsernameTakenError struct {
	Username string
}

func (e *UsernameTakenError) Error() string {
	return fmt.Sprintf("The username '%q' is already taken", e.Username)
}

type RequiredParameterNotFoundError struct {
	Parameter string
}

func (e *RequiredParameterNotFoundError) Error() string {
	return fmt.Sprintf("Required parameter not found: %q", e.Parameter)
}

// Define Service Interface
type IUserService interface {
	// GetId(ctx context.Context, id uint) (*model.User, error)
	SearchByUsername(ctx context.Context, username string) (*model.User, error)
	// TODO - FindBySpokenDialects(ctx context.Context, spoken_dialects []string) (*[]model.User, error)
	// TODO - FindBySpokenLanguages(ctx context.Context, spoken_languages []string) (*[]model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, id uint, user *model.User) error
	Delete(ctx context.Context, id uint) error
}

type UserService struct {
	repo IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{repo: userRepository}
}

func (s *UserService) SearchByUsername(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, &RequiredParameterNotFoundError{Parameter: "username"}
	}

	user, err := s.repo.FindByUsername(ctx, username)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &UserNotFoundError{Username: username}
	}
	if err != nil {
		return nil, fmt.Errorf("An error has occured: %w", err)
	}

	return user, nil
}

func (s *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if user.Username == "" {
		return nil, &RequiredParameterNotFoundError{Parameter: "username"}
	}

	user, err := s.repo.Create(ctx, user)

	if errors.Is(err, gorm.ErrDuplicatedKey) { // TODO - INVESTIGATE why ErrDuplicatedKey isn't handled.
		return nil, &UsernameTakenError{Username: user.Username}
	}
	if err != nil {
		return nil, fmt.Errorf("Internal server error: %w", err)
	}

	return user, nil
}

func (s *UserService) Update(ctx context.Context, id uint, user *model.User) error {
	userExists, err := s.repo.ExistsById(ctx, id)

	if err != nil {
		return fmt.Errorf("internal server error: '%w", err)
	}
	if !userExists {
		return &UserNotFoundError{ID: id}
	}

	user.ID = id

	if err := s.repo.Update(ctx, user); err != nil {
		return fmt.Errorf("Internal server error: %w", err)
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id uint) error {
	userExists, err := s.repo.ExistsById(ctx, id)

	if err != nil {
		return fmt.Errorf("Internal server error: '%w'", err)
	}
	if !userExists {
		return &UserNotFoundError{ID: id}
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
