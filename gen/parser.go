package gen

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/net/html"
)

func parseNode(node *html.Node) *ElementNode {
	res := &ElementNode{}
	res.Children = make([]*ElementNode, 0)
	res.Type = node.Data

	for _, attr := range node.Attr {
		switch attr.Key {
		case "colspan":
			res.Colspan, _ = strconv.Atoi(attr.Val)
		case "namespace":
			res.Namespace = attr.Val
		case "background-color":
			if res.Background == nil {
				res.Background = &Background{}
			}
			res.Background.Color = attr.Val
		case "text-align":
			if res.Text == nil {
				res.Text = &Text{}
			}
			res.Text.Align = attr.Val
		case "font-size":
			if res.Font == nil {
				res.Font = &Font{}
			}
			res.Font.Size, _ = strconv.Atoi(attr.Val)
		case "font-color":
			if res.Font == nil {
				res.Font = &Font{}
			}
			res.Font.Color = attr.Val
		}
	}
	return res
}

func parse(filePath string) *ElementNode {
	file, _ := os.ReadFile(filePath)
	raw, _ := html.Parse(bytes.NewReader(file))

	var process func(*html.Node, *ElementNode)
	var doc *ElementNode

	process = func(node *html.Node, parentElement *ElementNode) {
		if node.Type == html.TextNode && parentElement != nil {
			parentElement.Data = node.Data
		} else {
			var ele *ElementNode
			if node.Type == html.ElementNode {
				ele = parseNode(node)
				fmt.Printf("Parse node %s\n", ele.Type)
				if ele.Type == "document" {
					doc = ele
				} else if parentElement != nil {
					parentElement.Children = append(parentElement.Children, ele)
				}
			}
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				process(child, ele)
			}
		}

	}

	process(raw, nil)
	return doc
}
