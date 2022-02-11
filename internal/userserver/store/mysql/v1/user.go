package v1

import (
	"context"
	"errors"
	"github.com/zxm1124/component-base/pkg/code"
	meta "github.com/zxm1124/component-base/pkg/meta/v1"
	"github.com/zxm1124/component-base/pkg/util/gormutil"
	model "github.com/zxm1124/store-ming/internal/userserver/model/v1"
	"gorm.io/gorm"
)

type Users struct {
	Db *gorm.DB
}

func (u *Users) Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error {
	return u.Db.Create(&user).Error
}

func (u *Users) Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error {
	return u.Db.Save(&user).Error
}

func (u *Users) Delete(ctx context.Context, username string, opts meta.DeleteOptions) error {
	// true: 永久删除
	if opts.Unscoped {
		u.Db = u.Db.Unscoped()
	}

	err := u.Db.Where("name = ?", username).Delete(&model.User{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return code.ErrDatabase
	}

	return nil
}

func (u *Users) DeleteCollection(ctx context.Context, usernames []string, opts meta.DeleteOptions) error {
	if opts.Unscoped {
		u.Db = u.Db.Unscoped()
	}

	err := u.Db.Where("name in (?)", usernames).Error
	if err != nil {
		return code.ErrDatabase
	}

	return nil
}

func (u *Users) Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error) {
	user := &model.User{}
	err := u.Db.Where("name = ? and status = 1", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.ErrUserNotFound
		}

		return nil, code.ErrDatabase
	}

	return user, nil
}

func (u *Users) List(ctx context.Context, opts meta.ListOptions) (*model.UserList, error) {
	ret := &model.UserList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	d := u.Db.Where("instanceID like ? and status = 1", "%user-").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Limit(-1).
		Count(&ret.Total)

	return ret, d.Error
}
