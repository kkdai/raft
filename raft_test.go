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
}

func TestFollowerElectionToLeader(t *testing.T) {
	nt := CreateNetwork(1, 2, 3, 4)
	nServer1 := NewServer(1, Follower, nt.getNodeNetwork(1), 2, 3, 4)
	nServer2 := NewServer(2, Follower, nt.getNodeNetwork(2), 1, 3, 4)
	nServer3 := NewServer(3, Follower, nt.getNodeNetwork(3), 2, 1, 4)
	nServer4 := NewServer(4, Follower, nt.getNodeNetwork(4), 2, 3, 1)

	go nServer1.RunServerLoop()
	go nServer2.RunServerLoop()
	go nServer3.RunServerLoop()
	go nServer4.RunServerLoop()

	//Set server1 an action.
	nServer1.AssignAction(datalog{term: 1, action: "x<-1"})
	log.Println("Assign value to server 1 ")

	//Wait server1 become Leader.
	for {
		if nServer1.Whoareyou() == Leader {
			break
		}
		time.Sleep(time.Second)
	}
}
