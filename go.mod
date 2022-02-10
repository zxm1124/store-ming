module github.com/zxm1124/store-ming

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/sirupsen/logrus v1.8.1
	github.com/zxm1124/component-base v0.1.9
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.22.5
)

replace github.com/zxm1124/component-base v0.1.9 => ../component-base
