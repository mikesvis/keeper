package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"keeper/internal/service/secret"
	"keeper/internal/service/user"
)

type SecretTypeMenuModel struct {
	parent MenuModel
	us     *user.Service
	ss     *secret.Service
	list   list.Model
}

func NewSecretTypeMenuModel(parent MenuModel, us *user.Service, ss *secret.Service) SecretTypeMenuModel {
	w, h := getListSizes()

	l := list.New([]list.Item{
		ModelItem{title: "Add login/password", desc: "Add login & password pair for site/service/etc"},
		ModelItem{title: "Add credit cart", desc: "Add credit card credentials"},
		ModelItem{title: "Add note", desc: "Add simple text note"},
		ModelItem{title: "Add file", desc: "Add binary file"},
	}, list.NewDefaultDelegate(), w, h)

	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.Title = "Chose new secret type"

	return SecretTypeMenuModel{
		parent: parent,
		us:     us,
		ss:     ss,
		list:   l,
	}
}

func (m SecretTypeMenuModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m SecretTypeMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m.parent, nil
		case tea.KeyEnter:
			switch m.list.Cursor() {
			case 0:
				return NewAddPasswordModel(m.parent, m.ss), nil
			case 1:
				return NewAddCreditCardModel(m.parent, m.ss), nil
			case 2:
				return NewNoteModel(m.parent, m.ss), nil
			//case 3:
			//	return NewAddSecretFileModel(m.parent, m.service), nil
			default:
				return m.parent, nil
			}
		}

	case tea.WindowSizeMsg:
		h, v := docForListStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m SecretTypeMenuModel) View() string {
	return docForListStyle.Render(m.list.View())
}
