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

type finder struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	c          *Cli
	service    *service.Service
}

func (c *Cli) Finder() {
	if _, err := tea.NewProgram(ll(c)).Run(); err != nil {
		fmt.Printf("could not start gram: %s\n", err)
		c.Help()
	}
}

func ll(c *Cli) finder {
	result, counter := c.FindStudent()

	m := finder{
		inputs: make([]textinput.Model, counter),
		c:      c,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = style.CursorStyle
		t.CharLimit = 32

		if i < len(result) {
			t.Placeholder = result[i]

			if i == 0 {
				t.Focus()
				t.PromptStyle = style.FocusedStyle
				t.TextStyle = style.FocusedStyle
			}
		}
		m.inputs[i] = t
	}

	return m
}

func (m finder) Init() tea.Cmd {
	return textinput.Blink
}

func (m finder) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+d", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {

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

func (m finder) Validator() entities.Student {
	var input entities.Student
	input.Name = m.inputs[0].Value()
	input.Lastname = m.inputs[1].Value()
	isuStr := m.inputs[2].Value()
	input.ISU, _ = strconv.Atoi(isuStr)
	return input
}

func (m *finder) updateInputs(msg tea.Msg) tea.Cmd {

	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
func (m finder) View() string {

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
