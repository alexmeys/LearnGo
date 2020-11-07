package main

// Made a test program for reading meta-data from image and parsing from json.

import (
	"encoding/json"
	"fmt"
)

type meta struct {
	Image struct {
		Width     int    `json:"Width"`
		Height    int    `json:"Height"`
		Title     string `json:"Title"`
		Thumbnail struct {
			URL    string `json:"Url"`
			Height int    `json:"Height"`
			Width  int    `json:"Width"`
		} `json:"Thumbnail"`
		Animated bool  `json:"Animated"`
		IDs      []int `json:"IDs"`
	} `json:"Image"`
}

func main() {
	s := `{
		"Image": {
			"Width": 800,
			"Height": 600,
			"Title": "View from 15th floor",
			"Thumbnail": {
				"Url": "http://www.example.com/image/123",
				"Height": 125,
				"Width": 100
			},
			"Animated": false,
			"IDs": [116, 943, 234, 38793]
		}
	}`

	bs := []byte(s)
	data := meta{}

	err := json.Unmarshal(bs, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("")
	fmt.Println("Our image area:", data.Image.Width*data.Image.Height)
	fmt.Println("Our image title:", data.Image.Title)
	fmt.Println("Animated?", data.Image.Animated)
	fmt.Println("Url:", data.Image.Thumbnail.URL)
	fmt.Println("The ID's of the image:")
	for _, v := range data.Image.IDs {
		fmt.Println(v)
	}
}
