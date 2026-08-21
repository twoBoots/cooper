package validator

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	titleRegex      = regexp.MustCompile(`(?i)^#\s+(Capability\s+Spec(ification)?(\s+Delta)?|Spec\s+Delta):\s*.+`)
	sectionReqRegex = regexp.MustCompile(`(?i)^##\s+.*Requirements.*`)
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
		if isDelta {
			if strings.HasPrefix(line, "+ ") {
				cleanLine = strings.TrimSpace(strings.TrimPrefix(line, "+ "))
			} else if strings.HasPrefix(line, "- ") {
				cleanLine = strings.TrimSpace(strings.TrimPrefix(line, "- "))
			} else if strings.HasPrefix(line, "+###") || strings.HasPrefix(line, "+##") || strings.HasPrefix(line, "+#") {
				cleanLine = strings.TrimSpace(strings.TrimPrefix(line, "+"))
			} else if strings.HasPrefix(line, "-###") || strings.HasPrefix(line, "-##") || strings.HasPrefix(line, "-#") {
				cleanLine = strings.TrimSpace(strings.TrimPrefix(line, "-"))
			}
		}

		// Check Top-Level Title Header
		if strings.HasPrefix(cleanLine, "# ") {
			if titleRegex.MatchString(cleanLine) {
				hasTitle = true
			}
			continue
		}

		// Check Requirements Section Header
		if strings.HasPrefix(cleanLine, "## ") {
			if sectionReqRegex.MatchString(cleanLine) {
				hasRequirementsSection = true
			}
			continue
		}

		// Check Direct Scenario Step without Requirement Header (e.g. + GIVEN / - GIVEN)
		if hasRequirementsSection && (strings.HasPrefix(cleanLine, "GIVEN ") || strings.HasPrefix(cleanLine, "- GIVEN ")) {
			inScenario = true
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
			if !strings.HasPrefix(reqHeader, "Requirement:") && !strings.HasPrefix(reqHeader, "+ Requirement:") && !strings.HasPrefix(reqHeader, "- Requirement:") {
				errors = append(errors, ValidationError{
					File:    filename,
					Line:    lineNum,
					Message: "Requirement header must start with 'Requirement:' or '+/- Requirement:'",
					Rule:    "spec/requirement-header",
				})
			}
			continue
		}

		// Check Scenario Headers
		if strings.HasPrefix(cleanLine, "#### ") {
			inScenario = true
			scenHeader := strings.TrimSpace(strings.TrimPrefix(cleanLine, "#### "))
			if !strings.HasPrefix(scenHeader, "Scenario:") && !strings.HasPrefix(scenHeader, "+ Scenario:") && !strings.HasPrefix(scenHeader, "- Scenario:") {
				errors = append(errors, ValidationError{
					File:    filename,
					Line:    lineNum,
					Message: "Scenario header must start with 'Scenario:' or '+/- Scenario:'",
					Rule:    "spec/scenario-header",
				})
			}
			continue
		}

		stepCandidate := cleanLine
		if strings.HasPrefix(stepCandidate, "- ") {
			stepCandidate = strings.TrimSpace(strings.TrimPrefix(stepCandidate, "- "))
		}

		if inScenario && (strings.HasPrefix(cleanLine, "- ") || strings.HasPrefix(cleanLine, "GIVEN ") || strings.HasPrefix(cleanLine, "WHEN ") || strings.HasPrefix(cleanLine, "THEN ") || strings.HasPrefix(cleanLine, "AND ")) {
			validKeywords := []string{"GIVEN", "WHEN", "THEN", "AND"}
			hasValidKeyword := false
			for _, kw := range validKeywords {
				if strings.HasPrefix(strings.ToUpper(stepCandidate), kw) {
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
		if inRequirement || inScenario {
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
		expectedTitle := "# Capability Specification: <Title> or # Capability Specification Delta: <Title>"
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
