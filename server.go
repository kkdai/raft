package raft

import (
	"math/rand"
	"time"
)

type Role int

const (
	Folloer Role = iota + 1
	Candidate
	Leader
)

type server struct {
	id          int
	expiredTime int
	role        Role
	nt          nodeNetwork
	msgRecvTime time.Time
}

func NewServer(id int, role Role, nt nodeNetwork) *server {
	r := rand.New(time.Now().UnixNano())
	expiredMiliSec := rand.Intn(5)
	serv := &server{id: id, role: role, nt: nt, expiredTime: expiredMiliSec}
	return serv
}

func (sev *server) run() {

	for {
		m := serv.nt.recev()
		if m == nil {
			continue
		}

		switch m.typ {
		case Heartbit:
		case RequestVote:
		case AcceptVote:
		case WinningVote:

		}
	}
}
