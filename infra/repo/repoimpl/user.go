package repoimpl

import (
	"context"

	"github.com/pkg/errors"
	kiterrors "github.com/quocdaitrn/golang-kit/errors"
	"gorm.io/gorm"

	"github.com/quocdaitrn/cp-user/domain/entity"
	"github.com/quocdaitrn/cp-user/domain/repo"
)

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo creates and returns an instance of UserRepo.
func NewUserRepo(db *gorm.DB) repo.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) FindOne(ctx context.Context, id uint) (*entity.User, error) {
	var data entity.User

	if err := r.db.
		Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, kiterrors.ErrRepoEntityNotFound
		}
		return nil, err
	}

	return &data, nil
}

func (r *userRepo) FindManyByIDs(ctx context.Context, ids []uint) ([]entity.User, error) {
	var result []entity.User

	if err := r.db.
		Table(entity.User{}.TableName()).
		Where("id in (?)", ids).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepo) InsertOne(ctx context.Context, user *entity.User) error {
	if err := r.db.Table(user.TableName()).Create(user).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
