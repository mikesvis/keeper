package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"keeper/internal/app"
)

type RegisterModel struct {
	app            *app.App
	form           *huh.Form
	password       string
	passwordRepeat string
	name           string
	email          string
}

func NewRegisterModel(app *app.App) tea.Model {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Username").Key("username"),
			huh.NewInput().Title("Password").EchoMode(huh.EchoModePassword),
		),
	)

	m := RegisterModel{app: app, form: form}
	return m
}

func (m RegisterModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m RegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	//if m.form != nil {
	//	f, cmd := m.form.Update(msg)
	//	m.form = f.(*huh.Form)
	//	cmds = append(cmds, cmd)
	//}

	if m.form.State == huh.StateAborted {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}

func (m RegisterModel) View() string {
	return m.form.View()
}
