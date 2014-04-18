package main

import (
	"time"
)

const startTimeFormat = "2006-01-02 15:04"
const sendTimeFormat = "1504"

type Flag byte

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
}

type Reminder struct {
	Name  string
	Email string
}

// fmt.Println(Saturday & set)

// m, _ := strconv.ParseInt("0111110", 2, 64)
// fmt.Println(m & 64)
// fmt.Println(Flag(m) & Tuesday)
// fmt.Println(Saturday)
