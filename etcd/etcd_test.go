package etcd

import (
	"context"
	"testing"
	"time"
)

func TestCreateEtcdClient(t *testing.T) {
	cli := CreateClient()
	/*cli.Put("abc", "123")
	cli.Put("abc", "1234")
	cli.Get("abc")*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rsp, err := cli.client.MemberList(ctx)
	t.Log(err)
	t.Log(rsp)
}

func TestWatch(t *testing.T) {
	Watch()
}
