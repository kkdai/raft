package raft

import (
	"log"
	"math/rand"
	"time"
)

type Role int

const (
	Follower Role = iota + 1
	Candidate
	Leader
)

type server struct {
	id          int
	expiredTime int //Hearbit expired time (by millisecond.)
	role        Role
	nt          nodeNetwork
	msgRecvTime time.Time //Message receive time
	isAlive     bool      //To determine if server still alive, for kill testing.
	nodeList    []int     //id list exist in this network.
	term        int       //term about current time seq
}

//New a server and given a random expired time.
func NewServer(id int, role Role, nt nodeNetwork, nodeList ...int) *server {
	rand.Seed(time.Now().UnixNano())
	expiredMiliSec := rand.Intn(5) + 1
	serv := &server{id: id, role: role, nt: nt, expiredTime: expiredMiliSec, isAlive: true, nodeList: nodeList}
	return serv
}

//AssignAction : Assign a assign to any of server.
func (sev *server) AssignAction(action string) {
	//TODO. Add action into logs and leader will announce to all other servers.
}

func (sev *server) RunServerLoop() {

	for {
		switch sev.role {
		case Leader:
			sev.runLeaderLoop()
		case Candidate:
			sev.runCandidateLoop()
		case Follower:
			sev.runFollowerLoop()
		}

		//timer base on milli-second.
		time.Sleep(time.Millisecond)
	}
}

func (sev *server) sendHearbit() {

	for _, node := range sev.nodeList {
		hbMsg := message{from: sev.id, to: node, typ: Heartbit, val: "HB"}
		sev.nt.send(hbMsg)
	}
}

func (sev *server) runLeaderLoop() {
	sev.sendHearbit()

	recevMsg := sev.nt.recev()
	if recevMsg == nil {
		return
	}
	switch recevMsg.typ {
	case Heartbit:
		return
	}

	//TODO. assign value to followers

	//TODO. if get bigger TERM request, back to follower
}

func (sev *server) runCandidateLoop() {

	//TODO. send RequestVote to all others

	//TODO. recev AcceptVote

	//TODO. check if prompt to leader.

	//TODO. If not, back to follower
}

func (sev *server) runFollowerLoop() {

	//TODO. check if leader no heartbeat to change to candidate.

	//TODO. check action from leader to add into committed logs
}

func (sev *server) roleChange(newRole Role) {
	log.Println("note:", sev.id, " change role from ", sev.role, " to ", newRole)
	sev.role = newRole
}
