package v1

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	meta "github.com/zxm1124/component-base/pkg/meta/v1"
	model "github.com/zxm1124/store-ming/internal/userserver/model/v1"
	"github.com/zxm1124/store-ming/pkg/db"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

var users *Users

func TestDelete(t *testing.T) {
	err := users.Delete(context.TODO(), "zxm1124", meta.DeleteOptions{Unscoped: false})
	if err != nil {
		fmt.Printf("delete user failed, err: %v", err)
		return
	}

	fmt.Println("删除成功")
}

//func TestCreate(t *testing.T) {
//
//	initConf()
//
//	id := idutil.GetDistributeID()
//	err := users.Create(context.TODO(), &model.User{
//		ObjectMeta: meta.ObjectMeta{
//			ID:         idutil.GetDistributeID(),
//			InstanceID: idutil.GetInstanceID(id, "user-"),
//			Name:       "139201124",
//		},
//		NickName:  "zxm1124",
//		Password:  "123456",
//		Email:     "xinmaozhu@foxmai.com",
//		Phone:     "13055722970",
//		LoginedAt: time.Now(),
//		IsAdmin:   0,
//		Status:    1,
//	}, meta.CreateOptions{})
//
//	if err != nil {
//		logrus.Errorf("create.err: %v", err)
//	}
//
//	fmt.Println("创建成功")
//}

func initConf() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	opt := &db.Options{
		Username:              "root",
		Password:              "123456",
		Database:              "store_ming",
		Host:                  "42.192.131.10:3306",
		MaxConnectionLifeTime: 100,
		MaxOpenConnections:    10,
		MaxIdleConnections:    10,
		Logger:                newLogger,
	}
	d, err := db.New(opt)
	if err != nil {
		logrus.Errorf("init.Db.config failed,err: %v\n", err)
	}

	err = d.AutoMigrate(&model.User{})
	if err != nil {
		logrus.Errorf("init.AutoMigrate failed,err: %v\n", err)
	}

	users = &Users{
		Db: d,
	}
}
