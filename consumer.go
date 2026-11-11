package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"

	"github.com/BABAR-TAHSEEN55/disco/producer"
)

func consumer(id int, emailHolder chan producer.Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recpient := range emailHolder {
		// fmt.Println(id, recpient)
		smptHost := "localhost"
		smptPort := "1025"
		formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Hi! %s\r\n\r\n%s", recpient.Email, recpient.Name, "Just Testing our email dispatcher")

		msg := []byte(formattedMsg)
		err := smtp.SendMail(smptHost+":"+smptPort, nil, "babarsmoke24@gmail.com", []string{recpient.Email}, msg)
		if err != nil {
			fmt.Println("Error sending email:", err)
		} else {
			fmt.Println("Email sent successfully to", recpient.Email)
		}
		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: Sent email to %s\n", id, recpient.Email)
	}
}
