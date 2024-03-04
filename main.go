package main

import (
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.203.8:12379", "192.168.203.8:22379", "192.168.203.8:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// print error
		fmt.Println(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cli.Put(ctx, "my-etcd-test-key", "Hello World!")
	cancel()
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		resp, err := cli.Get(ctx, "my-etcd-test-key")
		cancel()
		if err != nil {
			fmt.Println(err)
		}
		for _, ev := range resp.Kvs {
			fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		}
		time.Sleep(1 * time.Second)
	}

}
