package services

import (
	"go-smaole/repo"
)

func Get() []repo.Teacher {
	teachers := repo.GetTeachers()
	return teachers
}
