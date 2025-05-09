package recursive_parser

import (
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Parser struct {
	path string
}

func New(path string) *Parser {
	return &Parser{path}
}

func (p *Parser) Walk(root *template.Template) error {
	return filepath.WalkDir(p.path, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		var temp *template.Template = root.New(entry.Name())

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		content, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		if _, err := temp.Parse(string(content)); err != nil {
			return err
		}

		return nil
	})
}
