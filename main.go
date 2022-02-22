package main

import "flag"

func main() {
	swagger_directory := flag.String("directory", "./swagger", "Directory of the swagger files")
	encoding := flag.String("encoding", "json", "Sets default encoding. Values are json, safejson, xml, rss, string, no-op")
	global_timeout := flag.String("globalTimeout", "3000ms", "Sets global timeout")

	flag.Parse()

}
