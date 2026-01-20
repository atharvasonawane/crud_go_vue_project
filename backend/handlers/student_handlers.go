package handlers

import (
	"encoding/json"
	"first_project/config"
	"first_project/models"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20) // max 10MB

	student := models.Student{
		StudentName:  r.FormValue("studentName"),
		Address:      r.FormValue("address"),
		State:        r.FormValue("state"),
		District:     r.FormValue("district"),
		Taluka:       r.FormValue("taluka"),
		Gender:       r.FormValue("gender"),
		Dob:          r.FormValue("dob"),
		Handicapped:  r.FormValue("handicapped") == "true",
		Email:        r.FormValue("email"),
		MobileNumber: r.FormValue("mobileNumber"),
		BloodGroup:   r.FormValue("bloodGroup"),
	}

	file, handler, err := r.FormFile("photo")
	if err == nil {
		defer file.Close()
		os.MkdirAll("uploads", os.ModePerm)
		dst, _ := os.Create("uploads/" + handler.Filename)
		defer dst.Close()
		dst.ReadFrom(file)
		student.Photo = handler.Filename
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

	session, _ := config.Store.Get(r, "student-session")

	rawID := session.Values["student_id"]
	if rawID == nil {
		http.Error(w, "No student selected", http.StatusBadRequest)
		return
	}

	id, ok := rawID.(int)
	if !ok {
		http.Error(w, "Invalid session data", http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec(
		"DELETE FROM students WHERE id = ?",
		id,
	)

	if err != nil {
		http.Error(w, "Failed to delete student", http.StatusInternalServerError)
		return
	}

	// clear session after delete
	delete(session.Values, "student_id")
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student deleted successfully"))
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, _ := config.Store.Get(r, "student-session")

	rawID := session.Values["student_id"]
	if rawID == nil {
		http.Error(w, "No student selected", http.StatusBadRequest)
		return
	}

	id, ok := rawID.(int)
	if !ok {
		http.Error(w, "Invalid session data", http.StatusBadRequest)
		return
	}

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if student.Dob != "" {
		dob, err := time.Parse("2006-01-02", student.Dob)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}
		student.Dob = dob.Format("2006-01-02")
	}

	_, err := config.DB.Exec(`
		UPDATE students SET
			student_name = ?,
			address = ?,
			state = ?,
			district = ?,
			taluka = ?,
			gender = ?,
			dob = ?,
			photo = ?,
			handicapped = ?,
			email = ?,
			mobile_number = ?,
			blood_group = ?
		WHERE id = ?
	`,
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
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delete(session.Values, "student_id")
	session.Save(r, w)

	student.ID = id
	json.NewEncoder(w).Encode(student)
}

// Handler to store student ID in session
func SelectStudent(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, "student-session")

	var data struct {
		StudentID int `json:"student_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	session.Values["student_id"] = data.StudentID
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Student selected",
	})
}

// Function to get student using session
func GetSelectedStudent(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, "student-session")

	rawID := session.Values["student_id"]

	if rawID == nil {
		http.Error(w, "No student selected", http.StatusBadRequest)
		return
	}

	var id int

	switch v := rawID.(type) {
	case int:
		id = v
	case float64:
		id = int(v)
	default:
		http.Error(w, "Invalid session data", http.StatusBadRequest)
		return
	}

	if id <= 0 {
		http.Error(w, "No student selected", http.StatusBadRequest)
		return
	}

	var student models.Student
	query := `		SELECT id, student_name, address, state, district, taluka,
		       gender, dob, photo, handicapped, email,
		       mobile_number, blood_group
		FROM students
		WHERE id = ?`

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