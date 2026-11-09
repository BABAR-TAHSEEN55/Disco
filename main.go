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

	go func() {

		if err := producer.LoadRecipients("emails.csv", emailHolder); err != nil {
			log.Fatal(err)
		}
	}()
	wg.Add(1)
	go consumer(1, emailHolder, &wg)
	wg.Wait()

}
