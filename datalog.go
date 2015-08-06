package raft

type datalog struct {
	term   int
	action string
}

func (d *datalog) identical(data datalog) bool {
	return d.term == data.term && d.action == data.action
}

type submittedItems struct {
	logs     []datalog
	logIndex int
}

func (d *submittedItems) getLatestLogs() *datalog {
	return &(d.logs[len(d.logs)-1])
}

func (d *submittedItems) identicalWith(b *submittedItems) bool {
	if d.logIndex == b.logIndex && d.getLatestLogs().term == b.getLatestLogs().term {
		return true
	}

	return false
}

func (s *submittedItems) add(data datalog) {
	s.logIndex++
	s.logs = append(s.logs, data)
}
