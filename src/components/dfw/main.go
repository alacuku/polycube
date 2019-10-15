// dfw project main.go
package main

import (
	"fmt"

	//cm "github.com/polycube-network/src/components/dfw/clientmgr"
	fw "github.com/polycube-network/src/components/dfw/k8sfirewall"
	//node "github.com/polycube-network/src/components/dfw/node"
)

const (
	basepath = "http://node1:9000/polycube/v1"
	fw_name  = "my-fw"
)

var (
	fwAPI *fw.FirewallApiService
)

func main() {
	cfg := fw.Configuration{BasePath: basepath}
	fwAPI := fw.NewAPIClient(&cfg).FirewallApi

	_, err := fwAPI.CreateFirewallByID(nil, fw_name, fw.Firewall{Name: fw_name})
	if err != nil {
		fmt.Printf("failed to create %s\n", fw_name)
		return
	}
	fmt.Printf("created slb %s successfully\n", fw_name)

}
