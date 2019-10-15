package node

import (
	//"fmt"
	"time"
	"log"
	"context"

	"github.com/coreos/etcd/clientv3"
    config "github.com/polycube-network/src/components/dfw/config"
)

const (
	dialtimeout time.Duration = 2 * time.Second
	default_cfg string = "/etc/dfw.cfg"

	// node path: /config/dfw/default_cluster/<node_name>/attributes
	node_path string = "/config/dfw/default_cluster/"
)

var (
    default_endpoints []string = []string{"localhost:2379"}
	endpoints []string
)

type node struct {
	name    string
	IP      string
	cluster string
	zone    string
}

type NodeOps interface {
	join() (*node, error)
	leave() error
}

func (n *node) join() (*node, error) {
	return nil, nil
}

func main() {
	eList, _ := config.ReadCfg(default_cfg)
	if eList != nil {
		endpoints = eList
	} else {
		endpoints = default_endpoints
	}

	cli, error := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialtimeout,
	})
    if error != nil {
		log.Fatal(error)
	}
	defer cli.Close()

	key := node_path
	value := "value"
	_, error = cli.Put(context.TODO(), key, value)
	if error != nil {
		log.Fatal(error)
	}
}
