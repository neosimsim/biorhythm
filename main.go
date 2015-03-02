package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type outputError struct {
	message string
}

func (err *outputError) Error() string {
	return err.message
}

func terminalDim() (height, width int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	out = out[:len(out)-1] // slice "\n" at the end
	return parseStringToIntTouple(string(out))
}

func parseStringToIntTouple(str string) (x, y int, err error) {
	xy := strings.Split(str, " ")
	if len(xy) != 2 {
		return 0, 0, &outputError{message: "Unexpected output from 'stty size'"}
	}
	x, err = strconv.Atoi(xy[0])
	if err != nil {
		return 0, 0, err
	}
	y, err = strconv.Atoi(xy[1])
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

func main() {
	birthday, err := time.Parse("2006-01-02", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	_, width, err := terminalDim()
	if err != nil {
		log.Fatal(err)
	}

	daysAlive := int(time.Since(birthday).Hours() / 24)
	fmt.Printf("Sie haben bis jetzt %v Tage gelebt.\n", daysAlive)
	plots := biorythm(width, daysAlive)
	printPlot(width, plots[0], plots[1], plots[2])
}
