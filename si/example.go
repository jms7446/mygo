package si

import "time"

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

// Deadline :
type Deadline struct {
	time.Time
}

// NewDeadline :
func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// OverDue :
func (d *Deadline) OverDue() bool {
	return d != nil && (*d).Before(time.Now())
}

// Task :
type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

// OverDue :
func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}
