package textExtract

import (
	e "file-extractor/errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
)

func IsValidExt(Filename string, acceptedExt []string) bool {
	parsedUrl, err := url.Parse(Filename)
	if err != nil {
		fmt.Printf("error is %s\n", err.Error())
	}
	fileExt := filepath.Ext(parsedUrl.Path)
	for _, singleImageType := range acceptedExt {
		if fileExt == "."+singleImageType {
			return true
		}
	}
	return false
}
func TestExtension(t *testing.T) {
	_, err := GetDocsExt("")
	fmt.Println(err.Error())
}

func GetDocsExt(Filename string) (string, error) {
	parsedUrl, err := url.Parse(Filename)
	if err != nil {
		return "", err
	}
	return filepath.Ext(parsedUrl.Path), nil
}

func TestExtractImage(t *testing.T) {

	fmt.Printf("ai_service_key %s\n", viper.GetString("azure.ai_service_key"))
	fmt.Printf("ai_service_endpoint %s\n", viper.GetString("azure.ai_service_endpoint"))
	fmt.Printf("ai_service_region %s\n", viper.GetString("azure.ai_service_region"))

	image, err := ExtractDocument("https://c0.piktochart.com/v2/themes/3208-modern-student-resume/snapshot/large.jpg", "url")
	if err != nil {
		fmt.Printf("err is %s", err.Error())
	}
	fmt.Println(image)
}

func TestGetInputsUrl(t *testing.T) {
	req := RequestUrl{"https://freetestdata.com/wp-content/uploads/2023/03/sample_html_file.html", false}
	getInputsUrl(req)
}

type RequestUrl struct {
	Url            string `json:"url"`
	ParagraphBreak bool   `json:"preserve_paragraphs"`
}

func getInputsUrl(urlRequest RequestUrl) (*http.Response, string, int, int) {
	urlRequest.Url = strings.Trim(urlRequest.Url, " ")

	if urlRequest.Url == "" {
		return nil, "", 400, e.ErrURLMissing
	}

	refinedUrl, err := AttachUrlScheme(urlRequest.Url, "https://")
	if err != nil {
		log.Output(1, fmt.Sprintf("error in GetHtmlFromSite: %v", err.Error()))
		return nil, "", 500, e.ErrTechnicalissue
	}
	res, err := http.Get(refinedUrl)
	if err != nil || !(200 <= res.StatusCode && res.StatusCode <= 299) {
		refinedUrl, err := AttachUrlScheme(urlRequest.Url, "http://")
		if err != nil {
			log.Output(1, fmt.Sprintf("error in GetHtmlFromSite: %v", err.Error()))
			return nil, "", 500, e.ErrTechnicalissue
		}
		res, err := http.Get(refinedUrl)
		if err != nil || !(200 <= res.StatusCode && res.StatusCode <= 299) {
			return nil, "", 400, e.ErrURLMissing
		}
	}
	return res, refinedUrl, 200, -1
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
