package cmd

import (
	"github.com/mshirdel/quick/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var _serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "Serve API server",
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("test is ok")
		app := app.New(_configPath)
		if err := app.InitAll(); err != nil {
			return err
		}

		defer app.Shutdown()

		// init http server
		return nil
	},
}
