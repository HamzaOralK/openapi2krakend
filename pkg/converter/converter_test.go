package converter

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLoadFromFile(t *testing.T) {

	path := getTestDefinitionPath()
	ans := loadFromFile(path)
	println(ans)
	if len(ans.Servers) != 1 {
		t.Errorf("Got %d servers; expected 1", len(ans.Servers))
	}
}

func getTestDefinitionPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path = strings.ReplaceAll(path, "/pkg/converter", "")
	return fmt.Sprintf("%s/swagger/pet-store.json", path)
}
