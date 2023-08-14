package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"isu1.0/entities"
)

type Commands interface {
	AddStudent(name, lastname, group string, isu int) error
	AddGroup(number, faculty string, course int) error
	GroupChecker(groupNumber string) error
	TransferStudent(isu int, groupNumber string) error
	FindStudent(param string) ([]entities.Student, error)
}

type Repository struct {
	Commands
}

func NewRepository(db *sqlx.DB) *Repository {
	if db == nil {
		fmt.Println("db nil")
		return nil
	}
	return &Repository{
		Commands: NewCommandRepository(db),
	}
}
