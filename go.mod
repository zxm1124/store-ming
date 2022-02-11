module github.com/zxm1124/store-ming

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/zxm1124/component-base v0.1.9
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
)

replace github.com/zxm1124/component-base v0.1.9 => ../component-base
