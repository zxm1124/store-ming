package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Options struct {
	Username              string
	Password              string
	Database              string
	Host                  string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogeLevel             int
	Logger                logger.Interface
}

func New(opt *Options) (*gorm.DB, error) {
	//  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Database,
		true,
		"Local",
	)
	//gorm.Open("mysql", url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: opt.Logger,
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(opt.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(opt.MaxOpenConnections)
	sqlDb.SetConnMaxLifetime(opt.MaxConnectionLifeTime)

	return db, nil
}
