package checksum

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	// FilesDir is a directory path for the db
	FilesDir = "/etc/checkfile"

	// FilesDB is a db file path
	FilesDB = FilesDir + "/checkfile_db"

	// FilesDBChecksum is a file path of db's checksum
	FilesDBChecksum = FilesDir + "/checkfile_db.sum"

	// TargetFilesEnv is an env. key (comma-separated file path list)
	TargetFilesEnv = "CHECK_FILES"
)

// SaveSumsMap saves sums map into a file
func SaveSumsMap(sums map[string]string) error {
	// Check or create parent dir
	s, err := os.Stat(FilesDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(FilesDir, 0655); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	if err == nil && !s.IsDir() {
		return fmt.Errorf("%s is not a directory", FilesDir)
	}

	// If DB file already exists, stop here
	if _, err := os.Stat(FilesDB); err == nil {
		return fmt.Errorf("DB file already exists")
	}

	// Open DB file
	dbFile, err := os.OpenFile(FilesDB, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = dbFile.Close() }()

	var b bytes.Buffer

	for file, sum := range sums {
		// '<FilePath> <sha1sum>'
		if _, err := b.WriteString(fmt.Sprintf("%s %s\n", file, sum)); err != nil {
			return err
		}
	}

	// Save DB file
	if _, err := dbFile.Write(b.Bytes()); err != nil {
		return err
	}

	// Open checksum file
	sumFile, err := os.OpenFile(FilesDBChecksum, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = sumFile.Close() }()

	dbSum, err := CalculateSum(FilesDB)
	if err != nil {
		return err
	}

	// Save checksum file
	if _, err := sumFile.WriteString(dbSum); err != nil {
		return err
	}

	return nil
}

// LoadSumsMap loads sums map from a file
func LoadSumsMap() (map[string]string, error) {
	// Open DB file
	f, err := os.OpenFile(FilesDB, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	// Open checksum file
	sumFile, err := os.OpenFile(FilesDBChecksum, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer func() { _ = sumFile.Close() }()

	// Verify DB file's checksum
	sumBytes, err := ioutil.ReadAll(sumFile)
	if err != nil {
		return nil, err
	}
	dbSum, err := CalculateSum(FilesDB)
	if err != nil {
		return nil, err
	}
	if string(sumBytes) != dbSum {
		return nil, fmt.Errorf("db is tainted")
	}

	sums := map[string]string{}

	// Read DB files line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		token := strings.Split(t, " ")
		if len(token) != 2 {
			return nil, fmt.Errorf("%s is not in form of '[file path] [sha1 sum]'", t)
		}
		// '<FilePath> <sha1sum>'
		sums[token[0]] = token[1]
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sums, nil
}

// InitSumsDB initiates files' sums database
func InitSumsDB(files []string) error {
	list, err := TargetFiles(files)
	if err != nil {
		return err
	}

	// Calculate sums for files
	sums := map[string]string{}
	for _, f := range list {
		sum, err := CalculateSum(f)
		if err != nil {
			return err
		}
		sums[f] = sum
	}

	// Save it to a file
	return SaveSumsMap(sums)
}

// VerifySums verifies files' checksums
func VerifySums(files []string) error {
	list, err := TargetFiles(files)
	if err != nil {
		return err
	}

	// Load from a file
	sums, err := LoadSumsMap()
	if err != nil {
		return err
	}

	// Verify them
	for _, f := range list {
		expected := sums[f]
		sum, err := CalculateSum(f)
		if err != nil {
			return err
		}

		if expected != sum {
			return fmt.Errorf("file %s is changed (%s -> %s)", f, expected, sum)
		}
	}

	return nil
}

// CalculateSum calculates a sha1sum of a file
func CalculateSum(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	hash := sha1.New()
	if _, err = io.Copy(hash, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// TargetFiles parses environment variables and append all child nodes recursively
func TargetFiles(argList []string) ([]string, error) {
	rawFiles := argList

	// Look up from env.
	if e, ok := os.LookupEnv(TargetFilesEnv); ok {
		list := strings.Split(e, ",")
		for _, f := range list {
			s := strings.TrimSpace(f)
			if s != "" {
				rawFiles = append(rawFiles, s)
			}
		}
	}

	var result []string
	// Add all their children files
	for _, f := range rawFiles {
		children, err := subFiles(f)
		if err != nil {
			return nil, err
		}
		result = append(result, children...)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no target files are specified")
	}
	return result, nil
}

func subFiles(fileName string) ([]string, error) {
	// Absolute path
	if !filepath.IsAbs(fileName) {
		return nil, fmt.Errorf("target file (%s) is not an absolute path", fileName)
	}

	// Exists
	s, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("target file (%s) does not exist", fileName)
	}

	// If it's file, return
	if !s.IsDir() {
		return []string{fileName}, nil
	}

	// If it's a dir, append all children
	var result []string
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		c, err := subFiles(filepath.Join(fileName, f.Name()))
		if err != nil {
			return nil, err
		}
		result = append(result, c...)
	}

	return result, nil
}
