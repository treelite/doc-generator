package gen

import (
	"bytes"
	"io"
	"os"
	"text/template"
)

func createTemplate(filePath string, data interface{}) io.Reader {
	file, _ := os.ReadFile(filePath)

	tmpl, _ := template.New("test").Parse(string(file))

	var b bytes.Buffer
	tmpl.Execute(&b, data)

	return bytes.NewReader(b.Bytes())
}
