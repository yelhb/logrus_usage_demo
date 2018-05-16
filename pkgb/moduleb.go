package pkgb

import (
	"fmt"
	"sync"
	"time"

	log "github.com/yelhb/logrus_usage_demo/logger"
)

func TestFunc(wg *sync.WaitGroup) {

	cnt := 0
	for {
		cnt++
		log.Info(fmt.Sprintf("Test func running... in pkgb, cnt: %d", cnt))
		time.Sleep(2 * time.Second)
		if cnt >= 5 {
			break
		}
	}

	wg.Done()
}
