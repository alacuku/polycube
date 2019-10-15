package node

import (
	"errors"
	"fmt"
)

type nodemgr struct {
	nodes map[string]*node
}

type NodeManager interface {
	add_node(n *node) error
	remove_node(n *node) error
}

func (m *nodemgr) add_node(n *node) error {
	_, ok := m.nodes[n.name]
	if ok {
		fmt.Println("node %s exists", n.name)
		return errors.New("node exists")
	}
	m.nodes[n.name] = n
	return nil
}

func (m *nodemgr) remove_node(n *node) error {
	_, ok := m.nodes[n.name]
	if !ok {
		fmt.Println("node %s doesn't exist", n.name)
		return errors.New("node doesn't exist")
	}
	delete(m.nodes, n.name)
	return nil
}
