package serviceimpl

import (
	"context"

	kitcontext "github.com/quocdaitrn/golang-kit/context"
	kiterrors "github.com/quocdaitrn/golang-kit/errors"
	"github.com/quocdaitrn/golang-kit/validator"
	"github.com/viettranx/service-context/core"

	"github.com/quocdaitrn/cp-user/domain/entity"
	"github.com/quocdaitrn/cp-user/domain/repo"
	"github.com/quocdaitrn/cp-user/domain/service"
)

type userService struct {
	userRepo  repo.UserRepo
	validator validator.Validator
}

// NewUserService creates and returns an instance of UserService.
func NewUserService(
	userRepo repo.UserRepo,
	validator validator.Validator,
) service.UserService {
	return &userService{
		userRepo:  userRepo,
		validator: validator,
	}
}

// GetCurrentUserProfile gets current user's profile.
func (s *userService) GetCurrentUserProfile(ctx context.Context, req *service.GetCurrentUserProfileRequest) (*service.GetCurrentUserProfileResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	uid := kitcontext.UIDFromContext(ctx)
	id, _ := core.FromBase58(uid.Sub)
	requesterID := uint(id.GetLocalID())

	user, err := s.userRepo.FindOne(ctx, requesterID)
	if err != nil {
		return nil, err
	}

	user.Mask()
	return &service.GetCurrentUserProfileResponse{User: user}, nil
}

// GetUser gets a user's profile by id.
func (s *userService) GetUser(ctx context.Context, req *service.GetUserRequest) (*service.GetUserResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	user, err := s.userRepo.FindOne(ctx, req.ID)
	if err != nil {
		if err == kiterrors.ErrRepoEntityNotFound {
			return nil, kiterrors.ErrNotFound
		}

		return nil, err
	}

	return &service.GetUserResponse{User: user}, nil
}

// GetUsers gets list of user's profiles by list of ids.
func (s *userService) GetUsers(ctx context.Context, req *service.GetUsersRequest) (*service.GetUsersResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	users, err := s.userRepo.FindManyByIDs(ctx, req.IDs)
	if err != nil {
		return nil, err
	}

	return &service.GetUsersResponse{Users: users}, nil
}

// CreateNewUser creates a new user's profile.
func (s *userService) CreateNewUser(ctx context.Context, req *service.CreateNewUserRequest) (*service.CreateNewUserResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, kiterrors.WithStack(err)
	}

	user := &entity.User{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Phone:      req.Phone,
		Avatar:     "",
		Gender:     entity.GenderUnknown, // TODO: do not hardcode
		SystemRole: entity.RoleUser,
		Status:     entity.StatusActive,
	}
	err := s.userRepo.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &service.CreateNewUserResponse{UserID: user.ID}, nil
}
