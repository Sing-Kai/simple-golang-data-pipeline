package main

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type inputData struct {
	id        string
	timeStamp int64
}

//converts incoming data into a struct and returns a chan
func prepareData(ic <-chan uuid.UUID) <-chan inputData {
	oc := make(chan inputData)
	go func() {
		for id := range ic {
			input := inputData{id: id.String(), timeStamp: time.Now().UnixNano()}
			log.Printf("data has been processed:  %+v", input)
			oc <- input
		}
		close(oc)
	}()

	return oc
}
