package xcassets

import (
	"strconv"
	"strings"
)

type Contents struct {
	Images []Image `json:"images"`
	Info   Info    `json:"info"`
}

type Image struct {
	Size     string `json:"size"`
	Idiom    string `json:"idiom"`
	Filename string `json:"filename"`
	Scale    string `json:"scale"`
}

func (i Image) Px() uint {
	size := strings.Split(i.Size, "x")
	sizePX, _ := strconv.ParseUint(size[0], 10, 32)
	scale, _ := strconv.ParseUint(string(i.Scale[0]), 10, 32)
	return uint(sizePX * scale)
}

type Info struct {
	Version int    `json:"version"`
	Author  string `json:"author"`
}

func NewContents() Contents {
	contents := Contents{}
	contents.Info = Info{
		Version: 1,
		Author:  "xcode",
	}
	contents.Images = make([]Image, 0)
	return contents
}
