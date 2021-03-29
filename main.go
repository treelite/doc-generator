package main

import (
	"io/ioutil"

	"doc-generator/gen"
)

type Item struct {
	Question string
	BgColor  string
	Answer   string
}

type Data struct {
	Items []Item
}

func main() {
	data := Data{
		Items: []Item{
			{
				Question: "Question 1",
				BgColor:  "#CCC",
				Answer:   "Yes",
			},
			{
				Question: "Question 2",
				BgColor:  "#F00",
				Answer:   "No",
			},
		},
	}

	doc := gen.Generate("./pdf.xml", data)

	ioutil.WriteFile("output.pdf", doc, 0644)
}
