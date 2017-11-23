// Package wait provides helper functions that make it easy to wait for common
// events or times.
package wait

import "time"

// UntilNextMinute waits until the minute after the provided startTime. If nil
// is provided as the start time, then it will use the current time.
func UntilNextMinute(startTime *time.Time) {
	// Set startTime to now if nil is provided.
	if startTime == nil {
		now := time.Now()
		startTime = &now
	}

	// Calculate the next minute.
	targetTime := time.Date(
		startTime.Year(),
		startTime.Month(),
		startTime.Day(),
		startTime.Hour(),
		startTime.Minute(),
		0,
		0,
		startTime.Location(),
	).Add(time.Minute * 1)

	// Create a time that will fire on the next minute.
	t := time.NewTimer(time.Until(targetTime))

	// Wait for the timer to fire.
	<-t.C
}
