package store

import (
	"context"
	meta "github.com/zxm1124/component-base/pkg/meta/v1"
	model "github.com/zxm1124/store-ming/internal/userserver/model/v1"
)

// UserStore 定义用户服务的dao接口
type UserStore interface {
	Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error
	Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error
	Delete(ctx context.Context, username string, opts meta.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts meta.UpdateOptions) error
	Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error)
	List(ctx context.Context, opts meta.ListOptions) (*model.UserList, error)
}
