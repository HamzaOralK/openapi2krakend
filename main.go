package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/venture-justbuild/openapitokrakend/models"
	"io/fs"
	"io/ioutil"
	"strings"
)

func FilterFiles(files []fs.FileInfo) []fs.FileInfo {
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

func main() {
	swaggerDirectory := flag.String("directory", "./swagger", "Directory of the swagger files")
	encoding := flag.String("encoding", "json", "Sets default encoding. Values are json, safejson, xml, rss, string, no-op")
	globalTimeout := flag.String("globalTimeout", "3000ms", "Sets global timeout")

	flag.Parse()

	var swaggerFiles []fs.FileInfo
	if files, err := ioutil.ReadDir(*swaggerDirectory); err == nil {
		swaggerFiles = FilterFiles(files)
	}

	numberOfFiles := len(swaggerFiles)
	configuration := models.NewConfiguration(*encoding, *globalTimeout)

	loader := openapi3.NewLoader()
	for _, file := range swaggerFiles {

		openApiDefinition, _ := loader.LoadFromFile(fmt.Sprintf("%s/%s", *swaggerDirectory, file.Name()))

		host := openApiDefinition.Servers[0].URL
		path := strings.ToLower(openApiDefinition.Info.Title)
		apiTimeout := *globalTimeout

		if extensionValue := getExtension(openApiDefinition.Extensions, "x-timeout"); extensionValue != "" {
			apiTimeout = extensionValue
		}

		for pathUrl, pathItem := range openApiDefinition.Paths {
			for methodName, methodObject := range pathItem.Operations() {

				endpoint := fmt.Sprintf("%s", pathUrl)
				if numberOfFiles > 1 {
					endpoint = fmt.Sprintf("/%s%s", path, pathUrl)
				}

				methodTimeout := apiTimeout
				if extensionValue := getExtension(methodObject.Extensions, "x-timeout"); extensionValue != "" {
					methodTimeout = extensionValue
				}

				krakendEndpoint := models.NewEndpoint(host, endpoint, pathUrl, methodName, *encoding, methodTimeout)
				lengthOfSecurity := len(*methodObject.Security)
				if lengthOfSecurity > 0 {
					krakendEndpoint.InsertHeadersToPass("Authorization")
				}
				for _, parameterObject := range methodObject.Parameters {
					parameter := parameterObject.Value
					if parameter.In == "query" {
						if *parameter.Explode == true && parameter.Schema.Ref != "" {
							explodedParams := getComponentFromReferenceAddress(*openApiDefinition, parameter.Schema.Ref)
							if explodedParams.Type == "object" {
								for k, _ := range explodedParams.Properties {
									krakendEndpoint.InsertQuerystringParams(k)
								}
							} else if explodedParams.Type == "Array" {
								krakendEndpoint.InsertQuerystringParams(parameter.Name)
							}
						} else {
							krakendEndpoint.InsertQuerystringParams(parameter.Name)
						}
					}
				}

				configuration.InsertEndpoint(krakendEndpoint)
			}
		}
	}

	file, _ := json.MarshalIndent(configuration, "", " ")
	_ = ioutil.WriteFile("krakend.json", file, 0644)
}
