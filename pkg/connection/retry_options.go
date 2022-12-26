package connection

import "time"

type Policy struct {
	Timeout        time.Duration
	RetryCount     int           // The maximum backoff number of retries to attempt.
	RetryMaxDelay  time.Duration // The max backoff delay between retries.
	RetryBaseDelay time.Duration // The base delay for calculation backoff between retries.
}

func DefaultPolicy() Policy {
	return Policy{
		Timeout:        10 * time.Second,
		RetryCount:     5,
		RetryMaxDelay:  5 * time.Second,
		RetryBaseDelay: 50 * time.Millisecond,
	}
}
