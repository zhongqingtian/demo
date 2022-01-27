package discov

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"strings"
	"sync"
	"time"
)

type Watcher struct {
	lock      sync.RWMutex
	client    *clientv3.Client
	prefix    string
	addresses map[string]string
	Notifies  []Notify
}

func NewWatcher(client *clientv3.Client) *Watcher {
	return &Watcher{
		lock:      sync.RWMutex{},
		client:    client,
		addresses: make(map[string]string),
	}
}

type Notify func(ev *clientv3.Event, address []string)

func (r *Watcher) defaultNotify(event *clientv3.Event) {
	switch event.Type {
	case mvccpb.PUT:
		r.setAddress(string(event.Kv.Key), string(event.Kv.Value))
	case mvccpb.DELETE:
		r.delAddress(string(event.Kv.Key))
	}
}

func (r *Watcher) AddEvent(notify Notify) *Watcher {
	r.Notifies = append(r.Notifies, notify)
	return r
}

func (r *Watcher) SetPrefix(prefix string) *Watcher {
	r.prefix = prefix
	return r
}

func (r *Watcher) Run() {
	for _, notify := range r.Notifies {
		notify(nil, []string{})
	}
	r.addresses = make(map[string]string)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	response, err := r.client.Get(ctx, r.prefix, clientv3.WithPrefix())
	if err == nil {
		for _, kv := range response.Kvs {
			r.setAddress(string(kv.Key), string(kv.Value))
			addresses := r.getAddresses()
			for _, notify := range r.Notifies {
				notify(&clientv3.Event{Type: mvccpb.PUT, Kv: kv}, addresses)
			}
		}
	} else {
		fmt.Println("etcd get fail", err)
	}

	watch := r.client.Watch(context.Background(), r.prefix, clientv3.WithPrefix())
	for response := range watch {
		for _, event := range response.Events {
			r.defaultNotify(event)
			addresses := r.getAddresses()
			for _, notify := range r.Notifies {
				notify(event, addresses)
			}
		}
	}
}

func (r *Watcher) setAddress(key, address string) {
	if strings.TrimSpace(address) == "" {
		return
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	r.addresses[key] = address
	fmt.Println("set address", r.addresses)
}

func (r *Watcher) delAddress(key string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.addresses, key)
	fmt.Println("del address", r.addresses)
}

func (r *Watcher) getAddresses() []string {
	var addresses []string
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, address := range r.addresses {
		if strings.TrimSpace(address) == "" {
			continue
		}
		addresses = append(addresses, address)
	}
	fmt.Println("get address", addresses)
	return addresses
}
