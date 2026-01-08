package handlers

import (
	"encoding/json"
	"first_project/config"
	"first_project/models"
	"net/http"
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
