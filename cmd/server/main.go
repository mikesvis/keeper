package main

import (
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"keeper/internal/app"
	"keeper/internal/config"
	"keeper/internal/logger"
	"keeper/internal/tui"
	"os"
)

func main() {
	// Config
	appConfig, err := config.NewConfig()
	if err != nil {
		fmt.Printf("error while initializing server config: %s\n", err.Error())
		os.Exit(1)
	}

	// Logger
	appLogger, err := logger.NewLogger(appConfig)
	if err != nil {
		fmt.Printf("error while initializing server logger: %s\n", err.Error())
		os.Exit(1)
	}

	appLogger.Infof("initialized logger and config: %+v", appConfig)

	keeper, err := app.NewApp(appConfig, appLogger)
	if err != nil {
		fmt.Printf("error while initializing keeper app: %s\n", err.Error())
		os.Exit(1)
	}

	srv, err := wish.NewServer(
		wish.WithAddress(appConfig.ServerAddress),
		wish.WithHostKeyPath(appConfig.ServerCertPath),
		wish.WithMiddleware(
			bubbletea.Middleware(myHandler(keeper)),
			activeterm.Middleware(),
		),
	)

	if err != nil {
		appLogger.Error("Could not start server", "error", err)
	}

	appLogger.Info("Starting SSH server ", appConfig.ServerAddress)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not start server", "error", err)
	}
}

func myHandler(keeper *app.App) bubbletea.Handler {
	return func(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
		_, _, active := sess.Pty()
		if !active {
			wish.Fatalln(sess, "no active terminal, skipping")
			return nil, nil
		}

		m := tui.NewStart(keeper.UserService, keeper.SecretService)
		return m, []tea.ProgramOption{tea.WithAltScreen()}
	}
}
