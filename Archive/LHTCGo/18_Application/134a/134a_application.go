package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type meta struct {
	Image1 struct {
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
	} `json:"Image1"`
	Image2 struct {
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
	} `json:"Image2"`
	Image3 struct {
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
	} `json:"Image3"`
}

func main() {

	//bad example of password protection, but fun.
	pw := `MyPassword`

	hash, err := bcrypt.GenerateFromPassword([]byte(pw), 4)
	if err != nil {
		fmt.Println(err)
	}

	checkpw := `MyPassword2`
	err = bcrypt.CompareHashAndPassword(hash, []byte(checkpw))
	if err != nil {
		fmt.Println("Illegal action, access denied!")
		fmt.Println(err)
		return
	}

	// You will get here if the password of the raw text matches above...
	s := []byte(`{
		"Image1": {
			"Width": 300,
			"Height": 200,
			"Title": "Building A",
			"Thumbnail": {
				"Url": "http://www.example.com/image/123",
				"Height": 125,
				"Width": 100
			},
			"Animated": false,
			"IDs": [116, 943, 234, 38793]
		},
		"Image2": {
			"Width": 800,
			"Height": 600,
			"Title": "Building B",
			"Thumbnail": {
				"Url": "http://www.example.com/image/124",
				"Height": 125,
				"Width": 100
			},
			"Animated": false,
			"IDs": [117, 944, 235, 38794]
		},
		"Image3": {
			"Width": 500,
			"Height": 500,
			"Title": "Building C",
			"Thumbnail": {
				"Url": "http://www.example.com/image/125",
				"Height": 125,
				"Width": 100
			},
			"Animated": false,
			"IDs": [115, 942, 233, 38792]
		}
	}`)

	var data meta

	err = json.Unmarshal(s, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Image1.Title)
	fmt.Println(data.Image1.Width)
	fmt.Println(data.Image1.Height)
	fmt.Println(data.Image1.Thumbnail.URL)

	fmt.Println(data.Image2.Title)
	fmt.Println(data.Image2.Width)
	fmt.Println(data.Image2.Height)
	fmt.Println(data.Image2.Thumbnail.URL)

	fmt.Println(data.Image3.Title)
	fmt.Println(data.Image3.Width)
	fmt.Println(data.Image3.Height)
	fmt.Println(data.Image3.Thumbnail.URL)

}
