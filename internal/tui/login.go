package tui

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"keeper/internal/service/secret"
	"keeper/internal/service/user"
)

const (
	loginLogin = iota
	loginPassword
)

type LoginModel struct {
	parent  tea.Model
	us      *user.Service
	ss      *secret.Service
	inputs  []textinput.Model
	focused int
	err     error
}

func NewLoginModel(parent tea.Model, us *user.Service, ss *secret.Service) LoginModel {
	inputs := make([]textinput.Model, 2)

	inputs[loginLogin] = NewInput(inputText, true)
	inputs[loginPassword] = NewInput(inputPassword, false)

	return LoginModel{
		parent:  parent,
		us:      us,
		ss:      ss,
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				login := m.inputs[loginLogin].Value()
				password := m.inputs[loginPassword].Value()

				auth, err := m.us.Login(context.Background(), login, password)
				if err != nil {
					m.err = err
					return m, nil
				}

				return NewMenuModel(auth, m.us, m.ss), nil
			}
			m.nextInput()
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m.parent, nil
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}

		m.err = nil
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m LoginModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		"%s\n\n%s:\n%s\n\n%s:\n%s\n%s\n%s",
		titleStyle.Render("Authorization"),
		inputLabelStyle.Render("Login"),
		m.inputs[loginLogin].View(),
		inputLabelStyle.Render("Password"),
		m.inputs[loginPassword].View(),
		errToString(m.err),
		continueStyle.Render("enter: sign in, esc: back"),
	))
}

func (m *LoginModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *LoginModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
