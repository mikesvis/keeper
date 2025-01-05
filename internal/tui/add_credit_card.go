package tui

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"keeper/internal/domain"
	"keeper/internal/service/secret"
	"strconv"
	"strings"
)

const (
	ccn = iota
	exp
	cvv
)

type AddCreditCardModel struct {
	parent  MenuModel
	ss      *secret.Service
	inputs  []textinput.Model
	focused int
	err     error
}

func NewAddCreditCardModel(parent MenuModel, ss *secret.Service) AddCreditCardModel {
	inputs := make([]textinput.Model, 3)

	inputs[ccn] = NewInput(inputText, true)
	inputs[exp] = NewInput(inputText, false)
	inputs[cvv] = NewInput(inputText, false)

	inputs[ccn].Placeholder = "4505 **** **** 1234"
	inputs[ccn].CharLimit = 20
	inputs[ccn].Width = 30

	inputs[exp].Placeholder = "MM/YY"
	inputs[exp].CharLimit = 5
	inputs[exp].Width = 5

	inputs[cvv].Placeholder = "XXX"
	inputs[cvv].CharLimit = 3
	inputs[cvv].Width = 5

	return AddCreditCardModel{
		parent:  parent,
		ss:      ss,
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m AddCreditCardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddCreditCardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				ccnV := m.inputs[ccn].Value()
				expV := m.inputs[exp].Value()
				cvvV := m.inputs[cvv].Value()

				parts := strings.Split(expV, "/")
				month, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					m.err = fmt.Errorf("EXP is invalid")
					return m, nil
				}
				year, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					m.err = fmt.Errorf("EXP is invalid")
					return m, nil
				}

				year += 2000

				cvvI, err := strconv.ParseInt(cvvV, 10, 64)
				if err != nil {
					m.err = fmt.Errorf("CVV is invalid")
					return m, nil
				}

				rawData, err := domain.NewUserSecretBankCard(ccnV, month, year, cvvI).GetData()
				if err != nil {
					m.err = err
					return m, nil
				}

				data, err := domain.MakeUserSecretData(domain.UserSecretBankCardType, rawData)
				if err != nil {
					m.err = err
					return m, nil
				}

				err = m.ss.Create(context.Background(), m.parent.GetUserID(), domain.UserSecretBankCardType, ccnV, &data)
				if err != nil {
					m.err = err
					return m, nil
				}

				return NewSuccessModel(m.parent), nil
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

func (m AddCreditCardModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		"%s\n\n%s:\n%s\n\n%s:\n%s\n\n%s:\n%s\n%s\n%s",
		titleStyle.Render("Add credit card"),
		inputLabelStyle.Render("Card Number"),
		m.inputs[ccn].View(),
		inputLabelStyle.Render("EXP"),
		m.inputs[exp].View(),
		inputLabelStyle.Render("CVV"),
		m.inputs[cvv].View(),
		errToString(m.err),
		continueStyle.Render("enter: add, esc: back"),
	))
}

func (m *AddCreditCardModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *AddCreditCardModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
