package repositories

import (
	"api-server/internal/models"
	"database/sql"
	"time"
)

// CreateCourse creates a new course in the database

func CreateCourse(db *sql.DB, course models.Course) (models.Course, error) {

	_, err := db.Exec(
		"INSERT INTO webapp.courses (course_id, date_added, date_last_updated, user_id, code, name, description, instructor_id, department, school, credit_hours, semester_term, section, enrollment_count) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		course.CourseID, course.DateAdded, course.DateLastUpdated, course.UserID, course.Code, course.Name, course.Description, course.InstructorID, course.Department, course.School, course.CreditHours, course.SemesterTerm, course.Section, course.EnrollmentCount,
	)
	if err != nil {
		return models.Course{}, err
	}
	return course, err
}

// GetCourseByID retrieves a course by its ID
func GetCourseByID(db *sql.DB, courseID string) (*models.Course, error) {
	course := &models.Course{}
	err := db.QueryRow(
		"SELECT course_id, date_added, date_last_updated, user_id, code, name, description, instructor_id, department, school, credit_hours, semester_term, section, enrollment_count FROM webapp.courses WHERE course_id = $1",
		courseID,
	).Scan(&course.CourseID, &course.DateAdded, &course.DateLastUpdated, &course.UserID, &course.Code, &course.Name, &course.Description, &course.InstructorID, &course.Department, &course.School, &course.CreditHours, &course.SemesterTerm, &course.Section, &course.EnrollmentCount)
	return course, err
}

// UpdateCourse updates a course in the database
func UpdateCourse(db *sql.DB, course *models.Course) error {
	course.DateLastUpdated = time.Now().UTC()
	_, err := db.Exec(
		"UPDATE webapp.courses SET date_last_updated=$1, code=$2, name=$3, description=$4, instructor_id=$5, department=$6, school=$7, credit_hours=$8, semester_term=$9, section=$10, enrollment_count=$11 WHERE course_id=$12",
		course.DateLastUpdated, course.Code, course.Name, course.Description, course.InstructorID, course.Department, course.School, course.CreditHours, course.SemesterTerm, course.Section, course.EnrollmentCount, course.CourseID,
	)
	return err
}

// delete course by ID
func DeleteCourse(db *sql.DB, courseID string) error {
	result, err := db.Exec(
		"DELETE FROM webapp.courses WHERE course_id = $1",
		courseID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
