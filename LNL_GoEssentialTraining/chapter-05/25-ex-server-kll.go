package main

import (
	"github.com/pkg/errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Can't open pid file")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("can't remove pidfile - %s", err)
	}

	strPID := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPID)

	if err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	fmt.Printf("Kill process PID: %v\n", pid)
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}