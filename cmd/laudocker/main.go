package main

import (
	"os"

	"github.com/liuliqiang/laudocker/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	cmd.Execute()
}

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}