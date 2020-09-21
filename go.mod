module github.com/liang24/go-gin-example

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.61.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.1.8 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200917073148-efd3b9a0ff20 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/liang24/go-gin-example/conf => ./pkg/conf
	github.com/liang24/go-gin-example/middleware => ./middleware
	github.com/liang24/go-gin-example/models => ./models
	github.com/liang24/go-gin-example/pkg/setting => ./pkg/setting
	github.com/liang24/go-gin-example/routers => ./routers
)