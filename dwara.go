package main

import (
	"github.com/dminGod/Dwara/config"
	"github.com/dminGod/Dwara/servers/http"
	"os"
	"github.com/sirupsen/logrus"
	"fmt"
	"time"
)

var log = logrus.New()


func main() {

	config.LoadInitialConfig()

	Conf := config.Conf.Get()

	log.Out = os.Stdout

	d := time.Now()

	logFile := fmt.Sprintf("%v/%v-%v", Conf.DwaraConfig.LogFolder, Conf.DwaraConfig.LogFilePrefix, d.Format("06-01-02_15_04.log") )

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0666)

	if err == nil {

		log.Out = file
	} else {

		log.Info("Failed to log to file, using default stderr")
	}

	log.Info("This is a log message")

	http_request_server.InitServer()
}