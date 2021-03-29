package gen

func Gen(filePath string) []byte {
	doc := parse(filePath)
	return render(doc)
}
