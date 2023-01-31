package discov

/*
服务发现代码
从ETCD读取指定Key配置，监听并更新连接池
*/
import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/client/v3"

	"google.golang.org/grpc/resolver"
)

type Builder struct {
	watcher *Watcher
	domain  string
}

func NewBuilder(client *clientv3.Client, domain string) *Builder {
	if domain == "" {
		log.Fatal("domain cannot be empty")
	}
	r := &Builder{
		watcher: NewWatcher(client),
		domain:  domain,
	}
	resolver.Register(r)
	return r
}

func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &Resolver{
		cc: cc,
	}
	notify := Notify(func(_ *clientv3.Event, address []string) {
		if len(address) == 0 {
			address = append(address, b.domain)
			log.Infof("grpc address is empty , add the application address  %+v", address)
		}
		log.Infof("grpc address change to %+v", address)
		r.Update(address)
	})
	prefix := fmt.Sprintf("/%s/%s/%s/", "xxx_discov", clusterKey, target.Endpoint)
	log.Infof("grpc address prefix is %+s", prefix)
	b.watcher.SetPrefix(prefix)
	b.watcher.AddEvent(notify)
	go b.watcher.Run()
	return r, nil
}

func (b *Builder) Scheme() string {
	return schema
}
