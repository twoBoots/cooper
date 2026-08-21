package scaffold

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
)

//go:embed assets/* assets/templates/* assets/skills/* assets/skills/*/* assets/templates/code_styleguides/*
var embeddedAssets embed.FS

// GetEmbeddedFS returns the underlying embed.FS.
func GetEmbeddedFS() embed.FS {
	return embeddedAssets
}

// ListSkills returns the list of all available embedded skill names.
func ListSkills() ([]string, error) {
	entries, err := embeddedAssets.ReadDir("assets/skills")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded skills directory: %w", err)
	}

	var skills []string
	for _, entry := range entries {
		if entry.IsDir() {
			skills = append(skills, entry.Name())
		}
	}
	return skills, nil
}

// GetSkill returns the contents of a given skill's SKILL.md.
func GetSkill(skillName string) ([]byte, error) {
	filePath := path.Join("assets/skills", skillName, "SKILL.md")
	data, err := embeddedAssets.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("skill '%s' not found: %w", skillName, err)
	}
	return data, nil
}

// GetTemplate returns the contents of a template file from assets/templates.
func GetTemplate(templatePath string) ([]byte, error) {
	filePath := path.Join("assets/templates", templatePath)
	data, err := embeddedAssets.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("template '%s' not found: %w", templatePath, err)
	}
	return data, nil
}

// GetAgentsTemplate returns the contents of AGENTS.template.md.
func GetAgentsTemplate() ([]byte, error) {
	data, err := embeddedAssets.ReadFile("assets/AGENTS.template.md")
	if err != nil {
		return nil, fmt.Errorf("AGENTS template not found: %w", err)
	}
	return data, nil
}

// WalkAssets walks through all embedded assets rooted at assets/.
func WalkAssets(fn func(path string, d fs.DirEntry, err error) error) error {
	return fs.WalkDir(embeddedAssets, "assets", fn)
}
