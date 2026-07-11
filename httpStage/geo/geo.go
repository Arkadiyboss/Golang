package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	City string `json:"city"`
}

type respJson struct {
	Error bool `json:"error"`
}

func GetLocation(city string) (*Location, error) {
	if city != "1" {
		return &Location{
			City: city,
		}, nil
	}

	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return nil, errors.New("Получен не код 200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var Geo Location

	json.Unmarshal(body, &Geo)
	defer resp.Body.Close()
	return &Geo, nil
}

func CityValidate(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		fmt.Println(err)
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer resp.Body.Close()

	var response respJson

	json.Unmarshal(body, &response)


	return response.Error
}
