package recursive_parser_test

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"

	recursive_parser "github.com/Z3NTL3/recursive-parser"
)

func TestRecursiveTempParser(t *testing.T) {
	p := recursive_parser.New("data") // views dir
	temp := template.New("views")

	if err := p.Walk(temp); err != nil {
		t.Fatal(err)
	}

	for _, t := range temp.Templates() {
		fmt.Println(t.Name())
	}
	var out bytes.Buffer
	if err := temp.ExecuteTemplate(&out, "data\\hello.html", map[string]string{"title": "yolo"}); err != nil {
		t.Fatal(err)
	}

	t.Logf("executed temp:\n%s\n", out.String())
}
