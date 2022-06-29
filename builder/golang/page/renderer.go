package page

import (
	"bufio"
	"bytes"
	"text/template"
)

type Renderer interface {
	Render(page Page) string
}

type HTMLRenderer struct{}

func (r *HTMLRenderer) Render(page Page) string {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{.Title}}</title>
</head>
<body>
	<h1>{{.Title}}</h1>
	{{range .ContentSections}}
	<div class="section">
		{{.}}
	</div>
	{{end}}
</body>
</html>
`
	// use text.template to avoid escaping
	tmpl := template.New("page")
	if _, err := tmpl.Parse(html); err != nil {
		return ""
	}

	var b bytes.Buffer
	byteWriter := bufio.NewWriter(&b)
	err := tmpl.Execute(byteWriter, page)
	if err != nil {
		return ""
	}
	byteWriter.Flush()

	return b.String()
}

type MarkdownRenderer struct{}

func (r *MarkdownRenderer) Render(page Page) string {
	md := `
# {{.Title}}

{{range .ContentSections}}
{{.}}
{{end}}
`
	// use text.template to avoid escaping
	tmpl := template.New("page")
	if _, err := tmpl.Parse(md); err != nil {
		return ""
	}

	var b bytes.Buffer
	byteWriter := bufio.NewWriter(&b)
	err := tmpl.Execute(byteWriter, page)
	if err != nil {
		return ""
	}
	byteWriter.Flush()

	return b.String()
}
