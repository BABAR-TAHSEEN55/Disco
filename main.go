package main

import (
	"log"
	"sync"

	"github.com/BABAR-TAHSEEN55/disco/producer"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	var wg sync.WaitGroup
	emailHolder := make(chan producer.Recipient)
	workerCount := 5
	for i := range workerCount {

		wg.Add(1)
		go consumer(i, emailHolder, &wg)

	}

	go func() {
		if err := producer.LoadRecipients("emails.csv", emailHolder); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
