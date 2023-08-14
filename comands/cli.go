package comands

import (
	tea "github.com/charmbracelet/bubbletea"
	"isu1.0/service"
)

type Cli struct {
	service *service.Service
}

func (c *Cli) InitialView() tea.Model {
	return c.Help()
}
func NewCLi(service *service.Service) *Cli {
	return &Cli{service: service}
}
