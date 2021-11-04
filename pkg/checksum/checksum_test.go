package checksum

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/tmax-cloud/checkfile/internal/utils"
)

func TestTargetFiles(t *testing.T) {

}

func TestCalculateSum(t *testing.T) {
	testFile := os.TempDir() + "/checksum_test_" + utils.RandomString(5)
	testContent := "tetetetetestContentntntns\n\n\n\nTEST!!!!123123123123"
	testSum := "9cb99952e0706d4779691da86e8b5a354fb696c2"

	if err := ioutil.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.Remove(testFile); err != nil {
			t.Fatal(err)
		}
	}()

	sum, err := CalculateSum(testFile)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, testSum, sum)
}
