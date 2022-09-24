package models

type User struct {
	ID       int
	Username string
	Password string // hash
}

var SaveData []User

func (s *User) SaveUser() []User {
	SaveData = append(SaveData, *s)
	return SaveData
}