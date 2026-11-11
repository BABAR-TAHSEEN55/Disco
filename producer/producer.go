package producer

import (
	"encoding/csv"
	"os"
)

type Recipient struct {
	Name  string
	Email string
}

func LoadRecipients(filename string, emailHolder chan Recipient) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	reader := csv.NewReader(f)
	rec, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, r := range rec[1:] {
		// fmt.Println(r)
		emailHolder <- Recipient{r[0], r[1]}
	}

	defer f.Close()
	defer close(emailHolder)

	return nil
}
