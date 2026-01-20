package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"reporting-utility/internal/db"
	"reporting-utility/internal/report"
	"reporting-utility/internal/repository"
	"reporting-utility/internal/service"
	"time"
)

// --- GET ALL ---
func GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := service.ListStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// --- CREATE ---
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // max 10MB

	student := repository.Student{
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

	// Handle File Upload
	file, handler, err := r.FormFile("photo")
	if err == nil {
		defer file.Close()
		os.MkdirAll("uploads", os.ModePerm)
		dst, _ := os.Create("uploads/" + handler.Filename)
		defer dst.Close()
		dst.ReadFrom(file)
		student.Photo = handler.Filename
	}

	createdStudent, err := service.AddStudent(student)
	if err != nil {
		http.Error(w, "Failed to create student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdStudent)
}

// --- UPDATE ---
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get ID from session (Old logic requirement)
	session, _ := db.Store.Get(r, "student-session")
	rawID := session.Values["student_id"]
	if rawID == nil {
		http.Error(w, "No student selected", http.StatusBadRequest)
		return
	}

	id, ok := rawID.(int)
	if !ok {
		// Handle case where session stored float64 (common in JSON serialization)
		if val, okFloat := rawID.(float64); okFloat {
			id = int(val)
		} else {
			http.Error(w, "Invalid session data", http.StatusBadRequest)
			return
		}
	}

	var student repository.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Format Date
	if student.Dob != "" {
		dob, err := time.Parse("2006-01-02", student.Dob)
		if err == nil {
			student.Dob = dob.Format("2006-01-02")
		}
	}

	student.ID = id // Enforce ID from session
	updatedStudent, err := service.EditStudent(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Clear session
	delete(session.Values, "student_id")
	session.Save(r, w)

	json.NewEncoder(w).Encode(updatedStudent)
}

// --- DELETE ---
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	session, _ := db.Store.Get(r, "student-session")
	rawID := session.Values["student_id"]
	if rawID == nil {
		http.Error(w, "No student selected", http.StatusBadRequest)
		return
	}

	id, ok := rawID.(int)
	if !ok {
		if val, okFloat := rawID.(float64); okFloat {
			id = int(val)
		} else {
			http.Error(w, "Invalid session data", http.StatusBadRequest)
			return
		}
	}

	if err := service.RemoveStudent(id); err != nil {
		http.Error(w, "Failed to delete student", http.StatusInternalServerError)
		return
	}

	delete(session.Values, "student_id")
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student deleted successfully"))
}

// --- SESSION HANDLERS ---
func SelectStudent(w http.ResponseWriter, r *http.Request) {
	session, _ := db.Store.Get(r, "student-session")
	var data struct {
		StudentID int `json:"student_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	session.Values["student_id"] = data.StudentID
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Student selected"})
}

func GetSelectedStudent(w http.ResponseWriter, r *http.Request) {
	session, _ := db.Store.Get(r, "student-session")
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

	student, err := service.GetStudent(id)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// --- PDF ---
func DownloadPDF(w http.ResponseWriter, r *http.Request) {
	students, err := service.ListStudents()
	if err != nil {
		http.Error(w, "Failed to fetch data", 500)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=students.pdf")
	report.GenerateStudentPDF(students, w)
}