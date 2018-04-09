package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/karthequian/demos/golang/async/pkg/dispatcher"
)

var (
	MaxWorker       = 3  //os.Getenv("MAX_WORKERS")
	MaxQueue        = 20 //os.Getenv("MAX_QUEUE")
	MaxLength int64 = 2048
)

func payloadHandler(w http.ResponseWriter, r *http.Request) {

	payload := dispatcher.Payload{Name: "Something"}

	work := dispatcher.Job{Payload: payload}
	fmt.Println("sending payload  to workque")
	// Push the work onto the queue.
	dispatcher.JobQueue <- work
	fmt.Println("sent payload  to workque")

	w.WriteHeader(http.StatusOK)
}
func main() {
	dispatcher.JobQueue = make(chan dispatcher.Job, MaxQueue)
	log.Println("main start")
	dispatcher := dispatcher.NewDispatcher(MaxWorker)
	dispatcher.Run()
	http.HandleFunc("/", payloadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("starting listening for payload messages")
	} else {
		fmt.Errorf("an error occured while starting payload server %s", err.Error())
	}

	time.Sleep(time.Hour)
}
