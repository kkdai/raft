package raft

type msgType int

const (
	Heartbit msgType = iota + 1
	HeartbitFeedback
	RequestVote
	AcceptVote
	WinningVote
)

//From has different meaning in different state:
//[AppendEntries]: from = leaderID
//[RequestVote]: from = candidateID

type Message struct {
	from int
	to   int
	typ  msgType
	term int
	val  datalog

	lastLogIndex int
	lastLogTerm  int
	leaderCommit int
	success      bool //return false if any error on specific message handle (ex: RequestVote, AppendEntries)
}

func (m *Message) GetMsgTerm() int {
	return m.term
}

func (m *Message) GetVal() datalog {
	return m.val
}
