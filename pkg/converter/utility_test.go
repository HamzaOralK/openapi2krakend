package converter

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLoadFromFile(t *testing.T) {

	path := getSwaggerFolder()
	openApiDefinition := loadFromFile(fmt.Sprintf("%s/pet-store.json", path))

	if openApiDefinition == nil {
		t.Error("Got nil servers; expected openApiDefinition")
	}
}

func TestFilterFiles(t *testing.T) {
	var swaggerFiles []fs.FileInfo
	if files, err := ioutil.ReadDir(getSwaggerFolder()); err == nil {
		swaggerFiles = filterFiles(files)
	}
	numberOfFiles := len(swaggerFiles)
	if numberOfFiles != 1 {
		t.Errorf("Got %d #files; expected 1", numberOfFiles)

	}
}

func getSwaggerFolder() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path = strings.ReplaceAll(path, "/pkg/converter", "")
	return fmt.Sprintf("%s/swagger", path)
}
