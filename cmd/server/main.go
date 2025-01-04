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
	"keeper/internal/middleware"
	"keeper/internal/postgres"
	uService "keeper/internal/service/user"
	uStorage "keeper/internal/storage/user"
	"keeper/internal/tui"
	"os"
)

func main() {
	// Config
	appConfig, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error while initializing server config: %s\n", err.Error())
		os.Exit(1)
	}

	// Logger
	appLogger, err := logger.NewLogger(appConfig)
	if err != nil {
		fmt.Printf("Error while initializing server logger: %s\n", err.Error())
		os.Exit(1)
	}

	appLogger.Infof("Initialized config: %+v", appConfig)

	// Storage
	db, err := postgres.NewPostgres(appConfig)
	if err != nil {
		appLogger.Fatalf("Error while initializing db: %s\n", err.Error())
		os.Exit(1)
	}

	// User Storage & Service
	userStorage, err := uStorage.NewStorage(db, appLogger)
	if err != nil {
		appLogger.Fatalf("Error while initializing user storage: %s\n", err.Error())
		os.Exit(1)
	}
	userService := uService.NewService(userStorage, appLogger)

	keeperApp := app.NewApp(userService)

	srv, err := wish.NewServer(
		wish.WithAddress(appConfig.ServerAddress),
		wish.WithHostKeyPath(`/home/mikhail/Learning/Go/certs/server`),
		wish.WithPublicKeyAuth(middleware.SignIn(keeperApp, appLogger)),
		wish.WithMiddleware(
			bubbletea.Middleware(myHandler(keeperApp)),
			activeterm.Middleware(),
		),
	)

	if err != nil {
		appLogger.Error("Could not start server", "error", err)
	}

	appLogger.Info("Starting SSH server ", appConfig.ServerAddress)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		// We ignore ErrServerClosed because it is expected.
		log.Error("Could not start server", "error", err)
	}
}

func myHandler(ka *app.App) bubbletea.Handler {
	return func(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
		_, _, active := sess.Pty()
		if !active {
			wish.Fatalln(sess, "no active terminal, skipping")
			return nil, nil
		}

		m := tui.NewWelcomeModel(sess, ka)
		return m, []tea.ProgramOption{tea.WithAltScreen()}
	}
}

//func myHandler(ka *app.App) func(next ssh.Handler) ssh.Handler {
//	return func(next ssh.Handler) ssh.Handler {
//		m := tui.NewWelcomeModel(ka, s)
//		return m, []tea.ProgramOption{tea.WithAltScreen()}
//	}
//}

//func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
//	m := tui.NewWelcomeModel(s)
//	return m, []tea.ProgramOption{tea.WithAltScreen()}
//}

//func myMiddleware(ka *app.App) wish.Middleware {
//	newProg := func(m tea.Model, opts ...tea.ProgramOption) *tea.Program {
//		p := tea.NewProgram(m, opts...)
//		return p
//	}
//	teaHandler := func(s ssh.Session) *tea.Program {
//		pty, _, active := s.Pty()
//		if !active {
//			wish.Fatalln(s, "no active terminal, skipping")
//			return nil
//		}
//		m := tui.NewWelcomeModel(s, pty, ka)
//		return newProg(m, append(bubbletea.MakeOptions(s), tea.WithAltScreen())...)
//	}
//	return bubbletea.MiddlewareWithProgramHandler(teaHandler, termenv.ANSI256)
//}
