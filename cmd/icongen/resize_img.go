package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/draaglom/xctools/xcassets"
)

func main() {
	var target = flag.String("target", "all", "The icon sizes to generate: options are 'iphone', 'ipad', 'mac', 'ios', 'all'; defaults to 'all'.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s source.png /project/path/to/Images.xcassets\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't get working dir:", err)
		os.Exit(-1)
	}
	source := flag.Arg(0)
	if len(source) == 0 {
		source = "source.png"
	}
	if source[:2] == "./" {
		source = strings.Replace(source, "./", dir, 1)
	}
	dest := flag.Arg(1)
	if len(dest) == 0 {
		dest = "./"
	}
	if dest[:2] == "./" {
		dest = strings.Replace(dest, "./", dir, 1)
	}
	var formats []xcassets.Image
	switch {
	case *target == "all":
		formats = xcassets.All
	case *target == "ios":
		formats = xcassets.IOS
	case *target == "mac":
		formats = xcassets.Mac
	case *target == "ipad":
		formats = xcassets.Ipad
	case *target == "iphone":
		formats = xcassets.Iphone
	default:
		fmt.Println("Not a valid icon set:", *target)
		os.Exit(-1)
	}
	err = xcassets.GenerateAppIconSet(source, dest, formats)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(-1)
	}
}
