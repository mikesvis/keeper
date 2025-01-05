package tui

import (
	"github.com/google/uuid"
	"keeper/internal/domain"
	"keeper/internal/service/secret"
	"keeper/internal/service/user"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	auth *domain.AuthenticatedUser
	us   *user.Service
	ss   *secret.Service
	list list.Model
}

func NewMenuModel(auth *domain.AuthenticatedUser, us *user.Service, ss *secret.Service) MenuModel {
	w, h := getListSizes()

	l := list.New([]list.Item{
		ModelItem{title: "Add a secret", desc: "Add and store new secret"},
		ModelItem{title: "Show all secrets", desc: "Show list of stored secrets"},
		ModelItem{title: "Delete secret", desc: "Delete stored secret"},
	}, list.NewDefaultDelegate(), w, h)

	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.Title = "Choose an option"

	return MenuModel{
		auth: auth,
		list: l,
		us:   us,
		ss:   ss,
	}
}

func (m MenuModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.list.Cursor() {
			case 0:
				return NewSecretTypeMenuModel(m, m.us, m.ss), nil
			case 1, 2:
				//secrets, _ := m.service.GetUserSecrets(context.Background())
				// todo tui error
				return m, nil
				//return NewSecretsList(m, secrets, m.service, m.list.Cursor() == 2), nil
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

func (m MenuModel) View() string {
	return docForListStyle.Render(m.list.View())
}

func (m MenuModel) GetUserID() uuid.UUID {
	return m.auth.ID
}
