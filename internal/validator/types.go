package validator

// ValidationError represents a single linting or schema error found in a file.
type ValidationError struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Message string `json:"message"`
	Rule    string `json:"rule"`
}
