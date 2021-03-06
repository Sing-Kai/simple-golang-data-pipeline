package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

type externalData struct {
	inputData
	// fields fetched from external service
	relatedIds []string
}

func fetchData(ic <-chan inputData) <-chan externalData {
	oc := make(chan externalData)
	go func() {
		wg := &sync.WaitGroup{}

		for input := range ic {
			wg.Add(1)
			go fetchExternalService(input, oc, wg)
		}

		wg.Wait()
		close(oc)
	}()

	return oc
}

//mimics http request
func fetchExternalService(input inputData, output chan<- externalData, wg *sync.WaitGroup) {

	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	s := make([]string, 0)

	for i := 0; i < rand.Intn(10); i++ {
		s = append(s, uuid.New().String())
	}

	output <- externalData{input, s}

	wg.Done()
}
