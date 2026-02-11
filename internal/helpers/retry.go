package helpers

import (
	"log"
	"time"
)

func Retry(fn func(int) error, retriesWait time.Duration) {
	for currIter := 1; ; currIter++ {
		if err := fn(currIter); err != nil {
			log.Println(err)
			time.Sleep(retriesWait)
			continue
		}
		return
	}
}
