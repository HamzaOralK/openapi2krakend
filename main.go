package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	"github.com/venture-justbuild/openapitokrakend/models"
)

func main() {
	swaggerDirectory := flag.String("directory", "./swagger", "Directory of the swagger files")
	encoding := flag.String("encoding", "json", "Sets default encoding. Values are json, safejson, xml, rss, string, no-op")
	globalTimeout := flag.String("globalTimeout", "3000ms", "Sets global timeout")

	flag.Parse()

	var swaggerFiles []fs.FileInfo
	if files, err := ioutil.ReadDir(*swaggerDirectory); err == nil {
		swaggerFiles = filterFiles(files)
	}

	numberOfFiles := len(swaggerFiles)
	configuration := models.NewConfiguration(*encoding, *globalTimeout)

	for _, file := range swaggerFiles {

		openApiDefinition := loadFromFile(fmt.Sprintf("%s/%s", *swaggerDirectory, file.Name()))

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
