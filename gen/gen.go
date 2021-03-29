package gen

func Generate(filePath string, data interface{}) []byte {
	template := createTemplate(filePath, data)
	doc := parse(template)
	return render(doc)
}
