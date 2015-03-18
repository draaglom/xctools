package xcassets

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

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

//GenerateAppIconSet takes a source file and a dest Images.xcassets filepath, and overwrites the containing AppIcon.appiconset with a new one containing all the image formats in targets.
func GenerateAppIconSet(source, dest string, targets []Image) error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't get working dir:", err)
		os.Exit(-1)
	}
	if source[:2] == "./" {
		source = strings.Replace(source, "./", dir, 1)
	}
	if dest[:2] == "./" {
		dest = strings.Replace(dest, "./", dir, 1)
	}
	//Have we been given a valid input icon?
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return errors.New(fmt.Sprint("Can't find the source file:", source))
	}
	//Are we in (or given the correct path of) the Gleepost project dir?
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		return errors.New(fmt.Sprint("You didn't specify the project directory correctly -- can't find Images.xcassets here:", dest))
	}
	destDir := dest + "/AppIcon.appiconset"
	//Delete the existing icon if it exists
	if _, err := os.Stat(destDir); err == nil {
		err = os.RemoveAll(destDir)
		if err != nil {
			return errors.New(fmt.Sprint("Error deleting existing icon:", err))
		}
	}
	//Create base dir
	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		return errors.New(fmt.Sprint("Error creating appiconset:", err))
	}

	contents := NewContents()
	for i, t := range targets {
		skip := false
		for _, before := range targets[:i] {
			if before.Px() == t.Px() {
				fmt.Sprint("Already generated this resolution, skipping:", t.Px())
				skip = true
				break
			}
		}
		if !skip {
			fmt.Println("Generating icon for resolution:", t.Px())
			err := resizePNG(source, destDir+"/"+t.Filename, t.Px())
			if err != nil {
				return errors.New(fmt.Sprint("Error resizing icon:", err))
			}
		}
	}
	contents.Images = targets
	f, err := os.Create(destDir + "/Contents.json")
	if err != nil {
		return errors.New(fmt.Sprint("Error creating contents file:", err))
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	err = enc.Encode(contents)
	if err != nil {
		return errors.New(fmt.Sprint("Error writing contents file:", err))
	}
	return nil
}
