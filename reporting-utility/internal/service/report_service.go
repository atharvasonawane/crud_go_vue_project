package service

import (
	"reporting-utility/internal/repository"
)

// ListStudents
func ListStudents() ([]repository.Student, error) {
	return repository.FetchAllStudents()
}

// GetStudent
func GetStudent(id int) (repository.Student, error) {
	return repository.GetStudentByID(id)
}

// AddStudent
func AddStudent(s repository.Student) (repository.Student, error) {
	id, err := repository.CreateStudent(s)
	if err != nil {
		return s, err
	}
	s.ID = id
	return s, nil
}

// EditStudent
func EditStudent(s repository.Student) (repository.Student, error) {
	// If needed, we could fetch the existing student first to merge data, 
	// but strictly following the old logic, we just run Update.
	err := repository.UpdateStudent(s)
	return s, err
}

// RemoveStudent
func RemoveStudent(id int) error {
	return repository.DeleteStudent(id)
}