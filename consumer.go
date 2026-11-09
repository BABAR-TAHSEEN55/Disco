package main

import (
	"fmt"
	"sync"

	"github.com/BABAR-TAHSEEN55/disco/producer"
)

func consumer(id int, emailHolder chan producer.Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recpient := range emailHolder {
		fmt.Println(recpient, id)
	}
}
