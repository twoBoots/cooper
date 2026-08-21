package validator

import (
	"fmt"
	"os"
	"strings"
)

// LintSpecFile reads and lints a specification or spec-delta markdown file.
func LintSpecFile(filePath string, isDelta bool) ([]ValidationError, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file %s: %w", filePath, err)
	}
	return LintSpecContent(filePath, string(data), isDelta), nil
}

// LintSpecContent validates the content of a capability specification or spec delta.
func LintSpecContent(filename string, content string, isDelta bool) []ValidationError {
	var errors []ValidationError
	lines := strings.Split(content, "\n")

	hasTitle := false
	hasRequirementsSection := false

	inScenario := false
	inRequirement := false
	requirementHasModal := false
	requirementLine := 0

	for i, rawLine := range lines {
		lineNum := i + 1
		line := strings.TrimSpace(rawLine)

		if line == "" || (isDelta && (line == "+" || line == "-")) {
			continue
		}

		// Clean delta prefix if present
		cleanLine := line
		hasDeltaPrefix := false
		if isDelta {
			if strings.HasPrefix(line, "+ ") {
				cleanLine = strings.TrimSpace(strings.TrimPrefix(line, "+ "))
				hasDeltaPrefix = true
			} else if strings.HasPrefix(line, "- ") {
				cleanLine = strings.TrimSpace(strings.TrimPrefix(line, "- "))
				hasDeltaPrefix = true
			} else if strings.HasPrefix(line, "+###") || strings.HasPrefix(line, "+##") || strings.HasPrefix(line, "+#") {
				cleanLine = strings.TrimSpace(strings.TrimPrefix(line, "+"))
				hasDeltaPrefix = true
			} else if strings.HasPrefix(line, "-###") || strings.HasPrefix(line, "-##") || strings.HasPrefix(line, "-#") {
				cleanLine = strings.TrimSpace(strings.TrimPrefix(line, "-"))
				hasDeltaPrefix = true
			}
		}

		// Check Top-Level Title Header
		if strings.HasPrefix(cleanLine, "# ") {
			titleText := strings.TrimSpace(strings.TrimPrefix(cleanLine, "# "))
			if isDelta {
				if strings.HasPrefix(titleText, "Capability Specification Delta:") {
					hasTitle = true
				}
			} else {
				if strings.HasPrefix(titleText, "Capability Specification:") {
					hasTitle = true
				}
			}
			continue
		}

		// Check Requirements Section Header
		if strings.HasPrefix(cleanLine, "## Requirements") {
			hasRequirementsSection = true
			continue
		}

		// Check Requirement Headers
		if strings.HasPrefix(cleanLine, "### ") {
			if inRequirement && !requirementHasModal {
				errors = append(errors, ValidationError{
					File:    filename,
					Line:    requirementLine,
					Message: "Requirement description must contain a normative keyword (SHALL, MUST, SHOULD, or MAY)",
					Rule:    "spec/normative-keyword",
				})
			}

			inRequirement = true
			inScenario = false
			requirementHasModal = false
			requirementLine = lineNum

			reqHeader := strings.TrimSpace(strings.TrimPrefix(cleanLine, "### "))
			if isDelta {
				if !strings.HasPrefix(reqHeader, "+ Requirement:") && !strings.HasPrefix(reqHeader, "- Requirement:") {
					errors = append(errors, ValidationError{
						File:    filename,
						Line:    lineNum,
						Message: "Spec Delta requirement header must start with '+ Requirement:' or '- Requirement:'",
						Rule:    "spec-delta/prefix",
					})
				}
			} else {
				if !strings.HasPrefix(reqHeader, "Requirement:") {
					errors = append(errors, ValidationError{
						File:    filename,
						Line:    lineNum,
						Message: "Living spec requirement header must start with 'Requirement:'",
						Rule:    "spec/requirement-header",
					})
				}
			}
			continue
		}

		// Check Scenario Headers
		if strings.HasPrefix(cleanLine, "#### ") {
			inScenario = true
			scenHeader := strings.TrimSpace(strings.TrimPrefix(cleanLine, "#### "))
			if isDelta {
				if !strings.HasPrefix(scenHeader, "+ Scenario:") && !strings.HasPrefix(scenHeader, "- Scenario:") {
					errors = append(errors, ValidationError{
						File:    filename,
						Line:    lineNum,
						Message: "Spec Delta scenario header must start with '+ Scenario:' or '- Scenario:'",
						Rule:    "spec-delta/prefix",
					})
				}
			} else {
				if !strings.HasPrefix(scenHeader, "Scenario:") {
					errors = append(errors, ValidationError{
						File:    filename,
						Line:    lineNum,
						Message: "Living spec scenario header must start with 'Scenario:'",
						Rule:    "spec/scenario-header",
					})
				}
			}
			continue
		}

		if isDelta && inRequirement && !hasDeltaPrefix && !strings.HasPrefix(cleanLine, "#") {
			errors = append(errors, ValidationError{
				File:    filename,
				Line:    lineNum,
				Message: "Line within Spec Delta must be prefixed with '+' (addition) or '-' (deletion)",
				Rule:    "spec-delta/prefix",
			})
		}

		if inScenario && strings.HasPrefix(cleanLine, "- ") {
			stepText := strings.TrimSpace(strings.TrimPrefix(cleanLine, "- "))
			validKeywords := []string{"GIVEN", "WHEN", "THEN", "AND"}
			hasValidKeyword := false
			for _, kw := range validKeywords {
				if strings.HasPrefix(strings.ToUpper(stepText), kw) {
					hasValidKeyword = true
					break
				}
			}
			if !hasValidKeyword {
				errors = append(errors, ValidationError{
					File:    filename,
					Line:    lineNum,
					Message: "Scenario step must start with GIVEN, WHEN, THEN, or AND",
					Rule:    "spec/scenario-keyword",
				})
			}
		}

		// Check Normative Keyword in Requirement Text
		if inRequirement && !inScenario {
			upper := strings.ToUpper(cleanLine)
			if strings.Contains(upper, "SHALL") || strings.Contains(upper, "MUST") || strings.Contains(upper, "SHOULD") || strings.Contains(upper, "MAY") {
				requirementHasModal = true
			}
		}
	}

	if inRequirement && !requirementHasModal {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    requirementLine,
			Message: "Requirement description must contain a normative keyword (SHALL, MUST, SHOULD, or MAY)",
			Rule:    "spec/normative-keyword",
		})
	}

	if !hasTitle {
		expectedTitle := "# Capability Specification: <Title>"
		if isDelta {
			expectedTitle = "# Capability Specification Delta: <Title>"
		}
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: fmt.Sprintf("Missing or invalid document title. Expected: '%s'", expectedTitle),
			Rule:    "spec/title-header",
		})
	}

	if !hasRequirementsSection {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: "Missing '## Requirements' section header",
			Rule:    "spec/requirements-section",
		})
	}

	return errors
}
