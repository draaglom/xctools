package xcassets

var (
	//Iphone contains all the required icon sizes for an iPhone app
	Iphone = []Image{
		{
			Size:     "29x29",
			Idiom:    "iphone",
			Filename: "Icon29.png",
			Scale:    "1x",
		},
		{
			Size:     "29x29",
			Idiom:    "iphone",
			Filename: "Icon58.png",
			Scale:    "2x",
		},
		{
			Size:     "29x29",
			Idiom:    "iphone",
			Filename: "Icon87.png",
			Scale:    "3x"},
		{
			Size:     "40x40",
			Idiom:    "iphone",
			Filename: "Icon80.png",
			Scale:    "2x",
		},
		{
			Size:     "40x40",
			Idiom:    "iphone",
			Filename: "Icon120.png",
			Scale:    "3x"},
		{
			Size:     "57x57",
			Idiom:    "iphone",
			Filename: "Icon57.png",
			Scale:    "1x"},
		{
			Size:     "57x57",
			Idiom:    "iphone",
			Filename: "Icon114.png",
			Scale:    "2x"},
		{
			Size:     "60x60",
			Idiom:    "iphone",
			Filename: "Icon120.png",
			Scale:    "2x"},
		{
			Size:     "60x60",
			Idiom:    "iphone",
			Filename: "Icon180.png",
			Scale:    "3x"},
	}
	//Ipad contains all the required icon sizes for an iPad app
	Ipad = []Image{
		{
			Size:     "29x29",
			Idiom:    "ipad",
			Filename: "Icon29.png",
			Scale:    "1x"},
		{
			Size:     "29x29",
			Idiom:    "ipad",
			Filename: "Icon58.png",
			Scale:    "2x"},
		{
			Size:     "40x40",
			Idiom:    "ipad",
			Filename: "Icon40.png",
			Scale:    "1x"},
		{
			Size:     "40x40",
			Idiom:    "ipad",
			Filename: "Icon80.png",
			Scale:    "2x"},
		{
			Size:     "50x50",
			Idiom:    "ipad",
			Filename: "Icon50.png",
			Scale:    "1x"},
		{
			Size:     "50x50",
			Idiom:    "ipad",
			Filename: "Icon100.png",
			Scale:    "2x"},
		{
			Size:     "72x72",
			Idiom:    "ipad",
			Filename: "Icon72.png",
			Scale:    "1x"},
		{
			Size:     "72x72",
			Idiom:    "ipad",
			Filename: "Icon144.png",
			Scale:    "2x"},
		{
			Size:     "76x76",
			Idiom:    "ipad",
			Filename: "Icon76.png",
			Scale:    "1x"},
		{
			Size:     "76x76",
			Idiom:    "ipad",
			Filename: "Icon152.png",
			Scale:    "2x"},
	}
	//Mac contains all the required icon sizes for a Mac app.
	Mac = []Image{
		{
			Size:     "16x16",
			Idiom:    "mac",
			Filename: "Icon16.png",
			Scale:    "1x"},
		{
			Size:     "16x16",
			Idiom:    "mac",
			Filename: "Icon32.png",
			Scale:    "2x"},
		{
			Size:     "32x32",
			Idiom:    "mac",
			Filename: "Icon32.png",
			Scale:    "1x"},
		{
			Size:     "32x32",
			Idiom:    "mac",
			Filename: "Icon64.png",
			Scale:    "2x"},
		{
			Size:     "128x128",
			Idiom:    "mac",
			Filename: "Icon128.png",
			Scale:    "1x"},
		{
			Size:     "128x128",
			Idiom:    "mac",
			Filename: "Icon256.png",
			Scale:    "2x"},
		{
			Size:     "256x256",
			Idiom:    "mac",
			Filename: "Icon256.png",
			Scale:    "1x"},
		{
			Size:     "256x256",
			Idiom:    "mac",
			Filename: "Icon512.png",
			Scale:    "2x"},
		{
			Size:     "512x512",
			Idiom:    "mac",
			Filename: "Icon512.png",
			Scale:    "1x"},
		{
			Size:     "512x512",
			Idiom:    "mac",
			Filename: "Icon1024.png",
			Scale:    "2x"},
	}
	//IOS = iphone, ipad
	IOS = append(Iphone, Ipad...)
	//All is a preset containing every possible icon size.
	All = append(IOS, Mac...)
)

//LaunchImages are presets for iPhone/iPad launch images.
var LaunchImages = []Image{
	{
		Orientation:          "portrait",
		Idiom:                "iphone",
		Extent:               "full-screen",
		MinimumSystemVersion: "7.0",
		Filename:             "smallsplash.png",
		Scale:                "2x",
	},
	{
		Extent:               "full-screen",
		Idiom:                "iphone",
		Subtype:              "retina4",
		Filename:             "big_splash.png",
		MinimumSystemVersion: "7.0",
		Orientation:          "portrait",
		Scale:                "2x",
	},
	{
		Orientation:          "portrait",
		Idiom:                "ipad",
		Extent:               "full-screen",
		MinimumSystemVersion: "7.0",
		Scale:                "1x",
	},
	{
		Orientation:          "landscape",
		Idiom:                "ipad",
		Extent:               "full-screen",
		MinimumSystemVersion: "7.0",
		Scale:                "1x",
	},
	{
		Orientation:          "portrait",
		Idiom:                "ipad",
		Extent:               "full-screen",
		MinimumSystemVersion: "7.0",
		Scale:                "2x",
	},
	{
		Orientation:          "landscape",
		Idiom:                "ipad",
		Extent:               "full-screen",
		MinimumSystemVersion: "7.0",
		Scale:                "2x",
	},
	{
		Orientation: "portrait",
		Idiom:       "iphone",
		Extent:      "full-screen",
		Filename:    "Default.png",
		Scale:       "1x",
	},
	{
		Orientation: "portrait",
		Idiom:       "iphone",
		Extent:      "full-screen",
		Filename:    "Default@2x.png",
		Scale:       "2x",
	},
	{
		Orientation: "portrait",
		Idiom:       "iphone",
		Extent:      "full-screen",
		Filename:    "Default-568h@2x.png",
		Subtype:     "retina4",
		Scale:       "2x",
	},
	{
		Orientation: "portrait",
		Idiom:       "ipad",
		Extent:      "to-status-bar",
		Scale:       "1x",
	},
	{
		Orientation: "landscape",
		Idiom:       "ipad",
		Extent:      "to-status-bar",
		Scale:       "1x",
	},
	{
		Orientation: "portrait",
		Idiom:       "ipad",
		Extent:      "to-status-bar",
		Scale:       "2x",
	},
	{
		Orientation: "landscape",
		Idiom:       "ipad",
		Extent:      "to-status-bar",
		Scale:       "2x",
	},
}
