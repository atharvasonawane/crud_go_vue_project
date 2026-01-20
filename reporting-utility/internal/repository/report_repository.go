package repository

import (
	"reporting-utility/internal/db"
)

// FetchAllStudents (Read)
func FetchAllStudents() ([]Student, error) {
	rows, err := db.DB.Query(`
		SELECT id, student_name, address, state, district, taluka,
		       gender, dob, photo, handicapped, email,
		       mobile_number, blood_group
		FROM students
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		err := rows.Scan(
			&s.ID, &s.StudentName, &s.Address, &s.State, &s.District, &s.Taluka,
			&s.Gender, &s.Dob, &s.Photo, &s.Handicapped, &s.Email,
			&s.MobileNumber, &s.BloodGroup,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

// GetStudentByID (Read Single)
func GetStudentByID(id int) (Student, error) {
	var s Student
	query := `
		SELECT id, student_name, address, state, district, taluka,
		       gender, dob, photo, handicapped, email,
		       mobile_number, blood_group
		FROM students
		WHERE id = ?
	`
	err := db.DB.QueryRow(query, id).Scan(
		&s.ID, &s.StudentName, &s.Address, &s.State, &s.District, &s.Taluka,
		&s.Gender, &s.Dob, &s.Photo, &s.Handicapped, &s.Email,
		&s.MobileNumber, &s.BloodGroup,
	)
	return s, err
}

// CreateStudent (Create)
func CreateStudent(s Student) (int, error) {
	query := `
		INSERT INTO students
		(student_name, address, state, district, taluka, gender, dob, photo,
		 handicapped, email, mobile_number, blood_group)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := db.DB.Exec(query,
		s.StudentName, s.Address, s.State, s.District, s.Taluka,
		s.Gender, s.Dob, s.Photo, s.Handicapped, s.Email,
		s.MobileNumber, s.BloodGroup,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

// UpdateStudent (Update)
func UpdateStudent(s Student) error {
	query := `
		UPDATE students SET
			student_name = ?, address = ?, state = ?, district = ?, taluka = ?,
			gender = ?, dob = ?, photo = ?, handicapped = ?, email = ?,
			mobile_number = ?, blood_group = ?
		WHERE id = ?
	`
	_, err := db.DB.Exec(query,
		s.StudentName, s.Address, s.State, s.District, s.Taluka,
		s.Gender, s.Dob, s.Photo, s.Handicapped, s.Email,
		s.MobileNumber, s.BloodGroup,
		s.ID,
	)
	return err
}

// DeleteStudent (Delete)
func DeleteStudent(id int) error {
	_, err := db.DB.Exec("DELETE FROM students WHERE id = ?", id)
	return err
}