package comands

import (
	"fmt"
	"isu1.0/service"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"by ISU", "by gr", "by course"}

type findStudent struct {
	cursor  int
	choice  []string
	c       *Cli
	service *service.Service
}

func (m findStudent) Init() tea.Cmd {
	return nil
}

func (m findStudent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		case "enter":
			return m, tea.Quit
		case "y":
			if m.cursor >= 0 && m.cursor < len(choices) {
				selectedChoice := choices[m.cursor]
				if !contains(m.choice, selectedChoice) {
					m.choice = append(m.choice, selectedChoice)
				} else {
					for i, choice := range m.choice {
						if choice == selectedChoice {
							m.choice = append(m.choice[:i], m.choice[i+1:]...)
							break
						}
					}
				}
			}

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}
func contains(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func (m findStudent) View() string {
	s := strings.Builder{}
	s.WriteString("By what criteria do you want to find a student?\n\n")

	for i := 0; i < len(choices); i++ {
		chosen := false
		for _, ch := range m.choice {
			if ch == choices[i] {
				chosen = true
				break
			}
		}

		if m.cursor == i || chosen {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}

	s.WriteString("\n(press y to choose)\n")
	s.WriteString("\n(press enter to go next)\n")

	return s.String()
}

func (c *Cli) FindStudent() ([]string, int) {
	p := tea.NewProgram(findStudent{})

	m, err := p.Run()
	if err != nil {
		fmt.Printf("could not start program: %s\n", err)
		c.Help()
	}
	if m, ok := m.(findStudent); ok && len(m.choice) != 0 {
		fmt.Printf("\n---\nYou chose %s!\n", m.choice)
		return m.choice, len(m.choice)
	}
	return nil, 0

}
