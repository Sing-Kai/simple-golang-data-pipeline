package main

import (
	"bufio"
	"log"
	"os"

	"github.com/google/uuid"
)

///read from a file with guids
func generateData() <-chan uuid.UUID {
	c := make(chan uuid.UUID)
	const filePath = "guids.txt"

	go func() {

		file, _ := os.Open(filePath)
		defer file.Close()

		s := bufio.NewScanner(file)

		for s.Scan() {
			line := s.Text()
			guid, err := uuid.Parse(line)

			if err != nil {
				log.Printf(err.Error())
				continue
			}

			c <- guid
		}
		close(c)

	}()

	return c
}
