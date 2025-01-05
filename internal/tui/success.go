package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type SuccessModel struct {
	parent MenuModel
}

func NewSuccessModel(parent MenuModel) SuccessModel {
	return SuccessModel{
		parent: parent,
	}
}

func (m SuccessModel) Init() tea.Cmd {
	return nil
}

func (m SuccessModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter, tea.KeyEsc:
			return m.parent, nil
		}
	}

	return m, nil
}

func (m SuccessModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		"%s\n\n%s",
		titleStyle.Render("Success"),
		continueStyle.Render("enter: continue"),
	) + "\n")
}
