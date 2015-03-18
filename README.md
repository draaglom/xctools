# xctools
Some simple utilities for dealing with xcode projects


## icongen -- app icon generator

`icongen` takes a 1024x1024 PNG app icon, and generates an icon set containing every required resolution.

Installation:

```
go install github.com/draaglom/xctools/cmd/icongen
```

```
Usage of icongen:
icongen source.png /project/path/to/Images.xcassets
  -target="all": The icon sizes to generate: options are 'iphone', 'ipad', 'mac', 'ios', 'all'; defaults to 'all'.
```
