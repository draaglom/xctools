package xcassets

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func resizePNG(source, dest string, dim uint) (err error) {
	file, err := os.Open(source)
	if err != nil {
		return
	}
	// decode png into image.Image
	img, err := png.Decode(file)
	if err != nil {
		return
	}
	file.Close()
	m := resize.Thumbnail(dim, dim, img, resize.Lanczos3)
	out, err := os.Create(dest)
	if err != nil {
		return
	}
	defer out.Close()
	// write new image to file
	err = png.Encode(out, m)
	return
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
		return errors.New(fmt.Sprint("Not a valid source file: ", source))
	}
	//Are we in (or given the correct path of) the Images.xcassets dir?
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		return errors.New(fmt.Sprint("Not a valid destination directory:", dest))
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
				fmt.Println("Already generated this resolution, skipping:", t.Px())
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

func FindPath(root, target string) (occurrences []string, err error) {
	occurrences = make([]string, 0)
	wf := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		_, file := filepath.Split(path)
		if file == target {
			occurrences = append(occurrences, path)
		}
		return nil
	}

	err = filepath.Walk(root, wf)
	return
}
