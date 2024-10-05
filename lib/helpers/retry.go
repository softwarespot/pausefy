package helpers

import (
	"log"
	"time"
)

func Retry(fn func(int) error, retriesWait time.Duration) {
	for currIter := 1; ; currIter++ {
		err := fn(currIter)
		if err == nil {
			return
		}

		log.Println(err)
		time.Sleep(retriesWait)
	}
}
