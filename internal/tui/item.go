package tui

type ModelItem struct {
	title, desc string
}

func (i ModelItem) Title() string       { return i.title }
func (i ModelItem) Description() string { return i.desc }
func (i ModelItem) FilterValue() string { return i.title }
