package comands

import (
	"fmt"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"isu1.0/entities"
	"isu1.0/service"
	"isu1.0/style"
	"strconv"
	"strings"
)

type addGroup struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	c          *Cli
	service    *service.Service
}

func (c *Cli) AddGroup() {
	if _, err := tea.NewProgram(AddGroupForm(c)).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		c.Help()
	}
	c.Help()

}

func AddGroupForm(c *Cli) addGroup {
	m := addGroup{
		inputs: make([]textinput.Model, 3),
		c:      c,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = style.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Number"
			t.Focus()
			t.PromptStyle = style.FocusedStyle
			t.TextStyle = style.FocusedStyle
		case 1:
			t.Placeholder = "Faculty"
		case 2:
			t.Placeholder = "Course"
		}

		m.inputs[i] = t
	}

	return m
}

func (m addGroup) Init() tea.Cmd {
	return textinput.Blink
}

func (m addGroup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+d", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				group := m.Validator()

				err := m.c.service.AddGroup(group.GroupNumber, group.Faculty, group.Course)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Group added")
				}

				return m, tea.Quit

			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {

					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = style.FocusedStyle
					m.inputs[i].TextStyle = style.FocusedStyle
					continue
				}

				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = style.NoStyle
				m.inputs[i].TextStyle = style.NoStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m addGroup) Validator() entities.GroupInfo {
	var input entities.GroupInfo
	faculty := "FTMI,PIIKT,ITIP"
	input.GroupNumber = m.inputs[0].Value()
	if len(input.GroupNumber) != 5 {
		fmt.Println("Len of the group number should be 6")
	}
	input.Faculty = m.inputs[1].Value()
	for i := 0; i < len(faculty); i++ {
		if strings.Count(faculty, input.Faculty) != 1 {
			fmt.Println("faculty wasnt found")
		}
	}
	course := m.inputs[2].Value()
	input.Course, _ = strconv.Atoi(course)
	if input.Course < 0 && input.Course > 4 {
		fmt.Println("course wasnt found")
	}
	return input
}

func (m *addGroup) updateInputs(msg tea.Msg) tea.Cmd {

	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
func (m addGroup) View() string {

	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &style.BlurredButton
	if m.focusIndex == len(m.inputs) {
		button = &style.FocusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
