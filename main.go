package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	log "github.com/yelhb/logrus_usage_demo/logger"
	"github.com/yelhb/logrus_usage_demo/pkga"
	"github.com/yelhb/logrus_usage_demo/pkgb"
)

func main() {
	logFile, err := os.OpenFile("log/run.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("failed to open log file")
		os.Exit(-1)
	}

	log.SetOutput(logFile)
	log.SetFormatter(new(logrus.JSONFormatter))
	log.SetLevel(logrus.DebugLevel)

	log.Info("System start ...")
	wg := sync.WaitGroup{}
	wg.Add(2)

	go pkga.TestFunc(&wg)
	go pkgb.TestFunc(&wg)

	wg.Wait()

	log.Info("System end ...")
}
