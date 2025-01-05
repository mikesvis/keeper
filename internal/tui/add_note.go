package tui

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"keeper/internal/domain"
	"keeper/internal/service/secret"
)

type AddNoteModel struct {
	parent   MenuModel
	ss       *secret.Service
	input    textinput.Model
	textarea textarea.Model
	focused  int
	err      error
}

func NewNoteModel(parent MenuModel, ss *secret.Service) AddNoteModel {
	return AddNoteModel{
		parent:   parent,
		ss:       ss,
		input:    NewInput(inputText, true),
		textarea: NewTextarea(false),
		focused:  0,
		err:      nil,
	}
}

func (m AddNoteModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddNoteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 2)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == 1 {
				name := m.input.Value()
				text := m.textarea.Value()

				rawData, err := domain.NewUserSecretText(text).GetData()
				if err != nil {
					m.err = err
					return m, nil
				}

				data, err := domain.MakeUserSecretData(domain.UserSecretTextType, rawData)
				if err != nil {
					m.err = err
					return m, nil
				}

				err = m.ss.Create(context.Background(), m.parent.GetUserID(), domain.UserSecretTextType, name, &data)
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

		m.input.Blur()
		m.textarea.Blur()

		if m.focused == 0 {
			m.input.Focus()
		} else {
			m.textarea.Focus()
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	if m.input.Err != nil {
		m.err = m.input.Err
	} else if m.textarea.Err != nil {
		m.err = m.textarea.Err
	}

	m.input, cmds[0] = m.input.Update(msg)
	m.textarea, cmds[1] = m.textarea.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m AddNoteModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		"%s\n\n%s:\n%s\n\n%s:\n%s\n%s\n%s",
		titleStyle.Render("Add note"),
		inputLabelStyle.Render("Title"),
		m.input.View(),
		inputLabelStyle.Render("Note"),
		m.textarea.View(),
		errToString(m.err),
		continueStyle.Render("enter: add, esc: back"),
	))
}

func (m *AddNoteModel) nextInput() {
	m.focused = (m.focused + 1) % 2
}

func (m *AddNoteModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = 1
	}
}
