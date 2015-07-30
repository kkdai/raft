package raft

import (
	"log"
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
	rand.Seed(time.Now().UnixNano())
	//expired time is random by spec.
	expiredMiliSec := rand.Intn(5) + 1
	serv := &server{id: id, role: role, nt: nt, expiredTime: expiredMiliSec}
	return serv
}

func (sev *server) run() {

	for {
		m := sev.nt.recev()
		if m == nil {
			continue
		}

		var retMsg message
		switch m.typ {
		case Heartbit:
			retMsg := sev.handleHearbit(*m)
		case RequestVote:
			retMsg := sev.handleRequestVote(*m)

		case AcceptVote:
			retMsg := sev.handleAcceptVote(*m)
		case WinningVote:
			retMsg := sev.handleWinningVote(*m)
		}

		sev.nt.send(retMsg)
	}
}

func (sev *server) handleHearbit(m message) message {

	switch sev.role {
	case Folloer:
	case Candidate:
	case Leader:
	}
	return m
}

func (sev *server) handleRequestVote(m message) message {

	switch sev.role {
	case Folloer:
	case Candidate:
	case Leader:
	}
	return m
}

func (sev *server) handleAcceptVote(m message) message {

	switch sev.role {
	case Folloer:
	case Candidate:
	case Leader:
	}
	return m
}

func (sev *server) handleWinningVote(m message) message {

	switch sev.role {
	case Folloer:
	case Candidate:
	case Leader:
	}
	return m
}

func (sev *server) roleChange(newRole Role) {
	log.Println("note:", sev.id, " change role from ", sev.role, " to ", newRole)
	sev.role = newRole
}
