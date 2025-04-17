package cmd

import (
	"github.com/mshirdel/quick/app"
	"github.com/mshirdel/quick/app/http"
	"github.com/spf13/cobra"
)

var _serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "Serve API server",
	RunE: func(cmd *cobra.Command, args []string) error {
		app := app.New(_configPath)
		if err := app.InitAll(); err != nil {
			return err
		}

		defer app.Shutdown()

		// init http server
		server := http.NewHTTPServer(app)
		defer server.Shutdown()

		go server.Start()

		ctx, cancel := handleInterrupts()
		defer cancel()

		<-ctx.Done()

		return nil
	},
}
