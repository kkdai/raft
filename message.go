package raft

type msgType int

const (
	Heartbit msgType = iota + 1
	RequestVote
	AcceptVote
	WinningVote
)

type message struct {
	from int
	to   int
	typ  msgType
	term int
	val  string
}

func (m *message) getMsgTerm() int {
	return m.term
}

func (m *message) getVal() string {
	return m.val
}
