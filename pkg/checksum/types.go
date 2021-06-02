package checksum

import "time"

// TamperedFile is a info of tampered files
type TamperedFile struct {
	FilePath string `json:"filePath"`

	OriginalHash string `json:"originalHash"`
	TamperedHash string `json:"tamperedHash"`
}

// VerificationResult is a result struct of checksum verification
type VerificationResult struct {
	IsTampered    bool           `json:"isTampered"`
	TamperedFiles []TamperedFile `json:"tamperedFiles"`
	Timestamp     *time.Time     `json:"timestamp"`
}
