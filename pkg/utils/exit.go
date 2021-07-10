package utils

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func Exit() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)
	go func() {
		<-sigs
		done <- true
	}()
	logrus.Info("Server Start Awaiting Signal")
	<-done
	logrus.Info("Exiting")
}
