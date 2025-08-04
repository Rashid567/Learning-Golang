package l_2_08

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

// Создаём переменные, что бы можно было их подменять в тестах
var ntpTime = ntp.Time
var stdOut = os.Stdout
var stdErr = os.Stderr

func PrintCurrentTime(server *string) int {
	defaultServer := "0.beevik-ntp.pool.ntp.org"
	if server == nil {
		server = &defaultServer
	}

	currentTime, err := ntpTime(*server)
	if err != nil {
		_, err := fmt.Fprintf(stdErr, "Failed to get current time: %s\n", err.Error())
		if err != nil {
			log.Fatal(err)
		}
		return 1
	}
	_, err = fmt.Fprintf(stdOut, "Current time: %s\n", currentTime)
	if err != nil {
		log.Fatal(err)
	}
	return 0
}
