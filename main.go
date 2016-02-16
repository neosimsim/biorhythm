package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type outputError struct {
	message string
}

func (err *outputError) Error() string {
	return err.message
}

func main() {
	birthday, err := time.Parse("2006-01-02", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	daysAlive := int(time.Since(birthday).Hours() / 24)
	fmt.Printf("Sie haben bis jetzt %v Tage gelebt.\n", daysAlive)
	p, e, i := biorythm(daysAlive)
	plotValues([]string{"p", "e", "i"}, p, e, i)
}
