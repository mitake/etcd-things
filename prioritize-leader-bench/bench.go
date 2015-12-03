package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mitake/etcd/client"
)

func main() {
	prioritize, err := strconv.ParseBool(os.Args[1])
	if err != nil {
		fmt.Printf("invalid flag: %s\n", os.Args[1])
		os.Exit(1)
	}

	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:12379", "http://127.0.0.1:22379", "http://127.0.0.1:32379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
		PrioritizeLeader:        prioritize,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)

	for i := 0; i < 1000; i++ {
		k := fmt.Sprintf("/foo%d", i)
		_, err := kapi.Set(context.Background(), k, "bar", nil)
		if err != nil {
			log.Fatal(err)
			log.Fatal("exiting\n")
			os.Exit(1)
		}
	}

	log.Printf("done\n")
}
