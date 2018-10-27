package main

import (
	"github.com/liuliqiang/laudocker/cmd"
	"github.com/liuliqiang/laudocker/internal/container"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var enableTty = false

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("docker run")

		if len(args) < 1 {
			logrus.Fatalf("Missing container command.")
		}

		logrus.Debugf("command: %s", args[0])
		container.Run(enableTty, args[0])
	},
}

func init() {
	runCmd.Flags().BoolVarP(&enableTty, "tty", "t", enableTty, "Enable tty")

	cmd.RootCmd.AddCommand(runCmd)
}
