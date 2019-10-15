// client_manager.go
package ClientManager

import (
	"fmt"

	"github.com/deckarep/golang-set"
	fw "github.com/polycube-network/src/components/dfw/k8sfirewall"
)

type clientmanager struct {
	cliapi []*fw.FirewallApiService
	sid    map[string]int
	ids    mapset.Set
}

type ClientManager interface {
	GetClientAPI(server string) *fw.FirewallApiService
	RegisterServer(server string) int
}

func (cm *clientmanager) GetClientAPI(server string) *fw.FirewallApiService {
	id := cm.sid[server]
	if id >= len(cm.cliapi) {
		return nil
	}

	return cm.cliapi[id]
}

func (cm *clientmanager) RegisterServer(server string) {
	if _, ok := cm.sid[server]; ok {
		fmt.Printf("server %s is already registered.", server)
		return
	}

	emptyset := mapset.NewSet()
	if cm.ids.Equal(emptyset) {
		fmt.Println("running out of ids")
		return
	}
	cm.sid[server] = cm.ids.Pop().(int)
}
