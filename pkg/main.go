package main

import (
	"encoding/json"
	"flag"
	"os"
	"path"

	"github.com/venture-justbuild/openapitokrakend/pkg/converter"
)

func main() {
	swaggerDirectory := flag.String("directory", "./swagger", "Directory of the swagger files")
	encoding := flag.String("encoding", "json", "Sets default encoding. Values are json, safejson, xml, rss, string, no-op")
	globalTimeout := flag.String("globalTimeout", "3000ms", "Sets global timeout")

	flag.Parse()

	configuration := converter.Convert(*swaggerDirectory, *encoding, *globalTimeout)

	file, _ := json.MarshalIndent(configuration, "", " ")
	//emptyLine := []byte("\n")
	//file = append(file, emptyLine...)
	_ = os.MkdirAll(path.Join(path.Base(""), "output"), 0777)
	_ = os.WriteFile("output/krakend.json", file, 0644)
}
