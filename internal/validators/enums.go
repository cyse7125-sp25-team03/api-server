package validators

// Define valid enum values using maps for fast lookup
var ValidSchools = map[string]bool{
	"Engineering": true,
	"Science":     true,
}

var ValidDepartments = map[string]bool{
	"Computer Science": true,
	"Mathematics":      true,
	"Physics":          true,
}

var ValidSemesterTerms = map[string]bool{
	"Fall 2024":   true,
	"Spring 2024": true,
	"Summer 2024": true,
}

// Generic validation function using maps
func IsValidEnum(value string, validValues map[string]bool) bool {
	return validValues[value]
}
