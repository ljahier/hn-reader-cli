package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor  int
	stories []string
}

func (m model) Init() tea.Cmd {
	return GetStories
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case storyMsg:
		m.stories = msg

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.stories)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string = "Hacker News Stories\n\n"
	var content string

	for i, story := range m.stories {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		content += fmt.Sprintf("%s * %s\n", cursor, story)
	}

	if len(content) > 0 {
		content += "\nPress q to quit.\n"
	}

	return s + content
}
