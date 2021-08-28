package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"etcd.io/etcd/client/v3"
)

func main() {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379", "http://127.0.0.1:22379", "http://127.0.0.1:32379"},
		DialTimeout: 2 * time.Second,
		Username:    "u1",
		Password:    "p",
	})

	if err != nil {
		fmt.Printf("failed to create a new client: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: %v\n", c)

	kv := clientv3.NewKV(c)

	for {
		fmt.Printf("reading k1\n")

		rsp, err := kv.Get(context.TODO(), "k1")
		if err != nil {
			fmt.Printf("Get failed: %s\n", err)
		}

		if len(rsp.Kvs) != 1 {
			fmt.Printf("hmm, length of Kvs is %d\n", len(rsp.Kvs))
		} else {
			fmt.Printf("value of k1: %s\n", string(rsp.Kvs[0].Value))
		}
		<-time.After(1 * time.Second)
	}
}
