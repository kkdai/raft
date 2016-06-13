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
	log.Println("--------TestFollowerElectionToLeader-------")
	nt := CreateNetwork(1, 2, 3, 4, 5)
	nServer1 := NewServer(1, Follower, nt.getNodeNetwork(1), 2, 3, 4, 5)
	nServer2 := NewServer(2, Follower, nt.getNodeNetwork(2), 1, 3, 4, 5)
	nServer3 := NewServer(3, Follower, nt.getNodeNetwork(3), 2, 1, 4, 5)
	nServer4 := NewServer(4, Follower, nt.getNodeNetwork(4), 2, 3, 1, 5)
	nServer5 := NewServer(5, Follower, nt.getNodeNetwork(5), 2, 3, 1, 5)

	log.Println(nServer1.id, nServer2.id, nServer3.id, nServer4.id, nServer5.id)
	// go nServer1.runServerLoop()
	// go nServer2.runServerLoop()
	// go nServer3.runServerLoop()
	// go nServer4.runServerLoop()
	// go nServer5.runServerLoop()

	//Set server1 an action.
	nServer1.AppendEntries(datalog{term: 1, action: "x<-1"})
	log.Println("Assign value to server 1 ")

	//Wait server1 become Leader.
	for i := 0; i <= 10; i++ {
		if nServer1.Whoareyou() == Leader {
			log.Println("1 become leader done.")
			return
		}
		log.Println("1 still not leader:", nServer1.Whoareyou())
		time.Sleep(time.Second)
	}
	t.Error("No one become leader on basic")
}
