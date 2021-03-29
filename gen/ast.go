package gen

type Font struct {
	Size  int
	Color string
}

type Text struct {
	Align string
}

type Background struct {
	Color string
}

type ElementNode struct {
	Type       string
	Data       string
	Font       *Font
	Text       *Text
	Background *Background
	Children   []*ElementNode

	Namespace string
	Colspan   int
}
