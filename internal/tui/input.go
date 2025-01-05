package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

type inputType string

var (
	inputText     inputType = "text"
	inputPassword inputType = "password"
)

func NewInput(t inputType, focus bool) textinput.Model {
	input := textinput.New()
	input.Placeholder = ""
	input.CharLimit = 255
	input.Width = 50
	input.Prompt = ""

	if focus {
		input.Focus()
	}

	if t == inputPassword {
		input.Placeholder = "******"
		input.EchoMode = textinput.EchoPassword
		input.EchoCharacter = '•'
	}

	return input
}

func NewTextarea(focus bool) textarea.Model {
	ti := textarea.New()

	if focus {
		ti.Focus()
	}

	return ti
}
