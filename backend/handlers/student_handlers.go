package handlers

import (
	"encoding/json"
	"first_project/config"
	"first_project/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	// Read JSON body
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO students
		(student_name, address, state, district, taluka, gender, dob, photo,
		 handicapped, email, mobile_number, blood_group)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := config.DB.Exec(query,
		student.StudentName,
		student.Address,
		student.State,
		student.District,
		student.Taluka,
		student.Gender,
		student.Dob,
		student.Photo,
		student.Handicapped,
		student.Email,
		student.MobileNumber,
		student.BloodGroup,
	)

	if err != nil {
		http.Error(w, "Failed to create student", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()

	student.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)

}

func GetStudents(w http.ResponseWriter, r *http.Request) {

	rows, err := config.DB.Query(`
		SELECT id, student_name, address, state, district, taluka,
		       gender, dob, photo, handicapped, email,
		       mobile_number, blood_group
		FROM students
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var s models.Student
		err := rows.Scan(
			&s.ID,
			&s.StudentName,
			&s.Address,
			&s.State,
			&s.District,
			&s.Taluka,
			&s.Gender,
			&s.Dob,
			&s.Photo,
			&s.Handicapped,
			&s.Email,
			&s.MobileNumber,
			&s.BloodGroup,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var student models.Student
	query := `
		SELECT id, student_name, address, state, district, taluka,
		       gender, dob, photo, handicapped, email,
		       mobile_number, blood_group
		FROM students
		WHERE id = ?
	`

	err := config.DB.QueryRow(query, id).Scan(
		&student.ID,
		&student.StudentName,
		&student.Address,
		&student.State,
		&student.District,
		&student.Taluka,
		&student.Gender,
		&student.Dob,
		&student.Photo,
		&student.Handicapped,
		&student.Email,
		&student.MobileNumber,
		&student.BloodGroup,
	)

	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := config.DB.Exec(
		"DELETE FROM students WHERE id = ?",
		id,
	)

	if err != nil {
		http.Error(w, "Failed to delete student", http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student deleted successfully"))
}