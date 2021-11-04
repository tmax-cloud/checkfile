package verify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/tmax-cloud/checkfile/pkg/checksum"
)

// WriteResult writes checksum.VerificationResult to a file/http
func WriteResult(result *checksum.VerificationResult, filePath, fileTemplates []string) error {
	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return WriteStrings(string(b), "application/json", filePath, fileTemplates)
}

// WriteStrings writes a string to the files
func WriteStrings(str, contentType string, filePaths, fileTemplates []string) error {
	var err error = nil
	for i, p := range filePaths {
		tmpl, err3 := template.New("body").Parse(fileTemplates[i])
		if err3 != nil {
			err = err3
			continue
		}
		var buf bytes.Buffer
		if err3 := tmpl.Execute(&buf, str); err3 != nil {
			err = err3
			continue
		}
		if err2 := WriteString(buf.String(), contentType, p); err2 != nil {
			err = err2
		}
	}
	return err
}

// WriteString writes a string to the file
func WriteString(str, contentType string, filePath string) error {
	// Handle HTTP
	if strings.HasPrefix(filePath, "http://") || strings.HasPrefix(filePath, "https://") {
		res, err := http.Post(filePath, contentType, bytes.NewBufferString(str))
		if err != nil {
			return err
		}
		if res.StatusCode < 200 || res.StatusCode >= 400 {
			result, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}
			return fmt.Errorf("error writing error to %s: %d(%s)", filePath, res.StatusCode, string(result))
		}
		return nil
	}

	// Handle file
	f, err := openFile(filePath)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(str + "\n"); err != nil {
		return err
	}
	return nil
}

func openFile(filePath string) (*os.File, error) {
	if err := os.MkdirAll(path.Dir(filePath), 0755); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}
