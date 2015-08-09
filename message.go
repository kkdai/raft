package raft

type msgType int

const (
	Heartbit msgType = iota + 1
	HeartbitFeedback
	RequestVote
	AcceptVote
	WinningVote
)

type message struct {
	from int
	to   int
	typ  msgType
	term int
	val  datalog
}

func (m *message) getMsgTerm() int {
	return m.term
}

func (m *message) getVal() datalog {
	return m.val
}
