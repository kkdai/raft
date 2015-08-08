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
	nodeList    []int     //id list exist in this network.
	term        int       //term about current time seq
	db          submittedItems

	isAlive bool //To determine if server still alive, for kill testing.

	//For candidator
	HasVoted      bool //To record if already vote others.
	acceptVoteMsg []message
}

//New a server and given a random expired time.
func NewServer(id int, role Role, nt nodeNetwork, nodeList ...int) *server {
	rand.Seed(time.Now().UnixNano())
	expiredMiliSec := rand.Intn(5) + 1
	serv := &server{id: id,
		role:        role,
		nt:          nt,
		expiredTime: expiredMiliSec,
		isAlive:     true,
		nodeList:    nodeList,
		db:          submittedItems{}}
	return serv
}

//AssignAction : Assign a assign to any of server.
//This is another thread, because the command send from outside, example app request to log.
func (sev *server) AssignAction(action datalog) {
	//TODO. Add action into logs and leader will announce to all other servers.
	switch sev.role {
	case Leader:
		//Apply to all followers
	case Candidate:
		//TBC.
	case Follower:
		//Run election to leader
		sev.requestVote(action)
	}
}

func (sev *server) Whoareyou() Role {
	return sev.role
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

//For flower -> candidate
func (sev *server) requestVote(action datalog) {
	m := message{from: sev.id,
		typ: RequestVote,
		val: action}
	for _, node := range sev.nodeList {
		m.to = node
		sev.nt.send(m)
	}

	//Send request Vote and change self to Candidate.
	sev.roleChange(Candidate)
	log.Println(" Now ID:", sev.id, " become ", sev.role)
}

func (sev *server) sendHearbit() {
	//
	for _, node := range sev.nodeList {
		hbMsg := message{from: sev.id, to: node, typ: Heartbit}
		sev.nt.send(hbMsg)
	}
}

func (sev *server) runLeaderLoop() {
	log.Println("ID:", sev.id, " Run leader loop")
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
	log.Println("ID:", sev.id, " Run candidate loop")
	//TODO. send RequestVote to all others
	recvMsg := sev.nt.recev()

	if recvMsg == nil {
		log.Println("ID:", sev.id, " no msg, return.")
		return
	}
	switch recvMsg.typ {
	case Heartbit:
		//TODO
		return
	case RequestVote:
		//TODO.
		return
	case AcceptVote:
		sev.acceptVoteMsg = append(sev.acceptVoteMsg, *recvMsg)
		if len(sev.acceptVoteMsg) > sev.majorityCount() {
			sev.roleChange(Leader)

			//TODO. send win vote to all note

		}

		return
	case WinningVote:
		//TODO
		return

	}
	//TODO. recev AcceptVote

	//TODO. check if prompt to leader.

	//TODO. If not, back to follower
}

func (sev *server) runFollowerLoop() {
	log.Println("ID:", sev.id, " Run follower loop")

	//TODO. check if leader no heartbeat to change to candidate.
	recvMsg := sev.nt.recev()

	if recvMsg == nil {
		log.Println("ID:", sev.id, " no msg, return.")
		return
	}
	switch recvMsg.typ {
	case Heartbit:
		// return
		if !sev.db.getLatestLogs().identical(recvMsg.getVal()) {
			//Data not exist, add it. (TODO)
			sev.db.add(recvMsg.getVal())
		}

		//Send it back HeartBeat
		recvMsg.to = recvMsg.from
		recvMsg.from = sev.id
		sev.nt.send(*recvMsg)
		return
	case RequestVote:
		//Handle Request Vote from candidate.
		//If doesn't vote before, will vote.
		if !sev.HasVoted {
			recvMsg.to = recvMsg.from
			recvMsg.from = sev.id
			recvMsg.typ = AcceptVote
			sev.nt.send(*recvMsg)
		} else {
			//Don't do anything if you already vote.
			//Only vote when first candidate request vote comes.
		}
	case WinningVote:
		//Clean variables.
		sev.HasVoted = false

	}
}

func (sev *server) majorityCount() int {
	return len(sev.nodeList)/2 + 1
}

func (sev *server) roleChange(newRole Role) {
	log.Println("note:", sev.id, " change role from ", sev.role, " to ", newRole)
	sev.role = newRole
}
