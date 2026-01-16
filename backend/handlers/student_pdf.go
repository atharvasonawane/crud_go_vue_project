package handlers

import (
	"first_project/config"
	"first_project/models"
	"fmt"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

func DownloadStudentsPDF(w http.ResponseWriter, r *http.Request) {

	rows, err := config.DB.Query(`
	SELECT id, student_name, email, mobile_number, blood_group
	FROM students
	`)

	if err != nil {
		http.Error(w, "Failed to fetch students", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var s models.Student
		rows.Scan(
			&s.ID,
			&s.StudentName,
			&s.Email,
			&s.MobileNumber,
			&s.BloodGroup,
		)
		students = append(students, s)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, "Student List")
	pdf.Ln(15)

	// Table Header
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(10, 10, "ID")
	pdf.Cell(40, 10, "Name")
	pdf.Cell(50, 10, "Email")
	pdf.Cell(40, 10, "Mobile")
	pdf.Cell(30, 10, "Blood")
	pdf.Ln(8)

	// Table Body
	pdf.SetFont("Arial", "", 10)
	for _, s := range students {
		pdf.Cell(10, 8, fmt.Sprintf("%d", s.ID))
		pdf.Cell(40, 10, s.StudentName)
		pdf.Cell(50, 10, s.Email)
		pdf.Cell(40, 10, s.MobileNumber)
		pdf.Cell(30, 10, s.BloodGroup)
		pdf.Ln(8)
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=students.pdf")

	err = pdf.Output(w)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}
}
