package service

import (
	"isu1.0/entities"
	"isu1.0/repository"
)

type CommandService struct {
	repo repository.Commands
}

func NewCommandService(repo repository.Commands) *CommandService {
	return &CommandService{repo: repo}
}
func (s *CommandService) AddStudent(name, lastname, group string, isu int) error {
	return s.repo.AddStudent(name, lastname, group, isu)
}
func (s *CommandService) AddGroup(number, faculty string, course int) error {
	return s.repo.AddGroup(number, faculty, course)
}
func (s *CommandService) GroupChecker(groupNumber string) error {
	return s.repo.GroupChecker(groupNumber)
}
func (s *CommandService) TransferStudent(isu int, groupNumber string) error {
	return s.repo.TransferStudent(isu, groupNumber)
}
func (s *CommandService) FindStudent(param string) ([]entities.Student, error) {
	return s.repo.FindStudent(param)
}
