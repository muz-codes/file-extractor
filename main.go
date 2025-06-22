package main

import (
	"bytes"
	"code.sajari.com/docconv"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"file-extractor/constant"
	e "file-extractor/errors"
	"file-extractor/middleware"
	"file-extractor/scrapper"
	"file-extractor/textExtract"
	"file-extractor/util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	logware "github.com/apyhub/logging-middleware"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	_ "github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"golang.org/x/net/html"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Isurl bool

var (
	//result  string
	//err     error
	browser *rod.Browser
)

type RequestUrl struct {
	Url            string `json:"url"`
	ParagraphBreak bool   `json:"preserve_paragraphs"`
}

func main() {
	//Define endpoint
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if err := StartBrowser(); err != nil {
		log.Fatal("error while starting Browser", err.Error())
	}
	router := mux.NewRouter()
	router.HandleFunc("/hello", healthCheck).Methods("GET")
	router.HandleFunc("/hello/version", getVersion).Methods("GET")

	// extractor-webpage-text
	extractorWebpageTextRouter := router.PathPrefix("/").Subrouter()
	extractorWebpageTextRouter.Use(logware.Logging("extractor-webpage-text"))
	extractorWebpageTextRouter.HandleFunc("/extractor/scrapper", Scrapper).Methods("GET")
	extractorWebpageTextRouter.HandleFunc("/extractor/scrapper/split", Scrapper).Methods("GET")

	// extractor-pdf-text
	extractorPdfTextRouter := router.PathPrefix("/").Subrouter()
	extractorPdfTextRouter.Use(logware.Logging("extractor-pdf-text"))
	extractorPdfTextRouter.HandleFunc("/extractor/multipart/pdf", Isurl(false).Pdf).Methods("POST")
	extractorPdfTextRouter.HandleFunc("/extractor/url/pdf", Isurl(true).Pdf).Methods("POST")

	// extractor-word-text
	extractorWordTextRouter := router.PathPrefix("/").Subrouter()
	extractorWordTextRouter.Use(logware.Logging("extractor-word-text"))
	extractorWordTextRouter.HandleFunc("/extractor/multipart/word", Isurl(false).Word).Methods("POST")
	extractorWordTextRouter.HandleFunc("/extractor/url/word", Isurl(true).Word).Methods("POST")

	// following not used in any util
	extractorDocumentTextRouter := router.PathPrefix("/").Subrouter()
	extractorDocumentTextRouter.HandleFunc("/extractor/multipart/document", Isurl(false).Document).Methods("POST")
	extractorDocumentTextRouter.HandleFunc("/extractor/url/document", Isurl(true).Document).Methods("POST")

	router.HandleFunc("/extractor/multipart/xml", Isurl(false).Xml).Methods("POST")
	router.HandleFunc("/extractor/url/xml", Isurl(true).Xml).Methods("POST")

	router.HandleFunc("/extractor/multipart/html", Isurl(false).Html).Methods("POST")
	router.HandleFunc("/extractor/url/html", Isurl(true).Html).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", middleware.CORS(router)))
}

// To check service health
func healthCheck(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./test_files/sample.pdf")
	if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
		return
	}
	if _, _, err = docconv.ConvertPDF(file); err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
		return
	}
	respondWithJSON(w, map[string]interface{}{"data": "Apy Day!"})
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	build := "undefined"
	if os.Getenv("APY_BUILD") != "" {
		build = os.Getenv("APY_BUILD")
	}
	respondWithJSON(w, map[string]interface{}{"data": build})
}

func Scrapper(w http.ResponseWriter, req *http.Request) {
	urlInput := strings.Trim(req.URL.Query().Get("url"), " ")
	if urlInput == "" && ErrorCheck(w, 400, e.ErrURLMissing, e.Errors[e.ErrURLMissing]) {
		return
	}
	//_, urlInput, statusCode, code := urlParse(urlInput)
	//if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
	//	return
	//}
	ScrappingProcess(w, req, urlInput)
}

// old
//func ScrappingProcess(w http.ResponseWriter, req *http.Request, url string) {
//	p := scrapper.StrictPolicy()
//	c := colly.NewCollector()
//	var scrape string
//
//	c.OnRequest(func(r *colly.Request) {
//		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
//	})
//	c.OnHTML("html", func(edom *colly.HTMLElement) {
//		scrape, err = edom.DOM.Html()
//		if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue]) {
//			return
//		}
//	})
//	c.Visit(url)
//	scrap := p.Sanitize(scrape)
//	if strings.Contains(req.URL.Path, "/split") {
//		respondWithJSON(w, map[string]interface{}{"data": scrap})
//		return
//	}
//	respondWithJSON(w, map[string]interface{}{"data": strings.Join(scrap, " ")})
//}

func StartBrowser() error {
	// *** custom chromium ***
	path := "/usr/bin/chromium"
	newLauncher := launcher.New().Headless(true).Bin(path).Set("no-sandbox")
	rodUrl, err := newLauncher.Launch()
	if err != nil {
		return err
	}
	browser = rod.New().ControlURL(rodUrl)
	// ** end custom chromium ***

	//browser := rod.New()
	if err = browser.Connect(); err != nil {
		return err
	}
	return nil
}

func GetHtmlFromSite(url string) (string, error) {
	domStableTimeout := time.Duration(3)

	//// *** custom chromium ***
	//path := "/usr/bin/chromium"
	//u := launcher.New().Headless(true).Bin(path)
	//fmt.Printf("flags are : %v", u.Flags)
	//a, err := u.Launch()
	//if err != nil {
	//	log.Output(1, fmt.Sprintf("error in ScrappingProcess: %v", err.Error()))
	//	ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue])
	//	return
	//}
	//browser := rod.New().ControlURL(a)
	//// ** end custom chromium ***
	//
	////browser := rod.New()
	//if err = browser.Connect(); err != nil {
	//	log.Output(1, fmt.Sprintf("error in ScrappingProcess: %v", err.Error()))
	//	ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue])
	//	return
	//}
	//defer browser.Close()
	tryNavigation := func(url string) (*rod.Page, error) {
		refinedUrl, err := util.AttachUrlScheme(url, "https://")
		if err != nil {
			log.Output(1, fmt.Sprintf("error in GetHtmlFromSite: %v", err.Error()))
			return nil, err
		}
		page, err := browser.Page(proto.TargetCreateTarget{URL: refinedUrl})
		navErr := rod.ErrNavigation{Reason: "http required instead of https"}
		if err != nil && (navErr.Is(err) || strings.Contains(err.Error(), "Cannot navigate to invalid URL")) {
			refinedUrl, err = util.AttachUrlScheme(url, "http://")
			if err != nil {
				log.Output(1, fmt.Sprintf("error in GetHtmlFromSite: %v", err.Error()))
				return nil, err
			}
			page, err = browser.Page(proto.TargetCreateTarget{URL: refinedUrl})
		}
		return page, err
	}
	page, err := tryNavigation(url)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in GetHtmlFromSite: %v", err.Error()))
		return "", err
	}
	defer page.Close()

	if err = page.WaitLoad(); err != nil {
		log.Output(1, fmt.Sprintf("error in GetHtmlFromSite: %v", err.Error()))
		return "", err
	}

	// setting the timeout in case if the DOM is never stable because of continuous async calls
	// newPage has the timeout set
	newPage := page.Timeout(domStableTimeout * time.Second)

	if err := newPage.WaitDOMStable(1*time.Second, 5); err != nil {
		if err != context.DeadlineExceeded {
			log.Output(1, fmt.Sprintf("error in GetHtmlFromSite: %v", err.Error()))
			return "", err
		}
	}

	// timeout is not needed after waiting for DOM Stability.
	newPage = newPage.CancelTimeout()

	// Get the HTML content
	htmlContent, err := newPage.HTML()
	if err != nil {
		log.Output(1, fmt.Sprintf("error in GetHtmlFromSite: %v", err.Error()))
		return "", err
	}
	return htmlContent, nil
}

// using rod
func ScrappingProcess(w http.ResponseWriter, req *http.Request, url string) {
	// Launch the browser and create a new page
	htmlContent, err := GetHtmlFromSite(url)
	if err != nil {
		navErr := rod.ErrNavigation{Reason: e.Errors[e.ErrURLMissing].Error()}
		if navErr.Is(err) || strings.Contains(err.Error(), "Cannot navigate to invalid URL") {
			log.Output(1, fmt.Sprintf("error in ScrappingProcess: %v", err.Error()))
			ErrorCheck(w, 400, e.ErrURLMissing, e.Errors[e.ErrURLMissing], err)
			return
		}
		log.Output(1, fmt.Sprintf("error in ScrappingProcess: %v", err.Error()))
		ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err)
		return
	}
	p := scrapper.StrictPolicy()
	scrap := p.Sanitize(htmlContent)
	if strings.Contains(req.URL.Path, "/split") {
		respondWithJSON(w, map[string]interface{}{"data": scrap})
		return
	}
	respondWithJSON(w, map[string]interface{}{"data": strings.Join(scrap, " ")})
}

func (url Isurl) Pdf(w http.ResponseWriter, req *http.Request) {
	var paragraphBreak bool
	var err error
	var result string
	var pdfAcceptedInputs = []string{"pdf"}

	if url {
		var urlRequest RequestUrl
		err = json.NewDecoder(req.Body).Decode(&urlRequest)
		if err != nil && ErrorCheck(w, 400, e.ErrJsonInvalid, e.Errors[e.ErrJsonInvalid], err) {
			return
		}
		paragraphBreak = urlRequest.ParagraphBreak
		res, path, statusCode, code := getInputsUrl(urlRequest)
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}
		urlRequest.Url = path
		if !util.IsValidExt(urlRequest.Url, pdfAcceptedInputs, true) && ErrorCheck(w, 400, e.ErrURLMissing, e.Errors[e.ErrURLMissing]) {
			return
		}
		result, _, err = textExtract.ExtractPdf(res.Body)
		if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
			return
		}
	} else {
		file, fileHeader, statusCode, code := getInputFile(w, req)
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}
		if strings.TrimSpace(req.FormValue("preserve_paragraphs")) == "true" {
			paragraphBreak = true
		}
		defer file.Close()
		if !util.IsValidExt(fileHeader.Filename, pdfAcceptedInputs, false) {
			ErrorCheck(w, 400, e.ErrFileMissing, e.Errors[e.ErrFileMissing])
			return
		}
		result, _, err = textExtract.ExtractPdf(file)
		if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
			return
		}
	}
	if paragraphBreak {
		//respondWithJSON(w, map[string]interface{}{"data": strings.Split(result, "\n")})
		respondWithJSON(w, map[string]interface{}{"data": strings.Trim(util.CondenseNewlines(result), "\n")})
		return
	}
	respondWithJSON(w, map[string]interface{}{"data": strings.Replace(util.CondenseNewlines(result), "\n", " ", -1)})
}

// get url or file from image
// azure ai accepts image url or base 64
// if url from user, forward to azure ai
// if image from user, convert to base64 and then forward
// get result id and call the azure ai to get the result

func (url Isurl) Document(w http.ResponseWriter, req *http.Request) {
	var acceptedDocumentInputs = []string{"jpeg", "jpg", "png", "bmp", "tiff", "xlsx", "pptx"}
	var paragraphBreak bool
	var urlRequest RequestUrl
	var extractDocument string
	var err error
	if url {
		err = json.NewDecoder(req.Body).Decode(&urlRequest)
		if err != nil && ErrorCheck(w, 400, e.ErrJsonInvalid, e.Errors[e.ErrJsonInvalid], err) {
			return
		}
		paragraphBreak = urlRequest.ParagraphBreak
		_, path, statusCode, code := getInputsUrl(urlRequest)
		urlRequest.Url = path
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}
		if !util.IsValidExt(urlRequest.Url, acceptedDocumentInputs, true) {
			errorOccurred := errors.New("invalid file extension")
			log.Output(1, fmt.Sprintf("error in Document: %v", errorOccurred.Error()))
			ErrorCheck(w, http.StatusBadRequest, e.ErrURLMissing, e.Errors[e.ErrURLMissing], errorOccurred)
			return
		}
		extractDocument, err = textExtract.ExtractDocument(urlRequest.Url, "url")
		if err != nil {
			log.Output(1, fmt.Sprintf("error in Document: %v", err.Error()))
			ErrorCheck(w, http.StatusInternalServerError, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err)
			return
		}
	} else {
		file, fileHeader, statusCode, code := getInputFile(w, req)
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}
		if strings.TrimSpace(req.FormValue("preserve_paragraphs")) == "true" {
			paragraphBreak = true
		}
		defer file.Close()

		if !util.IsValidExt(fileHeader.Filename, acceptedDocumentInputs, false) {
			errorOccurred := errors.New("invalid file extension")
			log.Output(1, fmt.Sprintf("error in Document: %v", errorOccurred.Error()))
			ErrorCheck(w, http.StatusBadRequest, e.ErrFileMissing, e.Errors[e.ErrFileMissing], errorOccurred)
			return
		}
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			log.Output(1, fmt.Sprintf("error in Document: %v", err.Error()))
			ErrorCheck(w, http.StatusInternalServerError, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err)
			return
		}
		base64String := base64.StdEncoding.EncodeToString(fileBytes)
		extractDocument, err = textExtract.ExtractDocument(base64String, "base64")
		if err != nil {
			log.Output(1, fmt.Sprintf("error in Document: %v", err.Error()))
			ErrorCheck(w, http.StatusInternalServerError, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err)
			return
		}
	}
	if paragraphBreak {
		respondWithJSON(w, map[string]interface{}{"data": extractDocument})
		return
	}
	respondWithJSON(w, map[string]interface{}{"data": strings.ReplaceAll(extractDocument, "\n", " ")})
	return
}

func (url Isurl) Word(w http.ResponseWriter, req *http.Request) {
	var paragraphBreak bool
	var result string
	var err error
	if url {
		var urlRequest RequestUrl
		err = json.NewDecoder(req.Body).Decode(&urlRequest)
		if err != nil && ErrorCheck(w, 400, e.ErrJsonInvalid, e.Errors[e.ErrJsonInvalid], err) {
			return
		}
		paragraphBreak = urlRequest.ParagraphBreak
		res, path, statusCode, code := getInputsUrl(urlRequest)
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}
		fmt.Printf("path is %v\n", path)

		ext, err := util.GetDocsExt(path)
		fmt.Printf("ext is %v\n", filepath.Ext(ext))

		if err != nil {
			log.Output(1, "error in Word: "+err.Error())
			ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err)
			return
		}
		switch ext {
		case ".doc":
			result, _, err = docconv.ConvertDoc(res.Body)
		case ".docx":
			result, _, err = docconv.ConvertDocx(res.Body)
		case "":
			result, _, err = docconv.ConvertDoc(res.Body)
			if err != nil {
				result, _, err = docconv.ConvertDocx(res.Body)
			}
		default:
			ErrorCheck(w, 400, e.ErrURLMissing, e.Errors[e.ErrURLMissing])
			return
		}
		if err != nil {
			if strings.Contains(err.Error(), "not a valid zip file") {
				ErrorCheck(w, 400, e.ErrURLMissing, e.Errors[e.ErrURLMissing], err)
				return
			}
			ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err)
			return
		}
	} else {
		file, fileHeader, statusCode, code := getInputFile(w, req)
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}

		if strings.TrimSpace(req.FormValue("preserve_paragraphs")) == "true" {
			paragraphBreak = true
		}
		if !wordExtChk(fileHeader.Filename) && ErrorCheck(w, 400, e.ErrFileMissing, e.Errors[e.ErrFileMissing]) {
			return
		}

		defer file.Close()
		if filepath.Ext(fileHeader.Filename) == ".doc" {
			result, _, err = docconv.ConvertDoc(file)
		} else if filepath.Ext(fileHeader.Filename) == ".docx" {
			result, _, err = docconv.ConvertDocx(file)
		}
		// check error in word file.
		if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
			return
		}
	}
	if paragraphBreak {
		respondWithJSON(w, map[string]interface{}{"data": strings.Trim(util.CondenseNewlines(result), "\n")})
		//respondWithJSON(w, map[string]interface{}{"data": strings.Split(condenseNewlines(result), "\n")})
		return
	}
	respondWithJSON(w, map[string]interface{}{"data": strings.Replace(strings.Trim(util.CondenseNewlines(result), "\n"), "\n", " ", -1)})
}

//func extractText(sel *goquery.Selection, blockLevelElements map[string]struct{}, inlineElements string, results *[]string) {
//	// Extract direct text of the current selection
//	directText := ""
//
//	sel.Contents().Each(func(i int, contentSelection *goquery.Selection) {
//		if contentSelection.Nodes[0].Type == html.TextNode {
//			textContent := strings.TrimSpace(contentSelection.Text())
//			if textContent != "" {
//				directText += textContent + " "
//			}
//		} else if contentSelection.Is(inlineElements) {
//			textContent := strings.TrimSpace(contentSelection.Text())
//			if textContent != "" {
//				directText += textContent + " "
//			}
//		} else if _, isBlock := blockLevelElements[goquery.NodeName(contentSelection)]; isBlock {
//			// If the current content is a block element, recursively extract its text
//			extractText(contentSelection, blockLevelElements, inlineElements, results)
//		}
//	})
//
//	directText = strings.TrimSpace(directText)
//	if directText != "" {
//		*results = append(*results, directText)
//	}
//}

//func getHtmlString2(content string) (string, error) {
//	reader := strings.NewReader(content)
//
//	doc, err := goquery.NewDocumentFromReader(reader)
//	if err != nil {
//		return "", err
//	}
//
//	var results []string
//
//	// Start the extraction from the root
//	doc.Find("*").Each(func(i int, sel *goquery.Selection) {
//		if _, isBlock := constant.BlockLevelElements[goquery.NodeName(sel)]; isBlock {
//			extractText(sel, constant.BlockLevelElements, constant.InlineElements, &results)
//		}
//	})
//
//	return strings.Join(results, "\n"), nil
//}

// main function
func getHtmlString(content string) (string, error) {
	reader := strings.NewReader(content)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	var results []string

	doc.Find("*").Each(func(i int, sel *goquery.Selection) {
		if _, isBlock := constant.BlockLevelElements[goquery.NodeName(sel)]; isBlock {
			// Extract direct text
			directText := ""
			sel.Contents().Each(func(i int, contentSelection *goquery.Selection) {
				if contentSelection.Nodes[0].Type == html.TextNode {
					textContent := strings.TrimSpace(contentSelection.Text())
					if textContent != "" {
						directText += textContent + " "
					}
				} else if contentSelection.Is(constant.InlineElements) { // Add more inline elements if needed
					textContent := strings.Trim(strings.Trim(strings.TrimSpace(contentSelection.Text()), "\n"), "\t")
					if textContent != "" {
						directText += textContent + " "
					}
				}
			})
			directText = strings.TrimSpace(directText)
			if directText != "" {
				results = append(results, directText)
			}
		}
	})
	return strings.Join(results, "\n"), nil
}

// todo - can be used for html
func validateAndCleanHtml4(content string) (string, error) {
	// Load HTML content into a reader
	reader := strings.NewReader(content)

	// Parse the HTML content with goquery
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	blockLevelElements := map[string]struct{}{
		"p":          {},
		"div":        {},
		"article":    {},
		"section":    {},
		"header":     {},
		"footer":     {},
		"main":       {},
		"aside":      {},
		"h1":         {},
		"h2":         {},
		"h3":         {},
		"h4":         {},
		"h5":         {},
		"h6":         {},
		"ol":         {},
		"ul":         {},
		"li":         {},
		"dl":         {},
		"dt":         {},
		"dd":         {},
		"blockquote": {},
		"figure":     {},
		"figcaption": {},
		"pre":        {},
		"table":      {},
		"caption":    {},
		"thead":      {},
		"tbody":      {},
		"tfoot":      {},
		"colgroup":   {},
		"col":        {},
		"tr":         {},
		"td":         {},
		"th":         {},
		"form":       {},
		"fieldset":   {},
		"legend":     {},
		"label":      {},
		"button":     {},
	}

	var htmlString string

	doc.Find("*").Each(func(index int, item *goquery.Selection) {
		nodeName := goquery.NodeName(item)

		// Check if the current node is a block-level element
		if _, exists := blockLevelElements[nodeName]; exists {
			// Extracting direct text of the current element, which isn't part of any child element
			directText := strings.TrimSpace(item.Contents().FilterFunction(func(i int, s *goquery.Selection) bool {
				return goquery.NodeName(s) != goquery.NodeName(item.Children())
			}).Text())
			if directText != "" {
				htmlString += directText + "\n"
			}

			// Check if it doesn't have block-level children, and add its entire text to the results
			hasBlockChild := false
			item.Children().Each(func(_ int, child *goquery.Selection) {
				if _, exists := blockLevelElements[goquery.NodeName(child)]; exists {
					hasBlockChild = true
				}
			})

			if !hasBlockChild {
				text := strings.TrimSpace(item.Text())
				if text != "" && text != directText { // Ensure we aren't duplicating the direct text
					htmlString += text + "\n"
				}
			}
		}
	})

	return htmlString, nil
}

func validateAndCleanHtml3(content string) (string, error) {
	// Load HTML content into a reader
	reader := strings.NewReader(content)

	// Parse the HTML content with goquery
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Define block-level elements
	blockLevelElements := map[string]struct{}{
		"p":       {},
		"div":     {},
		"article": {},
		"section": {},
		"header":  {},
		"footer":  {},
		"main":    {},
		"aside":   {},
		"h1":      {},
		"h2":      {},
		"h3":      {},
		"h4":      {},
		"h5":      {},
		"h6":      {},
	}

	var htmlString []string

	// Iterate over all nodes
	doc.Find("*").Each(func(index int, item *goquery.Selection) {
		nodeName := goquery.NodeName(item)
		// Check if the current node is a block-level element
		if _, exists := blockLevelElements[nodeName]; exists {
			// Check if it has any block-level children
			hasBlockChild := false
			item.Children().Each(func(_ int, child *goquery.Selection) {
				if _, exists := blockLevelElements[goquery.NodeName(child)]; exists {
					hasBlockChild = true
				}
			})

			// If it doesn't have block-level children, add its text to the results
			if !hasBlockChild {
				text := strings.TrimSpace(item.Text())
				if text != "" {
					htmlString = append(htmlString, text+"\n")
				}
			}
		}
	})
	return strings.Join(htmlString, " "), nil

}

func validateAndCleanHtml2(content string) (string, error) {
	// Load XML content into a reader
	reader := bytes.NewReader([]byte(content))

	// Parse the XML content with goquery
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// List of block-level elements. If a text node is inside any of these tags, we'll add a newline after it.
	blockLevelTags := map[string]bool{
		"p":       true,
		"div":     true,
		"h1":      true,
		"h2":      true,
		"h3":      true,
		"h4":      true,
		"h5":      true,
		"h6":      true,
		"ol":      true,
		"ul":      true,
		"li":      true,
		"article": true,
		"header":  true,
		"footer":  true,
		"section": true,
		// ... add other block-level tags if necessary
	}

	var htmlString []string

	// Traverse the document
	doc.Find("*").Each(func(index int, item *goquery.Selection) {
		nodeName := goquery.NodeName(item)
		// Check if the tag is a block-level element
		if _, exists := blockLevelTags[nodeName]; exists {
			nodeText := item.Text()
			if strings.TrimSpace(nodeText) != "" {
				htmlString = append(htmlString, nodeText+"\n")
			}
		}
	})
	return strings.Join(htmlString, " "), nil

}

func validateAndCleanHtml1(content string) (string, error) {
	// Load XML content into a reader
	reader := bytes.NewReader([]byte(content))

	// Parse the XML content with goquery
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	//var htmlString string
	//// Iterate over all nodes and sanitize text content
	//doc.Find("*").Each(func(index int, item *goquery.Selection) {
	//	// Sanitize node text
	//	nodeText := item.Text()
	//	//item.SetText(sanitizeContent(nodeText) + "\n")
	//	//item.SetText(nodeText + "\n")
	//	htmlString += nodeText + "\n"
	//})
	// Print the sanitized XML
	//sanitizedHtml, _ := doc.Html()
	//fmt.Println("sanitizedXML is")
	//fmt.Println(htmlString)
	//return htmlString, nil

	var htmlString []string

	doc.Find("*").Each(func(index int, item *goquery.Selection) {
		// If it's a <br> tag, skip
		if item.Is("br") {
			return
		}

		nodeText := item.Contents().FilterFunction(func(_ int, sel *goquery.Selection) bool {
			node := sel.Get(0)
			return node.Type == html.TextNode
		}).Text()

		// If the next sibling of this text is <br>, don't add a newline
		if item.Next().Is("br") {
			htmlString = append(htmlString, nodeText)
		} else if strings.TrimSpace(nodeText) != "" {
			htmlString = append(htmlString, nodeText+"\n")
		}
	})

	return strings.Join(htmlString, " "), nil

}

func validateAndCleanXml(input io.Reader) (string, error) {
	var buffer bytes.Buffer
	encoder := xml.NewEncoder(&buffer)
	decoder := xml.NewDecoder(input)

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Output(1, fmt.Sprintf("error in validateAndCleanXml: %s", err.Error()))
			return "", err
		}

		switch v := t.(type) {
		case xml.CharData:
			//cleanedData := strings.ReplaceAll(string(v), `&`, ampersand)
			cleanedData := strings.ReplaceAll(string(v), `>`, constant.Gt)
			//cleanedData = strings.ReplaceAll(cleanedData, `<`, lt)
			//cleanedData = strings.ReplaceAll(cleanedData, `'`, singleQuote)
			//cleanedData = strings.ReplaceAll(cleanedData, `"`, doubleQuote)
			if err := encoder.EncodeToken(xml.CharData(cleanedData)); err != nil {
				log.Output(1, fmt.Sprintf("error in validateAndCleanXml: %s", err.Error()))
				return "", err
			}
		default:
			if err := encoder.EncodeToken(t); err != nil {
				log.Output(1, fmt.Sprintf("error in validateAndCleanXml: %s", err.Error()))
				return "", err
			}
		}
	}

	if err := encoder.Flush(); err != nil {
		log.Output(1, fmt.Sprintf("error in validateAndCleanXml: %s", err.Error()))
		return "", err
	}

	return buffer.String(), nil
}

func removeUnnecessaryXmlNewlines(xml string) (string, error) {
	// Temporarily replace ">" inside content with "#tempArr"
	tempXml, err := validateAndCleanXml(strings.NewReader(xml))
	//tempXml, err := validateAndCleanXml(xml)
	if err != nil {
		log.Output(1, fmt.Sprintf("error in removeUnnecessaryXmlNewlines: %s", err.Error()))
		return "", err
	}

	// Replace newlines surrounded by spaces
	re1 := regexp.MustCompile(`\s*\n\s*`)
	cleaned := re1.ReplaceAllString(tempXml, " ")

	// Restore newlines after closing tags
	//re2 := regexp.MustCompile(`(>)\s+`)
	//re2 := regexp.MustCompile(`(/?[a-zA-Z0-9/_-]+>)\s+`)
	//aa := re2.ReplaceAllString(cleaned, "$1\n")

	// Restore newlines after closing tags
	re2 := regexp.MustCompile(`>\s+`)
	cleaned = re2.ReplaceAllString(cleaned, ">\n")
	//cleaned = strings.ReplaceAll(cleaned, gt, ">")

	return cleaned, nil
}

func (url Isurl) Xml(w http.ResponseWriter, req *http.Request) {
	var paragraphBreak bool
	var fileReader io.Reader
	var err error
	var result string
	var xmlAcceptedInputs = []string{"xml"}

	if url {
		var urlRequest RequestUrl
		err = json.NewDecoder(req.Body).Decode(&urlRequest)
		if err != nil && ErrorCheck(w, 400, e.ErrJsonInvalid, e.Errors[e.ErrJsonInvalid], err) {
			return
		}
		paragraphBreak = urlRequest.ParagraphBreak
		res, path, statusCode, code := getInputsUrl(urlRequest)
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}
		urlRequest.Url = path
		if !util.IsValidExt(urlRequest.Url, xmlAcceptedInputs, true) {
			ErrorCheck(w, 400, e.ErrFileMissing, e.Errors[e.ErrFileMissing])
			return
		}
		fileReader = res.Body
	} else {
		file, fileHeader, statusCode, code := getInputFile(w, req)
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}
		defer file.Close()
		if strings.TrimSpace(req.FormValue("preserve_paragraphs")) == "true" {
			paragraphBreak = true
		}
		if !util.IsValidExt(fileHeader.Filename, xmlAcceptedInputs, false) {
			ErrorCheck(w, 400, e.ErrFileMissing, e.Errors[e.ErrFileMissing])
			return
		}
		//if !xmlExtChk(fileHeader.Filename) && ErrorCheck(w, 400, e.ErrFileMissing, e.Errors[e.ErrFileMissing]) {
		//	return
		//}
		fileReader = file
	}
	// ***** removing \n before reading ****
	fileBytes, err := io.ReadAll(fileReader)
	if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
		return
	}
	xmlString, err := removeUnnecessaryXmlNewlines(util.CondenseNewlines(string(fileBytes)))
	if err != nil {
		errorOccurred := errors.New("Please provide valid xml content")
		if strings.Contains(req.URL.Path, "multipart") {
			if err != nil && ErrorCheck(w, 400, e.ErrFileMissing, errorOccurred, err) {
				return
			}
		} else {
			if err != nil && ErrorCheck(w, 400, e.ErrURLMissing, errorOccurred, err) {
				return
			}
		}
	}

	result, err = docconv.XMLToText(strings.NewReader(xmlString), []string{}, []string{}, false)
	if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
		log.Output(1, fmt.Sprintf("error is %s", err.Error()))
		return
	}
	//result = replaceSpecialToOriginal(result)

	// ***** end removing \n before reading ****
	if paragraphBreak {
		//respondWithJSON(w, map[string]interface{}{"data": strings.Split(strings.ReplaceAll(condenseNewlines(strings.Replace(result, "\t", "", -1)), "\n \n", "\n"), "\n")})
		//respondWithJSON(w, map[string]interface{}{"data": strings.Split(strings.ReplaceAll(condenseNewlines(strings.Replace(result, "\t", "", -1)), "\n \n", ""), "\n")})
		//respondWithJSON(w, map[string]interface{}{"data": strings.ReplaceAll(condenseNewlines(strings.Replace(result, "\t", "", -1)), "\n \n", "\n")})
		respondWithJSON(w, map[string]interface{}{"data": strings.Trim(util.CondenseNewlines(strings.Replace(result, "\t", "", -1)), "\n")})
		return
	}
	// if want to remove spaces as well, uncomment the following
	respondWithJSON(w, map[string]interface{}{"data": strings.Replace(strings.Replace(strings.Trim(util.CondenseNewlines(result), "\n"), "\n", " ", -1), "\t", "", -1)})
	//respondWithJSON(w, map[string]interface{}{"data": strings.Replace(strings.Replace(result, "\n", " ", -1), "\t", "", -1)})
}

func (url Isurl) Html(w http.ResponseWriter, req *http.Request) {
	var paragraphBreak bool
	//var fileReader io.Reader
	var err error
	var result string
	var htmlString string
	if url {
		var urlRequest RequestUrl
		err = json.NewDecoder(req.Body).Decode(&urlRequest)
		if err != nil && ErrorCheck(w, 400, e.ErrJsonInvalid, e.Errors[e.ErrJsonInvalid], err) {
			return
		}
		paragraphBreak = urlRequest.ParagraphBreak
		urlInput := strings.Trim(urlRequest.Url, " ")
		if urlInput == "" && ErrorCheck(w, 400, e.ErrURLMissing, e.Errors[e.ErrURLMissing]) {
			return
		}
		htmlContent, err := GetHtmlFromSite(urlInput)
		if err != nil {
			navErr := rod.ErrNavigation{Reason: e.Errors[e.ErrURLMissing].Error()}
			if navErr.Is(err) || strings.Contains(err.Error(), "Cannot navigate to invalid URL") {
				log.Output(1, fmt.Sprintf("error in ScrappingProcess: %v", err.Error()))
				ErrorCheck(w, 400, e.ErrURLMissing, e.Errors[e.ErrURLMissing], err)
				return
			}
			log.Output(1, fmt.Sprintf("error in ScrappingProcess: %v", err.Error()))
			ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err)
			return
		}
		htmlString = htmlContent

		//res, path, statusCode, code := getInputsUrl(urlRequest)
		//if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
		//	return
		//}
		//urlRequest.Url = path
		//if !util.IsValidExt(urlRequest.Url, constant.Html, true) && ErrorCheck(w, 400, e.ErrURLMissing, e.Errors[e.ErrURLMissing]) {
		//	return
		//}
		//fileReader = res.Body
		//result, err = docconv.XMLToText(res.Body, []string{}, []string{}, false)
		//if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
		//	return
		//}
	} else {
		file, fileHeader, statusCode, code := getInputFile(w, req)
		if code != -1 && ErrorCheck(w, statusCode, code, e.Errors[code]) {
			return
		}
		defer file.Close()
		if strings.TrimSpace(req.FormValue("preserve_paragraphs")) == "true" {
			paragraphBreak = true
		}

		if !util.IsValidExt(fileHeader.Filename, constant.Html, false) && ErrorCheck(w, 400, e.ErrFileMissing, e.Errors[e.ErrFileMissing]) {
			return
		}
		fileBytes, err := io.ReadAll(file)
		if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
			return
		}
		htmlString = string(fileBytes)

		//result, err = docconv.XMLToText(file, []string{}, []string{}, false)
		//if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
		//	return
		//}
	}
	// ***** removing \n before reading ****
	//var errorOccurred error
	//var rawHtmlErr error
	//var xmlString string

	result, err = textExtract.ExtractHtml(htmlString)
	if err != nil {
		errorOccurred := errors.New("Please provide valid html content")
		if strings.Contains(req.URL.Path, "multipart") {
			if err != nil && ErrorCheck(w, 400, e.ErrFileMissing, errorOccurred, err) {
				return
			}
		} else {
			if err != nil && ErrorCheck(w, 400, e.ErrURLMissing, errorOccurred, err) {
				return
			}
		}
	}

	//result, err = docconv.XMLToText(strings.NewReader(htmlString), []string{}, []string{}, false)
	//if err != nil && ErrorCheck(w, 500, e.ErrTechnicalissue, e.Errors[e.ErrTechnicalissue], err) {
	//	log.Output(1, fmt.Sprintf("error is %s", err.Error()))
	//	return
	//}
	//result = replaceSpecialToOriginal(htmlString)

	// ***** end removing \n before reading ****
	if paragraphBreak {
		//respondWithJSON(w, map[string]interface{}{"data": strings.Split(strings.ReplaceAll(condenseNewlines(strings.Replace(result, "\t", "", -1)), "\n \n", "\n"), "\n")})
		//respondWithJSON(w, map[string]interface{}{"data": strings.Split(strings.ReplaceAll(condenseNewlines(strings.Replace(result, "\t", "", -1)), "\n \n", ""), "\n")})
		//respondWithJSON(w, map[string]interface{}{"data": strings.ReplaceAll(condenseNewlines(strings.Replace(result, "\t", "", -1)), "\n \n", "\n")})
		respondWithJSON(w, map[string]interface{}{"data": strings.Trim(util.CondenseNewlines(strings.Replace(result, "\t", "", -1)), "\n")})
		return
	}
	// if want to remove spaces as well, uncomment the following
	respondWithJSON(w, map[string]interface{}{"data": strings.Replace(strings.Replace(strings.Trim(util.CondenseNewlines(result), "\n"), "\n", " ", -1), "\t", "", -1)})
	//respondWithJSON(w, map[string]interface{}{"data": strings.Replace(strings.Replace(result, "\n", " ", -1), "\t", "", -1)})
}

func xmlExtChk(Filename string) bool {
	if strings.ToLower(filepath.Ext(Filename)) == ".html" {
		return true
	} else if strings.ToLower(filepath.Ext(Filename)) == ".xml" {
		return true
	}
	return false
}

func wordExtChk(Filename string) bool {
	if strings.ToLower(filepath.Ext(Filename)) == ".doc" {
		return true
	} else if strings.ToLower(filepath.Ext(Filename)) == ".docx" {
		return true
	}
	return false
}

func getInputFile(w http.ResponseWriter, r *http.Request) (multipart.File, *multipart.FileHeader, int, int) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		if err.Error() == "http: no such file" {
			return nil, nil, 400, e.ErrFileMissing
		}
		return nil, nil, 500, e.ErrFileMissing
	}
	return file, fileHeader, 200, -1
}

func getInputsUrl(urlRequest RequestUrl) (*http.Response, string, int, int) {
	urlRequest.Url = strings.Trim(urlRequest.Url, " ")

	if urlRequest.Url == "" {
		return nil, "", 400, e.ErrURLMissing
	}

	refinedUrl, err := util.AttachUrlScheme(urlRequest.Url, "https://")
	if err != nil {
		log.Output(1, fmt.Sprintf("error in getInputsUrl: %v", err.Error()))
		return nil, "", 500, e.ErrTechnicalissue
	}
	res, err := http.Get(refinedUrl)
	if err != nil || !(200 <= res.StatusCode && res.StatusCode <= 299) {
		refinedUrl, err := util.AttachUrlScheme(urlRequest.Url, "http://")
		if err != nil {
			log.Output(1, fmt.Sprintf("error in getInputsUrl: %v", err.Error()))
			return nil, "", 500, e.ErrTechnicalissue
		}
		res, err := http.Get(refinedUrl)
		if err != nil || !(200 <= res.StatusCode && res.StatusCode <= 299) {
			return nil, "", 400, e.ErrURLMissing
		}
	}
	return res, refinedUrl, 200, -1
}

//func getInputsUrl(urlRequest RequestUrl) (*http.Response, string, int, int) {
//	urlRequest.Url = strings.Trim(urlRequest.Url, " ")
//
//	if urlRequest.Url == "" {
//		return nil, "", 400, e.ErrURLMissing
//	}
//
//	res, path, statusCode, code := urlParse(urlRequest.Url)
//	if code != -1 {
//		return nil, "", statusCode, code
//	}
//
//	return res, path, 200, -1
//}

//func urlParse(Url string) (*http.Response, string, int, int) {
//	urlparse, err := url.Parse(Url)
//	if err != nil {
//		return nil, "", 500, e.ErrTechnicalissue
//	}
//
//	Urlnew := Url
//	// if urlparse.Host == "" || urlparse.Scheme == "" {
//	// 	fmt.Println("here")
//	// 	return nil, "", 400, e.ErrURLMissing
//	// }
//	if !strings.HasPrefix(Url, "http://") && !strings.HasPrefix(Url, "https://") && urlparse.Scheme == "" {
//		// Check with "https://" prefix
//		Url = "https://" + Urlnew
//		res, err := http.Get(Url)
//		if err != nil || res == nil {
//			// Check with "http://" prefix
//			Url = "http://" + Urlnew
//			_, err := http.Get(Url)
//			if err != nil {
//				return nil, "", 400, e.ErrURLMissing
//			}
//		}
//	}
//
//	res, err := http.Get(Url)
//	if err != nil {
//		return nil, "", 400, e.ErrURLMissing
//	}
//	if res.StatusCode != 200 {
//		return nil, "", 400, e.ErrURLMissing
//	}
//	return res, Url, 200, -1
//}

// This function is used to return the payload to the user in JSON format
func respondWithJSON(w http.ResponseWriter, payload map[string]interface{}) {
	// Converting the payload to JSON
	response, _ := json.Marshal(payload)

	// Set the custom response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

// This function is used to check error
func ErrorCheck(w http.ResponseWriter, statusCode int, code int, err error, rawErr ...error) bool {
	// Error handling
	errInfo := "Error :-"
	if err != nil {
		var logErr error
		if len(rawErr) > 0 && rawErr[0] != nil {
			logErr = rawErr[0]
		} else {
			logErr = err
		}
		logware.CaptureError(w, logErr)
		jsn := map[string]interface{}{
			"error": e.ErrorInfo{
				Message: err.Error(),
				Code:    code,
			},
		}
		if statusCode == 400 {
			errInfo = "Warning :-"
		}
		log.Output(2, errInfo+" Service: File-Extractor"+" Message:"+err.Error())
		response, _ := json.Marshal(jsn)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(response)
		return true
	}
	return false
}
