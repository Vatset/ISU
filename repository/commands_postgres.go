package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"isu1.0/entities"
)

type CommandRepository struct {
	db *sqlx.DB
}

func NewCommandRepository(db *sqlx.DB) *CommandRepository {
	return &CommandRepository{db: db}
}

func (r *CommandRepository) AddStudent(name, lastname, group string, isu int) error {
	createAllStudentQuery := fmt.Sprintf("INSERT INTO %s (name, lastname, isu, groupnumber) values ($1, $2, $3, $4)", studentsTable)
	_, err := r.db.Exec(createAllStudentQuery, name, lastname, isu, group)
	if err != nil {
		return err
	}
	return err
}
func (r *CommandRepository) GroupChecker(groupNumber string) error {
	query := fmt.Sprintf("SELECT  COUNT(*)  FROM %s WHERE number=$1", groupsTable)
	var countOfGroups int
	err := r.db.QueryRow(query, groupNumber).Scan(&countOfGroups)
	if err != nil {
		return err
	}
	if countOfGroups != 1 {
		return errors.New("Group number wasnt found")
	}
	return err
}
func (r *CommandRepository) AddGroup(number, faculty string, course int) error {
	createAllStudentQuery := fmt.Sprintf("INSERT INTO %s (number, faculty, course) values ($1, $2, $3)", groupsTable)
	_, err := r.db.Exec(createAllStudentQuery, number, faculty, course)
	if err != nil {
		return err
	}
	return err
}
func (r *CommandRepository) TransferStudent(isu int, groupNumber string) error {
	createAllStudentQuery := fmt.Sprintf("UPDATE %s SET groupnumber=$1 WHERE isu = $2", studentsTable)
	_, err := r.db.Exec(createAllStudentQuery, groupNumber, isu)
	if err != nil {
		return errors.New("Student wasnt found")
	}
	return err
}
func (r *CommandRepository) FindStudent(param string) ([]entities.Student, error) {
	var students []entities.Student
	query := fmt.Sprintf("SELECT name, lastname, isu FROM %s WHERE %s", studentsTable, param)
	err := r.db.Select(&students, query)
	if err != nil {
		return nil, err
	}

	return students, err
}
