package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/draaglom/xctools/xcassets"
	"github.com/nfnt/resize"
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
	//Have we been given a valid input icon?
	if _, err := os.Stat(*source); os.IsNotExist(err) {
		fmt.Println("Can't find the source file:", *source)
		os.Exit(-1)
	}
	//Are we in (or given the correct path of) the Gleepost project dir?
	if _, err := os.Stat(*dest); os.IsNotExist(err) {
		fmt.Println("You didn't specify the project directory correctly -- can't find Images.xcassets here:", *dest)
		os.Exit(-1)
	}
	destDir := *dest + "/AppIcon.appiconset"
	//Delete the existing icon if it exists
	if _, err := os.Stat(destDir); err == nil {
		err = os.RemoveAll(destDir)
		if err != nil {
			fmt.Println("Error deleting existing icon:", err)
			os.Exit(-1)
		}
	}
	//Create base dir
	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		fmt.Println("Error creating appiconset:", err)
		os.Exit(-1)
	}

	contents := xcassets.NewContents()
	for i, t := range targets {
		skip := false
		for _, before := range targets[:i] {
			if before.Px() == t.Px() {
				fmt.Println("Already generated this resolution, skipping:", t.Px())
				skip = true
				break
			}
		}
		if !skip {
			fmt.Println("Generating icon for resolution:", t.Px())
			err := resizePNG(*source, destDir+"/"+t.Filename, t.Px())
			if err != nil {
				fmt.Println("Error resizing icon:", err)
				os.Exit(-1)
			}
		}
	}
	contents.Images = targets
	f, err := os.Create(destDir + "/Contents.json")
	if err != nil {
		fmt.Println("Error creating contents file:", err)
		os.Exit(-1)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	err = enc.Encode(contents)
	if err != nil {
		log.Println("Error writing contents file:", err)
		os.Exit(-1)
	}
}

func resizePNG(source, dest string, dim uint) (err error) {
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	// decode png into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	m := resize.Thumbnail(dim, dim, img, resize.Lanczos3)
	out, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// write new image to file
	png.Encode(out, m)
	return nil
}
