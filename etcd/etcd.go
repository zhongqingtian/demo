package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type EtcdClient struct {
	client *clientv3.Client
}

func CreateClient() *EtcdClient {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return nil
	}
	fmt.Println("connect to etcd success")
	// defer cli.Close()
	return &EtcdClient{client: cli}
}

// put
func (cli *EtcdClient) Put(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.client.Put(ctx, key, value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return err
	}
	return nil
}

// get

func (cli *EtcdClient) Get(key string) (*clientv3.GetResponse, error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.client.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return nil, nil
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	return resp, nil
}

// watch demo
func Watch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// watch key:q1mi change
	rch := cli.Watch(context.Background(), "q1mi") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
