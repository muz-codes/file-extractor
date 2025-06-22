package textExtract

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ExtractPdf(r io.Reader) (string, map[string]string, error) {

	f, err := NewLocalFile(r)
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	bodyResult, metaResult, convertErr := ConvertPDFText(f.Name())
	if convertErr != nil {
		return "", nil, convertErr
	}
	if bodyResult.err != nil {
		return "", nil, bodyResult.err
	}
	if metaResult.err != nil {
		return "", nil, metaResult.err
	}
	fmt.Println("ExtractPdf triggered")
	fmt.Println(bodyResult.body)
	newBody := strings.Replace(bodyResult.body, "\f", "", -1)
	newBody = strings.Replace(newBody, "\n", " ", -1)
	newBody = strings.Replace(newBody, "  ", "\n", -1)
	newBody = strings.Replace(newBody, "\n\n", "", -1)
	newBody = strings.Replace(newBody, "\n\f", "", -1)
	return newBody, metaResult.meta, nil

}

// LocalFile is a type which wraps an *os.File.  See NewLocalFile for more details.
type LocalFile struct {
	*os.File

	unlink bool
}

// NewLocalFile ensures that there is a file which contains the data provided by r.  If r is
// actually an instance of *os.File then this file is used, otherwise a temporary file is
// created and the data from r copied into it.  Callers must call Done() when
// the LocalFile is no longer needed to ensure all resources are cleaned up.
func NewLocalFile(r io.Reader) (*LocalFile, error) {
	if f, ok := r.(*os.File); ok {
		return &LocalFile{
			File: f,
		}, nil
	}

	f, err := ioutil.TempFile(os.TempDir(), "docconv")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary file: %v", err)
	}
	_, err = io.Copy(f, r)
	if err != nil {
		f.Close()
		os.Remove(f.Name())
		return nil, fmt.Errorf("error copying data into temporary file: %v", err)
	}

	return &LocalFile{
		File:   f,
		unlink: true,
	}, nil
}

// Done cleans up all resources.
func (l *LocalFile) Done() {
	l.Close()
	if l.unlink {
		os.Remove(l.Name())
	}
}

// Meta data
type MetaResult struct {
	meta map[string]string
	err  error
}

type BodyResult struct {
	body string
	err  error
}

// Convert PDF

func ConvertPDFText(path string) (BodyResult, MetaResult, error) {
	metaResult := MetaResult{meta: make(map[string]string)}
	bodyResult := BodyResult{}
	mr := make(chan MetaResult, 1)
	go func() {
		metaStr, err := exec.Command("pdfinfo", path).Output()
		if err != nil {
			metaResult.err = err
			mr <- metaResult
			return
		}

		// Parse meta output
		for _, line := range strings.Split(string(metaStr), "\n") {
			if parts := strings.SplitN(line, ":", 2); len(parts) > 1 {
				metaResult.meta[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}

		// Convert parsed meta
		if x, ok := metaResult.meta["ModDate"]; ok {
			if t, ok := pdfTimeLayouts.Parse(x); ok {
				metaResult.meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
		if x, ok := metaResult.meta["CreationDate"]; ok {
			if t, ok := pdfTimeLayouts.Parse(x); ok {
				metaResult.meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}

		mr <- metaResult
	}()

	br := make(chan BodyResult, 1)
	go func() {
		body, err := exec.Command("pdftotext", "-q", "-layout", "-enc", "UTF-8", "-eol", "unix", path, "-").Output()
		if err != nil {
			bodyResult.err = err
		}

		bodyResult.body = string(body)

		br <- bodyResult
	}()

	return <-br, <-mr, nil
}

var pdfTimeLayouts = timeLayouts{time.ANSIC, "Mon Jan _2 15:04:05 2006 MST"}

type timeLayouts []string

func (tl timeLayouts) Parse(x string) (time.Time, bool) {
	for _, layout := range tl {
		t, err := time.Parse(layout, x)
		if err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}
