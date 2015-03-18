package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/draaglom/xctools/xcassets"
)

var (
	targets = xcassets.All
)

func main() {
	var source = flag.String("source", "source.png", "The 1024x1024 source icon")
	var dest = flag.String("xcassets", "./Gleepost/Images.xcassets/", "The Images.xcassets folder, defaults to ./Gleepost/Images.xcassets/")
	flag.Parse()
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't get working dir:", err)
		os.Exit(-1)
	}
	sourceStr := *source
	if sourceStr[:2] == "./" {
		*source = strings.Replace(sourceStr, "./", dir, 1)
	}
	destStr := *dest
	if destStr[:2] == "./" {
		*dest = strings.Replace(destStr, "./", dir, 1)
	}
	err = xcassets.GenerateAppIconSet(*source, *dest, xcassets.All)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
