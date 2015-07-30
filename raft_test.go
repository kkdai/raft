package raft

import (
	"fmt"
	"log"
	"math/rand"
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

func TestSingleProser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5))

}
