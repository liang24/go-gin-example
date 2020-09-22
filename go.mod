module github.com/liang24/go-gin-example

go 1.15

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20200505102242-63963976dee0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.2
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.61.0
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/ugorji/go v1.1.8 // indirect
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli v1.22.4 // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200922070232-aee5d888a860 // indirect
	golang.org/x/tools v0.0.0-20200921210052-fa0125251cc4 // indirect
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
