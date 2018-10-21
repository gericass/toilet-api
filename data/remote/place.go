package remote

import (
	"github.com/BurntSushi/toml"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type Place struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	NextPageToken    string        `json:"next_page_token"`
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Viewport struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		Icon string `json:"icon"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Photos []struct {
			Height           int      `json:"height"`
			HTMLAttributions []string `json:"html_attributions"`
			PhotoReference   string   `json:"photo_reference"`
			Width            int      `json:"width"`
		} `json:"photos,omitempty"`
		PlaceID   string   `json:"place_id"`
		Reference string   `json:"reference"`
		Scope     string   `json:"scope"`
		Types     []string `json:"types"`
		Vicinity  string   `json:"vicinity"`
		PlusCode struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code,omitempty"`
		Rating int `json:"rating,omitempty"`
		OpeningHours struct {
			OpenNow bool `json:"open_now"`
		} `json:"opening_hours,omitempty"`
		PriceLevel int `json:"price_level,omitempty"`
	} `json:"results"`
	Status string `json:"status"`
}

type key struct {
	key string
}

type Result struct {
	Coordinate Coordinate `xml:"coordinate"`
	Error      string     `xml:"error"`
}

type Coordinate struct {
	Lat float64 `xml:"lat"`
	Lng float64 `xml:"lng"`
}

func getKey() (string, error) {
	var key key
	_, err := toml.DecodeFile("key.toml", &key)
	if err != nil {
		return "", err
	}
	return key.key, nil
}

func getJson(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getXML(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	return xml.Unmarshal(body, &target)
}

func getLatLngFromKeyWord(keyword string) (*Coordinate, error) {
	result := new(Result)
	url := "https://www.geocoding.jp/api/?v=1.1&q=" + keyword
	err := getXML(url, result)
	if err != nil {
		return nil, err
	}
	return &result.Coordinate, nil
}

func SearchPlaces(keyword string) (*Place, error) {
	coordinate, err := getLatLngFromKeyWord(keyword)
	if err != nil {
		return nil, err
	}
	key, err := getKey()
	if err != nil {
		return nil, err
	}
	lat := fmt.Sprint(coordinate.Lat)
	lng := fmt.Sprint(coordinate.Lng)
	var places *Place
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + lat + "," + lng + "&radius=250&language=ja&key=" + key + "&keyword=トイレ"
	getJson(url, places)
	return places, nil
}
