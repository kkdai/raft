package raft

import (
	"log"
	"testing"
	"time"
)

func TestBasicNetwork(t *testing.T) {
	log.Println("TestBasicNetowk........................")
	nt := CreateNetwork(1, 3, 5, 2, 4)
	go func() {
		nt.recevFrom(5)
		nt.recevFrom(1)
		nt.recevFrom(3)
		nt.recevFrom(2)
		m := nt.recevFrom(4)
		if m == nil {
			t.Errorf("No message detected.")
		}
	}()

	//m1 := message{from: 3, to: 1, typ: Prepare, seq: 1, preSeq: 0, val: "m1"}
	//nt.sendTo(m1)
	//m2 := message{from: 5, to: 3, typ: Accept, seq: 2, preSeq: 1, val: "m2"}
	//nt.sendTo(m2)
	//m3 := message{from: 4, to: 2, typ: Promise, seq: 3, preSeq: 2, val: "m3"}
	//nt.sendTo(m3)
	//time.Sleep(time.Second)
}

func TestFollowerToCandidate(t *testing.T) {
	nt := CreateNetwork(1, 2, 3, 4)
	var serverList []server
	nServer1 := NewServer(1, Follower, nt.getNodeNetwork(1), 2, 3, 4)
	serverList = append(serverList, *nServer1)
	nServer2 := NewServer(2, Follower, nt.getNodeNetwork(2), 1, 3, 4)
	serverList = append(serverList, *nServer2)
	nServer3 := NewServer(3, Follower, nt.getNodeNetwork(3), 2, 1, 4)
	serverList = append(serverList, *nServer3)
	nServer4 := NewServer(4, Follower, nt.getNodeNetwork(4), 2, 3, 1)
	serverList = append(serverList, *nServer4)

	for _, sev := range serverList {
		go sev.RunServerLoop()
	}

	//Set server1 an action.
	nServer1.AssignAction(datalog{term: 1, action: "x<-1"})

	//Wait server1 become candidate.
	for {
		if nServer1.Whoareyou() == Candidate {
			break
		}
		time.Sleep(time.Second)
	}

}
