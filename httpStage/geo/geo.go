package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	City string `json:"city"`
}



func GetLocation(city string) (*Location, error) {
	if city != "" {
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

	return &Geo, nil
}

