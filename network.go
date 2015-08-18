package raft

import (
	"log"
	"time"
)

func CreateNetwork(nodes ...int) *network {
	nt := network{recvQueue: make(map[int]chan Message, 0)}

	for _, node := range nodes {
		nt.recvQueue[node] = make(chan Message, 1024)
	}

	return &nt
}

type network struct {
	recvQueue map[int]chan Message
}

func (n *network) getNodeNetwork(id int) nodeNetwork {
	return nodeNetwork{id: id, net: n}
}

func (n *network) sendTo(m Message) {
	log.Println("Send msg from:", m.from, " send to", m.to, " val:", m.val, " typ:", m.typ)
	n.recvQueue[m.to] <- m
}

func (n *network) recevFrom(id int) *Message {
	select {
	case retMsg := <-n.recvQueue[id]:
		log.Println("Recev msg from:", retMsg.from, " send to", retMsg.to, " val:", retMsg.val, " typ:", retMsg.typ)
		return &retMsg
	case <-time.After(time.Second):
		// log.Println("id:", id, " don't get message.. time out.")
		return nil
	}
}

type nodeNetwork struct {
	id  int
	net *network
}

func (n *nodeNetwork) send(m Message) {
	n.net.sendTo(m)
}

func (n *nodeNetwork) recev() *Message {
	return n.net.recevFrom(n.id)
}
