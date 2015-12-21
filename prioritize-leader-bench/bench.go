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

	mode := client.EndpointSelectionDefault
	if prioritize {
		mode = client.EndpointSelectionPrioritizeLeader
	}

	cfg := client.Config{
		Endpoints: []string{"http://172.31.41.106:12379", "http://172.31.41.103:12379", "http://172.31.41.104:12379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
		SelectionMode:   mode,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)

	for i := 0; i < 10000; i++ {
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
