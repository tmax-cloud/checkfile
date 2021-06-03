package verify

import (
	"encoding/json"
	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"os"
	"path"
)

// WriteToFile writes checksum.VerificationResult to the file
func WriteToFile(result *checksum.VerificationResult, filePath string) error {
	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return WriteStringToFile(string(b), filePath)
}

// WriteStringToFile writes a string to the file
func WriteStringToFile(str string, filePath string) error {
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
