package recursive_parser_test

import (
	"bytes"
	"html/template"
	"os"
	"path"
	"testing"

	recursive_parser "github.com/Z3NTL3/recursive-parser"
)

func TestRecursiveTempParser(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	p := recursive_parser.New(path.Join(cwd, "data")) // views dir
	temp := template.New("views")

	if err := p.Walk(temp); err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	if err := temp.ExecuteTemplate(&out, "hello.html", map[string]string{"title": "yolo"}); err != nil {
		t.Fatal(err)
	}

	t.Logf("executed temp:\n%s\n", out.String())
}
