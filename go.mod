module demo

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/Braveheart7854/rabbitmqPool v1.1.2
	github.com/astaxie/session v0.0.0-20130408050157-95d7fe18579c
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/dgryski/go-tsz v0.0.0-20180227144327-03b7d791f4fe
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris v0.0.2
	github.com/kataras/iris/v12 v12.1.8 // indirect
	github.com/lxmgo/config v0.0.0-20180313024057-8db99aca0f7e
	github.com/milvus-io/milvus-sdk-go/v2 v2.0.0-alpha.2
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic/v7 v7.0.22
	github.com/prometheus/client_golang v1.8.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/vakenbolt/go-test-report v0.9.3 // indirect
	github.com/valyala/fasthttp v1.17.0 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	google.golang.org/grpc v1.38.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.4.0

)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

replace go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.5
