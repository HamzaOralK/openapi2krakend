package main

import (
	"encoding/json"
	"flag"
	"github.com/okhuz/openapi2krakend/pkg/converter"
	"github.com/okhuz/openapi2krakend/pkg/utility"
	"os"
	"path"
)

func main() {
	swaggerDirectory := flag.String("directory", "./swagger", "Directory of the swagger files")

	flag.Parse()

	encoding := utility.GetEnv("ENCODING", "json");
	globalTimeout := utility.GetEnv("GLOBAL_TIMEOUT", "3000ms")

	configuration := converter.Convert(*swaggerDirectory, encoding, globalTimeout)

	file, _ := json.MarshalIndent(configuration, "", " ")
	_ = os.MkdirAll(path.Join(path.Base(""), "output"), 0777)
	_ = os.WriteFile("output/krakend.json", file, 0644)
}
