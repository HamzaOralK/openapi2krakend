package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"io/fs"
	"strings"
)

func loadFromFile(filePath string) *openapi3.T {
	loader := openapi3.NewLoader()
	openApiDefinition, _ := loader.LoadFromFile(filePath)
	return openApiDefinition
}

func filterFiles(files []fs.FileInfo) []fs.FileInfo {
	for index, file := range files {
		if !(strings.HasSuffix(file.Name(), ".yaml") || strings.HasSuffix(file.Name(), ".json")) {
			files = append(files[:index], files[index+1:]...)
		}
	}
	return files
}

func getComponentFromReferenceAddress(openapiDef openapi3.T, ref string) openapi3.Schema {
	referenceSplit := strings.Split(ref, "/")
	referenceKey := referenceSplit[len(referenceSplit)-1]
	return *openapiDef.Components.Schemas[referenceKey].Value
}

func getExtension(extension map[string]interface{}, key string) string {
	if extension[key] != nil {
		return strings.Replace(fmt.Sprintf("%s", extension[key]), "\"", "", -1)
	} else {
		return ""
	}
}
