package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"keeper/internal/app"
)

type WelcomeModel struct {
	app  *app.App
	auth bool
}

func NewWelcomeModel(sess ssh.Session, app *app.App) tea.Model {
	return WelcomeModel{
		app:  app,
		auth: app.UserService.IsUserAuthed(sess.Context()),
	}
}

func (m WelcomeModel) Init() tea.Cmd {
	return nil
}

func (m WelcomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return NewRegisterModel(m.app), nil
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m WelcomeModel) View() string {
	var s string

	if m.auth {
		s += "Welcome to the keeper, <USERNAME>!\n\n"
		s += "Press enter to proceed, press q or ctrl+c to quit."
		return s
	}

	s += "Welcome to the keeper!\n\nYou must be registered to use this password keeper!\n\n"
	s += "Press enter to proceed to registration, press q / ctrl+c / esc to quit."
	return s
}
