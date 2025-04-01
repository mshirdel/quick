package app

import (
	"fmt"

	"github.com/mshirdel/quick/config"
)

type Application struct {
	configPath string
	Cfg        *config.Config
}

func New(configPath string) *Application {
	return &Application{
		configPath: configPath,
	}
}

func (a *Application) InitAll() error {
	if err := a.InitConfig(); err != nil {
		return err
	}

	return nil
}

func (a *Application) InitConfig() error {
	var err error
	if a.Cfg != nil {
		return nil
	}

	a.Cfg, err = config.InitViper(a.configPath)
	if err != nil {
		return fmt.Errorf("config init: %w", err)
	}

	return nil
}

func (a *Application) Shutdown() {
}
