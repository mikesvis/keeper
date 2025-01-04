package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

const (
	hotPink  = lipgloss.Color("#EE6FF8")
	darkGray = lipgloss.Color("#767676")
	red      = lipgloss.Color("#EE204D")
	green    = lipgloss.Color("#5fb458")
)

var (
	docForListStyle = lipgloss.NewStyle().Margin(1, 2)
	docStyle        = lipgloss.NewStyle().Margin(1, 4)
	titleStyle      = lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 1)
	inputLabelStyle = lipgloss.NewStyle().Foreground(hotPink)
	errorStyle      = lipgloss.NewStyle().Foreground(red)
	successStyle    = lipgloss.NewStyle().Foreground(green)
	continueStyle   = lipgloss.NewStyle().Foreground(darkGray)
)

func errToString(err error) string {
	message := ""
	if err != nil {
		message = "\n" + errorStyle.Render(err.Error()) + "\n"
	}

	return message
}

func getListSizes() (int, int) {
	w, h, _ := term.GetSize(0)

	dw, dh := docForListStyle.GetFrameSize()

	return w - dw, h - dh
}
