package verify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

// WriteResult writes checksum.VerificationResult to a file/http
func WriteResult(result *checksum.VerificationResult, filePath string) error {
	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return WriteString(string(b), "application/json", filePath)
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
