package repo

import (
	"context"

	"github.com/quocdaitrn/cp-user/domain/entity"
)

// UserRepo provides methods for interacting with user's data.
type UserRepo interface {
	// FindOne fetches a user from database by user id.
	FindOne(ctx context.Context, id uint) (*entity.User, error)

	// FindManyByIDs fetches list of users from database by list of user's ids.
	FindManyByIDs(ctx context.Context, ids []uint) ([]entity.User, error)

	// InsertOne inserts a new user into database.
	InsertOne(ctx context.Context, user *entity.User) error
}
