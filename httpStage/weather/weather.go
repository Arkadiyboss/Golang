package weather

import (
	"errors"
	"fmt"
	"http/geo"
	"io"
	"net/http"
	"net/url"
	"time"
)

type WeatherLocation struct {
	Location   geo.Location `json:"location"`
	Weather    string	`json:"weather"`
	UpdateTime time.Time `json:"updateTime"`
}

func Temper(geo geo.Location, format int) (weather *WeatherLocation, err error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	params := url.Values{}

	params.Add("format", fmt.Sprint(format))

	baseUrl.RawQuery = params.Encode()

	
	resp, err := http.Get(baseUrl.String())
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

	var info WeatherLocation

	info.Weather = string(body)


	return &WeatherLocation{
		Location: geo,
		Weather: info.Weather,
		UpdateTime: time.Now(),
	}, nil
}
