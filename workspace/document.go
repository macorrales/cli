package workspace

import "path/filepath"

// Document is a file in a directory.
type Document struct {
	Root     string
	Filepath string
}

// NewDocument creates a document from a filepath.
// The root is typically the root of the exercise, and
// path is the relative path to the file within the root directory.
func NewDocument(root, path string) Document {
	return Document{
		Root:     root,
		Filepath: path,
	}
}

// Path is the normalized path.
// It uses forward slashes regardless of the operating system.
func (doc Document) Path() (string, error) {
	path, err := filepath.Rel(doc.Root, doc.Filepath)
	if err != nil {
		return "", err
	}
	return filepath.ToSlash(path), nil
}
