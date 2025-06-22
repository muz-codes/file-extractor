package util

import (
	"file-extractor/constant"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

func CondenseNewlines(s string) string {
	// changes consecutive \n and spaces before, after or between them to single \n.
	re2 := regexp.MustCompile(`[\s]*\n[\s]*`)
	return re2.ReplaceAllString(s, "\n")
}

func CondenseSpaces(s string) string {
	re := regexp.MustCompile(`\s{2,}`)
	return re.ReplaceAllString(s, " ")
}

func IsValidExt(path string, acceptedExt []string, isUrl bool) bool {
	// updating path to parseUrl.path in case of url
	if isUrl {
		parsedUrl, err := url.Parse(path)
		if err != nil {
			fmt.Printf("error is %s\n", err.Error())
			return false
		}
		fmt.Printf("path is %s\n", parsedUrl.Path)
		path = parsedUrl.Path
	}
	fileExt := filepath.Ext(path)
	for _, singleImageType := range acceptedExt {
		if strings.ToLower(fileExt) == "."+singleImageType {
			return true
		}
	}
	return false
}

func GetDocsExt(Filename string) (string, error) {
	parsedUrl, err := url.Parse(Filename)
	if err != nil {
		return "", err
	}
	return filepath.Ext(parsedUrl.Path), nil
}

func ReplaceContentGtWithTemp(xml string) string {
	inTag := false
	var result strings.Builder

	for _, ch := range xml {
		switch ch {
		case '<':
			inTag = true
			result.WriteRune(ch)
		case '>':
			inTag = false
			result.WriteRune(ch)
		default:
			if ch == '>' && !inTag {
				result.WriteString("#tempArr")
			} else {
				result.WriteRune(ch)
			}
		}
	}
	return result.String()
}

func replaceSpecialToOriginal(s string) string {
	s = strings.ReplaceAll(s, constant.Gt, ">")
	s = strings.ReplaceAll(s, constant.Lt, "<")
	s = strings.ReplaceAll(s, constant.SingleQuote, `'`)
	s = strings.ReplaceAll(s, constant.DoubleQuote, `"`)
	s = strings.ReplaceAll(s, constant.Ampersand, "&")
	return s
}

func sanitizeContent(content string) string {
	// Define your sanitization logic here.
	// For example, replace problematic characters:
	content = strings.ReplaceAll(content, "&", constant.Ampersand)
	content = strings.ReplaceAll(content, ">", constant.Gt)
	content = strings.ReplaceAll(content, "<", constant.Lt)
	content = strings.ReplaceAll(content, `'`, constant.SingleQuote)
	content = strings.ReplaceAll(content, `"`, constant.DoubleQuote)
	content = strings.ReplaceAll(content, `\n`, "")
	return content
}

func AttachUrlScheme(inputUrl string, scheme string) (string, error) {
	// if url doesn't have anything, add https
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(inputUrl, "http://") && !strings.HasPrefix(inputUrl, "https://") && parsedUrl.Scheme == "" {
		// Check with "https://" prefix
		inputUrl = scheme + inputUrl
	} else if strings.HasPrefix(inputUrl, "http://") && scheme == "https://" {
		inputUrl = strings.Replace(inputUrl, "http://", scheme, 1)
	} else if strings.HasPrefix(inputUrl, "https://") && scheme == "http://" {
		inputUrl = strings.Replace(inputUrl, "https://", scheme, 1)
	}
	return inputUrl, nil
}
