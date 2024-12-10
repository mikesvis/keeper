package main

//type Model struct {
//	form *huh.Form // huh.Form is just a tea.Model
//}

func main() {
	//form := huh.NewForm(
	//	huh.NewGroup(huh.NewNote().
	//		Title("Password keeper").
	//		Description("Welcome to _Keeper_\n\n").
	//		Next(true).
	//		NextLabel("Go"),
	//	),
	//)
	//
	//err := form.Run()
	//if err != nil {
	//	panic(err)
	//}

	//p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	//if _, err := p.Run(); err != nil {
	//	log.Fatal(err)
	//}
}

//func NewModel() Model {
//	return Model{
//		form: huh.NewForm(
//			huh.NewGroup(
//				huh.NewSelect[string]().
//					Key("class").
//					Options(huh.NewOptions("Warrior", "Mage", "Rogue")...).
//					Title("Choose your class"),
//
//				huh.NewSelect[int]().
//					Key("level").
//					Options(huh.NewOptions(1, 20, 9999)...).
//					Title("Choose your level"),
//			),
//		),
//	}
//}
//
//func (m Model) Init() tea.Cmd {
//	return m.form.Init()
//}
//
//func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	// ...
//
//	form, cmd := m.form.Update(msg)
//	if f, ok := form.(*huh.Form); ok {
//		m.form = f
//	}
//
//	return m, cmd
//}
//
//func (m Model) View() string {
//	if m.form.State == huh.StateCompleted {
//		class := m.form.GetString("class")
//		level := m.form.GetString("level")
//		return fmt.Sprintf("You selected: %s, Lvl. %d", class, level)
//	}
//	return m.form.View()
//}
