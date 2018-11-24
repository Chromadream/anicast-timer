package utility

import (
	"time"
)

//A Timer structure, indicates the start of the recording time
type Timer struct {
	StartTime time.Time
}

//StartTiming starts the timing function, and returns a Timer instance
func StartTiming() (timer Timer) {
	timer = Timer{time.Now()}
	return
}

//GetDuration returns the length of current recording session
func (x Timer) GetDuration() string {
	duration := time.Since(x.StartTime)
	return duration.Round(2 * time.Second).String()
}
