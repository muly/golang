// simpler exponential backoff delay calculation scalled to the limits we provide
package main

import (
	"fmt"
	"math"
	"time"
)

// ExponentialBackoff calculates the exponential backoff delay.
// It calculates the exponential growth from initialDelay to maxDelay (y axis) to the maxRetryCount (x axis)
// to acheive that window, the x (x axis value) and b (base) values in the exponential growth formula are calculated as below.
// 	y = a * b^x
// 	a: initial delay
// 	b (scaled): Max Delay / initial delay
// 	x (scaled): current retry count / max retry count
func ExponentialBackoff(initialDelay time.Duration, maxDelay time.Duration, maxRetryCount int,  currentRetryCnt int) time.Duration {
	a := initialDelay.Seconds()
	scaledB := maxDelay.Seconds() / initialDelay.Seconds()
	scaledX := float64(currentRetryCnt) / float64(maxRetryCount)

	y := a * math.Pow(scaledB, scaledX)
	return time.Duration(math.Round(y)) * time.Second
}

func main() {
	initialDelay := 30 * time.Second
	MaxExpDelay := 180 * time.Second 
	maxRetryCount := 10             

	for i := 0; i <= maxRetryCount; i++ {
		delay := ExponentialBackoff(initialDelay, MaxExpDelay, maxRetryCount, i)
		fmt.Printf("Retry %d: Delay %v\n", i, delay)
		// time.Sleep(delay)
	}
}
