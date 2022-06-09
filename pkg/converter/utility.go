package converter

import (
	"fmt"
	"github.com/okhuz/openapi2krakend/pkg/models"
	"io/fs"
	"log"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/okhuz/openapi2krakend/pkg/extensions"
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

func getExtension(extension map[string]interface{}, key extensions.CustomExtensions) string {
	keyString := key.String()
	if extension[keyString] != nil {
		return strings.Replace(fmt.Sprintf("%s", extension[keyString]), "\"", "", -1)
	} else {
		return ""
	}
}

func sanitizeTitle(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToLower(reg.ReplaceAllString(input, ""))
}

func insertQueryParams(refComponent *openapi3.Schema, openApiDefinition *openapi3.T, krakendEndpoint *models.Endpoint) {
	for k, _ := range refComponent.Properties {
		krakendEndpoint.InsertQuerystringParams(k)
	}
	for _, i := range refComponent.AllOf {
		if i.Ref != "" {
			newParameterObject := getComponentFromReferenceAddress(*openApiDefinition, i.Ref)
			insertQueryParams(&newParameterObject, openApiDefinition, krakendEndpoint)
		}
	}
}
