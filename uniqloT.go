package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
)

type ControlMessage struct {
	Target string
	Count  int
}

func main() {
	controlChannel := make(chan ControlMessage)
	workerCompleteChan := make(chan bool)
	statusPollChannel := make(chan chan bool)
	workerActive := false

	go admin(controlChannel, statusPollChannel)

	for {
		select {
		case respChan := <-statusPollChannel:
			respChan <- workerActive
			workerActive = true
		case status := <-workerCompleteChan:
			workerActive = status
		}
	}
}

func admin(cc chan ControlMessage, statusPollChannel chan chan bool) {
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		count, err := strconv.ParseInt(r.FormValue("count"), 10, 32)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		msg := ControlMessage{Target: r.FormValue("target"), Count: int(count)}
		cc <- msg
		fmt.Fprintf(w, "Control message issued for Target: %s, Count: %d", html.EscapeString(r.FormValue("target")), count)
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		reqChan := make(chan bool)
		statusPollChannel <- reqChan
		select {
		case result := <-reqChan:
			if result {
				fmt.Fprint(w, "ACTIVE")
			} else {
				fmt.Fprint(w, "INACTIVE")
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}
