package service

import (
	"context"

	"github.com/quocdaitrn/cp-user/domain/entity"
)

// UserService exposes all available use cases of user's domain.
type UserService interface {
	// GetCurrentUserProfile gets current user's profile.
	GetCurrentUserProfile(ctx context.Context, req *GetCurrentUserProfileRequest) (*GetCurrentUserProfileResponse, error)

	// GetUser gets a user's profile by id.
	GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error)

	// GetUsers gets list of user's profiles by list of ids.
	GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error)

	// CreateNewUser creates a new user's profile.
	CreateNewUser(ctx context.Context, req *CreateNewUserRequest) (*CreateNewUserResponse, error)
}

// GetCurrentUserProfileRequest represents a request to get current user's profile.
type GetCurrentUserProfileRequest struct{}

// GetCurrentUserProfileResponse represents a response for getting current user's profile request.
type GetCurrentUserProfileResponse struct {
	*entity.User
}

// GetUserRequest represents a request to get user's profile by id.
type GetUserRequest struct {
	ID uint `param:"id" validate:"required"`
}

// GetUserResponse represents a response for getting user's profile request.
type GetUserResponse struct {
	*entity.User
}

// GetUsersRequest represents a request to get user's profiles by ids.
type GetUsersRequest struct {
	IDs []uint `param:"ids" validate:"required"`
}

// GetUsersResponse represents a response for getting user's profiles request.
type GetUsersResponse struct {
	Users []entity.User `json:"items"`
}

// CreateNewUserRequest represents a request to create new user's profile.
type CreateNewUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
}

// CreateNewUserResponse represents a response for creating new user's profile request.
type CreateNewUserResponse struct {
	UserID uint `json:"user_id"`
}
