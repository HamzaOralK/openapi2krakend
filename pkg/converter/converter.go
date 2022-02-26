package converter

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	"github.com/okhuz/openapi2krakend/pkg/extensions"
	"github.com/okhuz/openapi2krakend/pkg/models"
)

func Convert(swaggerDirectory string, encoding string, globalTimeout string) models.Configuration {
	var swaggerFiles []fs.FileInfo
	if files, err := ioutil.ReadDir(swaggerDirectory); err == nil {
		swaggerFiles = filterFiles(files)
	}

	numberOfFiles := len(swaggerFiles)
	configuration := models.NewConfiguration(encoding, globalTimeout)

	for _, file := range swaggerFiles {

		openApiDefinition := loadFromFile(fmt.Sprintf("%s/%s", swaggerDirectory, file.Name()))

		host := openApiDefinition.Servers[0].URL
		path := strings.ToLower(openApiDefinition.Info.Title)
		apiTimeout := globalTimeout

		if extensionValue := getExtension(openApiDefinition.Extensions, extensions.TimeOut); extensionValue != "" {
			apiTimeout = extensionValue
		}

		for pathUrl, pathItem := range openApiDefinition.Paths {
			for methodName, methodObject := range pathItem.Operations() {

				endpoint := fmt.Sprintf("%s", pathUrl)
				if numberOfFiles > 1 {
					endpoint = fmt.Sprintf("/%s%s", path, pathUrl)
				}

				methodTimeout := apiTimeout
				if extensionValue := getExtension(methodObject.Extensions, extensions.TimeOut); extensionValue != "" {
					methodTimeout = extensionValue
				}

				krakendEndpoint := models.NewEndpoint(host, endpoint, pathUrl, methodName, encoding, methodTimeout)
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
	return configuration
}