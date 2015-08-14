package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

func newNode(name, ip string) (*node, error) {
	n := &node{
		ip:   ip,
		name: name,
	}
	return n, nil
}

var nodes map[string]*node

type node struct {
	ip   string
	name string
	rps  float64
}

func (n *node) stop() error {
	return nil
}

func (n *node) start(checkpoint bool) error {
	return nil
}

func (n *node) clone() error {
	return nil
}

func (n *node) send(n2 *node) error {
	return nil
}

func (n *node) getFill() float64 {
	f, err := redis.Float64(do("GET", fmt.Sprintf("nodes.%s.fill", n.name)))
	if err != nil {
		logrus.Error(err)
	}
	return f
}

func (n *node) getResponseTime() float64 {
	f, err := redis.Float64(do("GET", fmt.Sprintf("nodes.%s.avg", n.name)))
	if err != nil {
		logrus.Error(err)
	}
	return f
}

func loadNodes() error {
	nodes = make(map[string]*node)
	servers, err := redis.StringMap(do("HGETALL", "servers"))
	if err != nil {
		return err
	}
	for name, ip := range servers {
		nn, err := newNode(name, ip)
		if err != nil {
			return err
		}
		nodes[nn.name] = nn
	}
	return nil
}
