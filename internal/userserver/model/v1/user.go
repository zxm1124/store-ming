package v1

import (
	meta "github.com/zxm1124/component-base/pkg/meta/v1"
	"github.com/zxm1124/component-base/pkg/util/idutil"
	"gorm.io/gorm"
	"time"
)

// User 存储用户信息的结构体
type User struct {
	meta.ObjectMeta `json:"metadata,omitempty"`

	// Required: true
	NickName string `json:"nickName" gorm:"column:nickname" validate:"required,min=6,max=30"`

	// Required: true
	Password string `json:"password,omitempty" gorm:"column:password" validate:"required,min=6,max=30"`

	// Required: true
	Email string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`

	Phone string `json:"phone" gorm:"column:phone" validate:"omitempty"`

	IsAdmin int `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`

	IsMember int `json:"isMember,omitempty" gorm:"column:isMember" validate:"omitempty"`

	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`

	Status int `json:"status" gorm:"column:status" validate:"omitempty"`
}

// UserList 存储用户列表信息的结构体
type UserList struct {
	meta.ListMeta `json:",inline"`
	Items         []*User `json:"items"`
}

// TableName 返回用户表名
func (u *User) TableName() string {
	return "user"
}

// AfterCreate 创建用户资源ID
func (u *User) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = idutil.GetInstanceID(u.ID, "user-")

	return tx.Save(u).Error
}
