package textExtract

import (
	"file-extractor/constant"
	"file-extractor/util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"strings"
)

func traverseHTML(node *goquery.Selection) (string, error) {
	var htmlString string
	var htmlErr error

	// Check for direct text in the current node
	//var previousTempNode html.NodeType

	node.Contents().Each(func(i int, contentSelection *goquery.Selection) {
		if htmlErr != nil {
			return
		}

		if contentSelection.Nodes[0].Type == html.TextNode {
			textContent := strings.Trim(strings.TrimSpace(contentSelection.Text()), "\n")
			if textContent != "" {
				//if previousTempNode == html.ElementNode {
				//	htmlString += "\n"
				//}
				//htmlString += condenseSpaces(strings.ReplaceAll(textContent, "\n", " ") + " ")
				htmlString += util.CondenseSpaces(strings.ReplaceAll(textContent, "\n", " ") + " ")
			}
			//previousTempNode = contentSelection.Nodes[0].Type
			// old
		} else if contentSelection.Nodes[0].Type == html.ElementNode {
			if isIgnoredElement(contentSelection) {
				return
			}
			if isBlock(contentSelection) || !isInLine(contentSelection) {
				//if isBlock(contentSelection) {
				// If block element, do a DFS on that block
				dfsContent, dfsErr := traverseHTML(contentSelection)
				if dfsErr != nil {
					htmlErr = dfsErr
					return
				}
				//dfs := strings.Trim(strings.TrimSpace(dfsContent), "\n") + "\n"
				dfs := strings.TrimSpace(strings.Trim(dfsContent, "\n"))
				if dfs != "" {
					htmlString += "\n" + dfs + "\n"
				}
			} else if isInLine(contentSelection) {
				// If inline element, directly add its text
				//previousTempNode = contentSelection.Nodes[0].Type

				textContent := util.CondenseSpaces(strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(contentSelection.Text(), "\n", " "), "\t", "")))
				if textContent != "" {
					htmlString += textContent + " "
				}
			} else {
				htmlErr = fmt.Errorf("unexpected node type: %v, node name: %v", contentSelection.Nodes[0].Type, goquery.NodeName(contentSelection))
				return
			}
		}
	})

	return strings.TrimSpace(htmlString), htmlErr
}

func isValidHtmlNode(contentSelection *goquery.Selection) bool {
	htmlNodeTypes := []html.NodeType{html.ErrorNode, html.TextNode, html.DocumentNode, html.ElementNode, html.CommentNode, html.DoctypeNode, html.RawNode, 7}
	for _, singleNode := range htmlNodeTypes {
		if contentSelection.Nodes[0].Type == singleNode {
			return true
		}
	}
	return false
}

func isIgnoredElement(contentSelection *goquery.Selection) bool {
	if _, exists := constant.IgnoredElements[goquery.NodeName(contentSelection)]; exists {
		return true
	}
	return false
}

func isBlock(contentSelection *goquery.Selection) bool {
	if _, exists := constant.BlockLevelElements[goquery.NodeName(contentSelection)]; exists {
		return true
	}
	if _, exists := constant.DeprecatedBlockLevelElements[goquery.NodeName(contentSelection)]; exists {
		return true
	}
	return false
}

func isInLine(contentSelection *goquery.Selection) bool {
	if contentSelection.Is(constant.InlineElements) || contentSelection.Is(constant.DeprecatedInlineElements) {
		return true
	}
	return false
}

func ExtractHtml(htmlContent string) (string, error) {
	reader := strings.NewReader(htmlContent)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	body := doc.Find("body")
	htmlString, err := traverseHTML(body)
	if err != nil {
		return "", err
	}
	htmlString = util.CondenseNewlines(htmlString)

	return htmlString, nil
}
