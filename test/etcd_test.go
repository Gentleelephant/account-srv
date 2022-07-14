package test

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		t.Error(err)
	}
	cli.Put(context.TODO(), "account-srv-00001", "localhost:8080")
	resp, err := cli.Get(context.TODO(), "account-srv", clientv3.WithPrefix())
	if err != nil {
		// handle error!
		t.Error(err)
	}
	t.Log(string(resp.Kvs[0].Value))
	defer cli.Close()
}
