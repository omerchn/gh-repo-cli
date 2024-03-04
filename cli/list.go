package cli

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func List(title string, options []ListItem) (string, error) {

	program := tea.NewProgram(initialListModel(title, options), tea.WithAltScreen())

	teaModel, err := program.Run()
	if err != nil {
		return "", err
	}

	model := teaModel.(ListModel)
	value := model.selected

	return value, nil
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type ListItem struct {
	TitleText       string
	DescriptionText string
}

func (i ListItem) Title() string       { return i.TitleText }
func (i ListItem) Description() string { return i.DescriptionText }
func (i ListItem) FilterValue() string { return i.TitleText }

type ListModel struct {
	list     list.Model
	selected string
}

func initialListModel(title string, options []ListItem) ListModel {
	items := make([]list.Item, len(options))
	for i, opt := range options {
		items[i] = ListItem(opt)
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = title

	m := ListModel{
		list: l,
	}

	return m
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.list.FilterState() != list.Filtering {
				m.selected = m.list.SelectedItem().FilterValue()
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListModel) View() string {
	return docStyle.Render(m.list.View())
}
