package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:type`
}

type TestCase struct {
	Input  string
	Output string
}

type Assigment struct {
	AssigmentID string `json:"assigmentId"`
	Question    string `json:"question"`
	TestCases   []TestCase
}

type Batch struct {
	Name        string
	Students    []Student
	Assignments []Assignment
}

type Submission struct {
	AssignmentID string            `json:"assignmentId`
	Code         string            `json:"code"`
	Output       string            `json:output"`
	Pass         map[TestCase]bool `json:"pass"`
}

type Student struct {
	User
	Submissions []Submission
}

type Teacher struct {
	Classes []Batch
	User
}

func (u *User) HashPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(password)
	return nil
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
