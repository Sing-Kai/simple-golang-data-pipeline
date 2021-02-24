package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {

	guid := randomGuidGenerator()

	for id := range guid {
		fmt.Println(id.String())
	}

	c := saveData(fetchData(
		prepareData(
			generateData(),
		),
	))

	for data := range c {
		log.Printf("Items saved: %+v", data)
	}

}

func randomGuidGenerator() <-chan uuid.UUID {

	c := make(chan uuid.UUID)

	go func() {

		c <- uuid.New()

		close(c)

	}()

	return c
}
