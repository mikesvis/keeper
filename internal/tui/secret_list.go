package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"keeper/internal/domain"
	"keeper/internal/service/secret"
)

type SecretsListModel struct {
	parent         MenuModel
	ss             *secret.Service
	list           list.Model
	deleteOnSelect bool
}

func NewSecretsListModel(
	parent MenuModel,
	secrets []*domain.UserSecret,
	ss *secret.Service,
	deleteOnSelect bool,
) SecretsListModel {
	delegate := list.NewDefaultDelegate()

	items := make([]list.Item, 0)
	for _, oneSecret := range secrets {
		items = append(items, SecretListItem{*oneSecret})
	}

	w, h := getListSizes()

	l := list.New(items, delegate, w, h)
	l.Title = "Your secrets"

	return SecretsListModel{
		parent:         parent,
		ss:             ss,
		list:           l,
		deleteOnSelect: deleteOnSelect,
	}
}

func (m SecretsListModel) Init() tea.Cmd {
	return nil
}

func (m SecretsListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m.parent, nil
		case tea.KeyEnter:
			if m.deleteOnSelect {
				item := m.list.SelectedItem().(SecretListItem)

				return NewConfirmDeleteModel(m.parent, m.ss, item), nil
			}

			return m.parent, nil
		}

	case tea.WindowSizeMsg:
		h, v := docForListStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m SecretsListModel) View() string {
	return docForListStyle.Render(m.list.View())
}
