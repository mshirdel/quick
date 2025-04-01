package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	_rootCMD = cobra.Command{
		Use:   "quick",
		Short: "quick",
	}

	_configPath string
)

func init() {
	_rootCMD.PersistentFlags().StringVarP(&_configPath, "config", "c", "config.yml", "config path (directory or file)")
	_rootCMD.AddCommand(_serveCMD)
}

func Execute() {
	if err := _rootCMD.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
