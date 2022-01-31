package services

import (
	"go-smaole/repo"
)

func GetTeacher() []repo.Teacher {
	teachers := repo.GetTeachers()
	return teachers
}

func CreateTeacher(name string, age int) int64 {
	id := repo.InserTeacher(name, age)
	return id
}

func UpdateTeacher(name string, age int, id int64) {
	repo.UpdateTeacher(name, age, id)
}

func DeleteTeacher(id int64) {
	repo.DeleteTeacher(id)
}
