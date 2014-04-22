package schedule

import (
	"fmt"
	"time"
)

type Flag byte

// fmt.Println(Saturday & set)

// m, _ := strconv.ParseInt("0111110", 2, 64)
// fmt.Println(m & 64)
// fmt.Println(Flag(m) & Tuesday)
// fmt.Println(Saturday)

const startTimeFormat = "2006-01-02 15:04"
const sendTimeFormat = "1504"

const (
	Sunday Flag = 1 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// "0900,1700|0111110|2014-04-18 00:00|1"
// [time_to_send]|[bitmask_of_week]|[start_datetime]|[repeat]
type scheduleFormat struct {
	sendTimes []time.Time
	days      []Flag
	startTime time.Time
	repeat    bool
}

type Scheduler struct {
	jobs          []Job
	stats         map[string]int
	checkInterval time.Duration
}

func NewScheduler(seconds int) *Scheduler {
	fmt.Println(seconds)
	return &Scheduler{
		jobs:          make([]Job, 0),
		stats:         make(map[string]int),
		checkInterval: time.Duration(seconds) * time.Second,
	}
}

func (s *Scheduler) Run() error {
	fmt.Println("Starting Run", cap(s.jobs))
	tick := time.NewTicker(s.checkInterval)
	for now := range tick.C {
		fmt.Println("Running at", now, cap(s.jobs))
		for _, j := range s.jobs {
			fmt.Println("Job ready")
			j.Work()
		}
		fmt.Println("Waiting...")
	}
	return nil
}

func (s *Scheduler) AddJob(j Job) (int, error) {
	s.jobs = append(s.jobs, j)
	return cap(s.jobs), nil
}

type Job interface {
	Work() (bool, error)
}

type Reminder interface {
	GetName() string
	GetEmail() string
	GetContent() string
}

type BasicReminder struct {
	Name    string
	Email   string
	Content string
}

func (b *BasicReminder) GetName() string {
	return b.Name
}

func (b *BasicReminder) GetEmail() string {
	return b.Email
}

func (b *BasicReminder) GetContent() string {
	return b.Content
}

func (b *BasicReminder) Work() (bool, error) {
	fmt.Printf("Sending %s to %s\n", b.GetName(), b.GetEmail())
	fmt.Println(b.GetContent())
	return true, nil
}
