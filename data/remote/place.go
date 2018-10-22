package remote

import (
	"github.com/BurntSushi/toml"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"path/filepath"
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

type PlaceDetail struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Result struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		AdrAddress           string `json:"adr_address"`
		FormattedAddress     string `json:"formatted_address"`
		FormattedPhoneNumber string `json:"formatted_phone_number"`
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
		Icon                     string `json:"icon"`
		ID                       string `json:"id"`
		InternationalPhoneNumber string `json:"international_phone_number"`
		Name                     string `json:"name"`
		OpeningHours struct {
			OpenNow bool `json:"open_now"`
			Periods []struct {
				Close struct {
					Day  int    `json:"day"`
					Time string `json:"time"`
				} `json:"close"`
				Open struct {
					Day  int    `json:"day"`
					Time string `json:"time"`
				} `json:"open"`
			} `json:"periods"`
			WeekdayText []string `json:"weekday_text"`
		} `json:"opening_hours"`
		Photos []struct {
			Height           int      `json:"height"`
			HTMLAttributions []string `json:"html_attributions"`
			PhotoReference   string   `json:"photo_reference"`
			Width            int      `json:"width"`
		} `json:"photos"`
		PlaceID string `json:"place_id"`
		PlusCode struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code"`
		Rating    float64 `json:"rating"`
		Reference string  `json:"reference"`
		Reviews []struct {
			AuthorName              string `json:"author_name"`
			AuthorURL               string `json:"author_url"`
			Language                string `json:"language"`
			ProfilePhotoURL         string `json:"profile_photo_url"`
			Rating                  int    `json:"rating"`
			RelativeTimeDescription string `json:"relative_time_description"`
			Text                    string `json:"text"`
			Time                    int    `json:"time"`
		} `json:"reviews"`
		Scope     string   `json:"scope"`
		Types     []string `json:"types"`
		URL       string   `json:"url"`
		UtcOffset int      `json:"utc_offset"`
		Vicinity  string   `json:"vicinity"`
		Website   string   `json:"website"`
	} `json:"result"`
	Status string `json:"status"`
}

type key struct {
	Key string
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
	path, err := filepath.Abs("key.toml")
	if err != nil {
		return "", err
	}
	_, err = toml.DecodeFile(path, &key)
	if err != nil {
		return "", err
	}
	return key.Key, nil
}

func getJson(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return json.Unmarshal(body, target)
}

func getXML(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return xml.Unmarshal(body, target)
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
	place := new(Place)
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + lat + "," + lng + "&radius=250&language=ja&key=" + key + "&keyword=トイレ"
	err = getJson(url, place)
	if err != nil {
		return nil, err
	}
	return place, nil
}

func GetPlaceDetail(placeId string) (*PlaceDetail, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}
	placeDetail := new(PlaceDetail)
	url := "https://maps.googleapis.com/maps/api/place/details/json?placeid=" + placeId + "&key=" + key
	err = getJson(url, PlaceDetail{})
	if err != nil {
		return nil, err
	}
	return placeDetail, nil
}
