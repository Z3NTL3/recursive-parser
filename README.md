# recursive-parser

Extends [html/template](https://pkg.go.dev/html/template) by adding the ability to recursively parse dirs and files that are deeply nested within each other. In contrast to the default parsers: they cannot work with views nested in deep folder hierarchies.

### Usage
```go
// provide your views folder
p := recursive_parser.New(path.Join(cwd, "views"))

// name the root template
temp := template.New("views")

// walk recursively and use temp as root template to build upon
if err := p.Walk(temp); err != nil {
    t.Fatal(err)
}

var out bytes.Buffer // where to write

// execute `hello.html`, injecting it with data props and writing output to `out`
if err := temp.ExecuteTemplate(&out, "hello.html", map[string]string{"title": "yolo"}); err != nil {
    t.Fatal(err)
}

// flush the output as string
fmt.Printf("executed temp:\n%s\n", out.String())
```