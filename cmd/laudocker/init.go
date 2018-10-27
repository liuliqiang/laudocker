package main

import (
	"github.com/liuliqiang/laudocker/cmd"
	"github.com/liuliqiang/laudocker/internal/container"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("docker init")

		if len(args) != 1 {
			logrus.Fatalf("Invalid args count: %d", len(args))
		}
		logrus.Debugf("init cmd: %s", args[0])

		container.RunContainerInitProcess(args[0], nil)
	},
}

func init() {
	cmd.RootCmd.AddCommand(initCmd)
}
