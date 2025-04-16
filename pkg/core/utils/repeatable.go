package utils

import "time"

func DoWithAttempts(fn func() error, maxAttempts int32, delay time.Duration) (err error) {
	t := time.NewTimer(delay)
	for maxAttempts > 0 {
		if err = fn(); err != nil {
			t.Reset(delay)
			<-t.C

			maxAttempts--

			continue
		}

		return nil
	}

	return err
}
