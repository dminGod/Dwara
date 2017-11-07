package http_request_server

import (
	"github.com/dminGod/Dwara/config"
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func InitServer() {

	Conf := config.Conf.Get()

	log.Info("This is from the initi server of http.go")

	if Conf.HttpServer.Enabled {

		l := fmt.Sprintf("%v:%v", Conf.HttpServer.ListenAddress, Conf.HttpServer.Port)
		http.HandleFunc("/", handler)

		err := http.ListenAndServe(l, nil)

		fmt.Printf("This is the server : %v", l)

		if err != nil {

			log.Errorf("Error in the string : %v", err)
		}
	} else {

		fmt.Printf("Its enabled...")
		log.Info("This http server is disabled")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}




