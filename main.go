package main

import (
	"io/ioutil"

	"doc-generator/gen"
)

func main() {
	doc := gen.Gen("./pdf.xml")

	ioutil.WriteFile("output.pdf", doc, 0644)
}
