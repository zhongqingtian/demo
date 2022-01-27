module demo

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/Braveheart7854/rabbitmqPool v1.1.2
	github.com/Shopify/sarama v1.19.0
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/astaxie/session v0.0.0-20130408050157-95d7fe18579c
	github.com/coreos/bbolt v1.3.2 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-tsz v0.0.0-20180227144327-03b7d791f4fe
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xman/go.emoji v0.1.2
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.5.5
	github.com/google/gops v0.3.14
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris v0.0.2
	github.com/kataras/iris/v12 v12.1.8 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/lxmgo/config v0.0.0-20180313024057-8db99aca0f7e
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/milvus-io/milvus-sdk-go/v2 v2.0.0-alpha.2
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic/v7 v7.0.22
	github.com/onsi/ginkgo v1.7.0 // indirect
	github.com/onsi/gomega v1.4.3 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	github.com/valyala/fasthttp v1.17.0 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.etcd.io/etcd/api/v3 v3.5.1
	go.etcd.io/etcd/client/v3 v3.5.1
	go.mongodb.org/mongo-driver v1.4.6
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.38.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0

)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

replace go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.5
