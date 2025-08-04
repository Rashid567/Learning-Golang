package l_2_08

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var now = time.Now()

type MockNTPClient struct{}

func (m *MockNTPClient) Time(server string) (time.Time, error) {
	if server == "error.server" {
		return time.Time{}, fmt.Errorf("NTP server unavailable")
	}
	return now, nil
}

func runPrintCurrentTime(server *string) (string, string, int) {
	originalStdout := stdOut
	originalStderr := stdErr
	originalNTPTime := ntpTime

	defer func() {
		stdOut = originalStdout
		stdErr = originalStderr
		ntpTime = originalNTPTime
	}()

	// Подменяем stdout и stderr
	r1, w1, _ := os.Pipe()
	r2, w2, _ := os.Pipe()
	stdOut = w1
	stdErr = w2

	// Подменяем NTP клиент
	ntpTime = (&MockNTPClient{}).Time

	exitCode := PrintCurrentTime(server)

	err := w1.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = w2.Close()
	if err != nil {
		log.Fatal(err)
	}

	var bufOut bytes.Buffer
	_, err = bufOut.ReadFrom(r1)
	if err != nil {
		log.Fatal(err)
	}

	var bufErr bytes.Buffer
	_, err = bufErr.ReadFrom(r2)
	if err != nil {
		log.Fatal(err)
	}

	return bufOut.String(), bufErr.String(), exitCode

}

func TestPrintCurrentTime_Success(t *testing.T) {
	out, err, exitCode := runPrintCurrentTime(nil)

	expectedOut := fmt.Sprintf("Current time: %s\n", now)
	assert.Equal(t, expectedOut, out)

	expectedErr := ""
	assert.Equal(t, expectedErr, err)

	assert.Equal(t, exitCode, 0)
}

func TestPrintCurrentTime_Error(t *testing.T) {
	server := "error.server"
	out, err, exitCode := runPrintCurrentTime(&server)

	expectedOut := ""
	assert.Equal(t, expectedOut, out)

	expectedErr := "Failed to get current time: NTP server unavailable\n"
	assert.Equal(t, expectedErr, err)

	assert.Equal(t, exitCode, 1)
}
