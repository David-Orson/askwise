package domain

import (
	"errors"
	"path/filepath"
	"strings"
)

var allowedExtensions = map[string]bool{
	".pdf":  true,
	".txt":  true,
	".docx": true,
	".md":   true,
}

func ValidateFileName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return errors.New("fileName cannot be empty")
	}

	if strings.ContainsAny(name, `/\`) {
		return errors.New("fileName cannot contain slashes")
	}

	if len(name) > 255 {
		return errors.New("fileName is too long")
	}

	ext := strings.ToLower(filepath.Ext(name))
	if !allowedExtensions[ext] {
		return errors.New("fileName must have a supported extension (.pdf, .txt, .docx)")
	}

	return nil
}
