/*
register.go
服务注册代码
从Discov读取配置，创建etcd连接客户端，将服务注册到etcd
*/
package discov

import (
	"context"
	"demo/utils"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"time"
)

type register struct {
	client      *clientv3.Client
	serviceName string
	serviceAddr string
	ttl         int64
}

func NewRegister(client *clientv3.Client) *register {
	r := &register{client: client}
	return r
}

//Register 服务注册
func (r *register) Discov(serviceName string, port string) <-chan error {
	e := make(chan error)
	split := strings.Split(port, ":")
	if len(split) > 0 {
		port = split[len(split)-1]
	}
	var err error
	var ip string

	for {
		ip, err = utils.IpAddrList()
		if err != nil {
			e <- err
			return e
		} else if ip == "" {
			time.Sleep(time.Second * 1)
		}
		break
	}

	r.serviceAddr = ip + ":" + port
	r.serviceName = serviceName

	ch := make(chan struct{}, 1)
	go r.async(ch, e)
	return e
}

func (r *register) async(ch chan struct{}, e chan error) {
	ch <- struct{}{}
	for {
		select {
		case <-ch:
			err := r.keepAlive(ch)
			if err != nil {
				e <- err
				time.Sleep(time.Second * time.Duration(defaultTTL))
				if len(ch) == 0 {
					ch <- struct{}{}
				}
			}
		}
	}
}

func (r *register) keepAlive(restart chan struct{}) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		if err != nil {
			cancel()
		}
	}()
	//创建租约
	lease, err := r.client.Grant(ctx, defaultTTL)
	if err != nil {
		return err
	}
	//设置键值对
	_, err = r.client.Put(ctx, r.key(), r.serviceAddr, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	//续签
	ch, err := r.client.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			if v := <-ch; v == nil {
				restart <- struct{}{}
				return
			}
		}
	}()
	return nil
}

func (r *register) key() string {
	return fmt.Sprintf("/%s/%s/%s/%s", "docer_discov", clusterKey, r.serviceName, r.serviceAddr)
}
