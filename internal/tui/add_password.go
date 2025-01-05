package tui

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"keeper/internal/domain"
	"keeper/internal/service/secret"
)

type AddPasswordModel struct {
	parent  MenuModel
	ss      *secret.Service
	inputs  []textinput.Model
	focused int
	err     error
}

const (
	secretPasswordName  = iota
	secretPasswordLogin = iota
	secretPasswordPassword
)

func NewAddPasswordModel(parent MenuModel, ss *secret.Service) AddPasswordModel {
	inputs := make([]textinput.Model, 3)

	inputs[secretPasswordName] = NewInput(inputText, true)
	inputs[secretPasswordLogin] = NewInput(inputText, false)
	inputs[secretPasswordPassword] = NewInput(inputPassword, false)

	return AddPasswordModel{
		parent:  parent,
		ss:      ss,
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m AddPasswordModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddPasswordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				name := m.inputs[secretPasswordName].Value()
				login := m.inputs[secretPasswordLogin].Value()
				password := m.inputs[secretPasswordPassword].Value()

				rawData, err := domain.NewUserSecretPassword(login, password).GetData()
				if err != nil {
					m.err = err
					return m, nil
				}

				data, err := domain.MakeUserSecretData(domain.UserSecretPasswordType, rawData)
				if err != nil {
					m.err = err
					return m, nil
				}

				err = m.ss.Create(context.Background(), m.parent.GetUserID(), domain.UserSecretPasswordType, name, &data)
				if err != nil {
					m.err = err
					return m, nil
				}

				return NewSuccessModel(m.parent), nil
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}

		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		if m.inputs[i].Err != nil {
			m.err = m.inputs[i].Err
		}
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m AddPasswordModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		"%s\n\n%s:\n%s\n\n%s:\n%s\n\n%s:\n%s\n%s\n%s",
		titleStyle.Render("Add login/password"),
		inputLabelStyle.Render("Title"),
		m.inputs[secretPasswordName].View(),
		inputLabelStyle.Render("Login"),
		m.inputs[secretPasswordLogin].View(),
		inputLabelStyle.Render("Password"),
		m.inputs[secretPasswordPassword].View(),
		errToString(m.err),
		continueStyle.Render("enter: add, esc: back"),
	))
}

func (m *AddPasswordModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *AddPasswordModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
