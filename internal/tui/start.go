package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"keeper/internal/service/secret"
	"keeper/internal/service/user"
)

type StartModel struct {
	list list.Model
	us   *user.Service
	ss   *secret.Service
}

func NewStart(us *user.Service, ss *secret.Service) StartModel {
	l := list.New([]list.Item{
		ModelItem{title: "Sign-In", desc: "Sing in by email and password"},
		ModelItem{title: "Sign-Up", desc: "Sign up for new user"},
	}, list.NewDefaultDelegate(), 0, 0)

	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.Title = "Authorization required"

	return StartModel{
		list: l,
		us:   us,
		ss:   ss,
	}
}

func (m StartModel) Init() tea.Cmd {
	return nil
}

func (m StartModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.list.Cursor() {
			case 0:
				return NewLoginModel(m, m.us, m.ss), nil
			case 1:
				return NewRegisterModel(m, m.us, m.ss), nil
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

func (m StartModel) View() string {
	return docForListStyle.Render(m.list.View())
}
