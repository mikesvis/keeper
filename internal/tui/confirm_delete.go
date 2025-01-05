package tui

import (
	"context"
	"fmt"
	"keeper/internal/service/secret"

	tea "github.com/charmbracelet/bubbletea"
)

type ConfirmDeleteModel struct {
	parent MenuModel
	ss     *secret.Service
	item   SecretListItem
	err    error
}

func NewConfirmDeleteModel(parent MenuModel, ss *secret.Service, item SecretListItem) *ConfirmDeleteModel {
	return &ConfirmDeleteModel{
		parent: parent,
		ss:     ss,
		item:   item,
	}
}

func (m ConfirmDeleteModel) Init() tea.Cmd {
	return nil
}

func (m ConfirmDeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.err = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			err := m.ss.Delete(context.Background(), m.item.ID)
			if err != nil {
				m.err = err
				return m, nil
			}

			return m.parent, nil
		case tea.KeyEsc:
			return m.parent, nil
		}
	}

	return m, nil
}

func (m ConfirmDeleteModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		"%s\n\n%s\n%s",
		titleStyle.Render("Delete secret?"),
		errToString(m.err),
		continueStyle.Render("enter: delete, esc: back"),
	) + "\n")
}
