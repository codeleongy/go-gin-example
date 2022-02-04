module github.com/leong-y/go-gin-example

go 1.17

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/astaxie/beego v1.12.3
	github.com/boombuler/barcode v1.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/swaggo/gin-swagger v1.4.0
	github.com/swaggo/swag v1.7.8
	github.com/unknwon/com v1.0.1
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/ugorji/go/codec v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20220126234351-aa10faf2a1f8 // indirect
	golang.org/x/net v0.0.0-20220127074510-2fabfed7e28f // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.9 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/leong-y/go-gin-example/conf/app.ini => ../go-gin-example-v2/conf/app.ini
	github.com/leong-y/go-gin-example/middleware => ../go-gin-example-v2/middleware
	github.com/leong-y/go-gin-example/middleware/jwt => ../go-gin-example-v2/middleware/jwt
	github.com/leong-y/go-gin-example/models => ../go-gin-example-v2/models
	github.com/leong-y/go-gin-example/pkg/app => ../go-gin-example-v2/pkg/app
	github.com/leong-y/go-gin-example/pkg/e => ../go-gin-example-v2/pkg/e
	github.com/leong-y/go-gin-example/pkg/file => ../go-gin-example-v2/pkg/file
	github.com/leong-y/go-gin-example/pkg/gredis => ../go-gin-example-v2/pkg/gredis
	github.com/leong-y/go-gin-example/pkg/logging => ../go-gin-example-v2/pkg/logging
	github.com/leong-y/go-gin-example/pkg/setting => ../go-gin-example-v2/pkg/setting
	github.com/leong-y/go-gin-example/pkg/util => ../go-gin-example-v2/pkg/util
	github.com/leong-y/go-gin-example/routers => ../go-gin-example-v2/routers
	github.com/leong-y/go-gin-example/routers/api => ../go-gin-example-v2/routers/api
	github.com/leong-y/go-gin-example/routers/api/version1 => /root/go-projects/go-gin-example-learn/go-gin-example-v2/routers/api/version1
)
