package service

import (
	"isu1.0/entities"
	"isu1.0/repository"
)

type Commands interface {
	AddStudent(name, lastname, group string, isu int) error
	AddGroup(number, faculty string, course int) error
	GroupChecker(groupNumber string) error
	TransferStudent(isu int, groupNumber string) error
	FindStudent(param string) ([]entities.Student, error)
	//find student by isu
	//find student by course number
	//find student by group number
}

type Service struct {
	Commands
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Commands: NewCommandService(repos.Commands),
	}
}
