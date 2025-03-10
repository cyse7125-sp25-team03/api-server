package validators

import (
	"errors"
	"strings"

	"api-server/internal/models"

	"github.com/google/uuid"
)

// validate Filename
func ValidateFilename(filename string) error {
	if strings.TrimSpace(filename) == "" {
		return errors.New("Filename cannot be empty or just blank")
	}
	return nil
}

// validate Bucket Path
func ValidateBucketPath(bucketPath string) error {
	if strings.TrimSpace(bucketPath) == "" {
		return errors.New("Bucket Path cannot be empty or just blank")
	}
	return nil
}

// validate Section
func ValidateSection(section string) error {
	if strings.TrimSpace(section) == "" {
		return errors.New("Section cannot be empty or just blank")
	}
	return nil
}

// validate courseID
func ValidateCourseID(courseID string) error {
	if strings.TrimSpace(courseID) == "" {
		return errors.New("Course ID cannot be empty or just blank")
	}
	if _, err := uuid.Parse(courseID); err != nil {
		return errors.New("Invalid UUID format for course ID")
	}
	return nil
}

// validate semester term
func ValidateSemesterTerm(semesterTerm string) error {
	if strings.TrimSpace(semesterTerm) == "" {
		return errors.New("Semester term cannot be empty or just blank")
	}
	return nil
}

// ValidateTraceRequest validates the request body for creating a new trace
func ValidateTraceRequest(trace models.TraceRequest) error {
	if err := ValidateFilename(trace.FileName); err != nil {
		return err
	}
	if err := ValidateBucketPath(trace.BucketPath); err != nil {
		return err
	}
	if err := ValidateSection(trace.Section); err != nil {
		return err
	}
	if err := ValidateSemesterTerm(trace.SemesterTerm); err != nil {
		return err
	}
	if err := ValidateCourseInstructorID(trace.InstructorID); err != nil {
		return err
	}
	return nil
}
