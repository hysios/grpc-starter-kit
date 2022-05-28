module github.com/hysios/grpc-starter-kit

go 1.18

replace (
	github.com/hysios/cloud v0.0.0 => ../../packages/cloud
	github.com/hysios/config v0.0.0 => ../../packages/config
)

require (
	github.com/go-oauth2/oauth2/v4 v4.5.1
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.2
	github.com/hysios/mx/platform v0.0.0-20220526081806-fa268bc178a9
	github.com/hysios/utils v0.0.11
	github.com/joho/godotenv v1.4.0
	google.golang.org/genproto v0.0.0-20220526192754-51939a95c655
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gorm.io/gorm v1.23.5
)

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.4 // indirect
	github.com/alibabacloud-go/darabonba-openapi v0.1.18 // indirect
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/alibabacloud-go/dm-20151123 v1.0.2 // indirect
	github.com/alibabacloud-go/dysmsapi-20170525/v2 v2.0.9 // indirect
	github.com/alibabacloud-go/endpoint-util v1.1.0 // indirect
	github.com/alibabacloud-go/openapi-util v0.0.11 // indirect
	github.com/alibabacloud-go/tea v1.1.17 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.3 // indirect
	github.com/alibabacloud-go/tea-xml v1.1.2 // indirect
	github.com/aliyun/credentials-go v1.1.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/casbin/casbin/v2 v2.43.2 // indirect
	github.com/casbin/gorm-adapter/v3 v3.5.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/clbanning/mxj/v2 v2.5.5 // indirect
	github.com/denisenkom/go-mssqldb v0.12.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-oauth2/redis/v4 v4.1.1 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang-sql/sqlexp v0.0.0-20170517235910-f1bb20e5a188 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hysios/cloud v0.0.0 // indirect
	github.com/hysios/config v0.0.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.11.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.10.0 // indirect
	github.com/jackc/pgx/v4 v4.15.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mcuadros/go-defaults v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/naucon/casbin-fs-adapter v0.1.0 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/tidwall/btree v0.0.0-20191029221954-400434d76274 // indirect
	github.com/tidwall/buntdb v1.1.2 // indirect
	github.com/tidwall/gjson v1.12.1 // indirect
	github.com/tidwall/grect v0.0.0-20161006141115-ba9a043346eb // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/rtree v0.0.0-20180113144539-6cd427091e0e // indirect
	github.com/tidwall/tinyqueue v0.0.0-20180302190814-1e39f5511563 // indirect
	github.com/tjfoc/gmsm v1.3.2 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/config v1.4.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.10 // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	gopkg.in/oauth2.v3 v3.12.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
	gorm.io/driver/postgres v1.3.3 // indirect
	gorm.io/driver/sqlserver v1.3.2 // indirect
	gorm.io/plugin/dbresolver v1.1.0 // indirect
)
