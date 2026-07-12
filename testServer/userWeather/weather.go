package userweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testServer/gigachat"
	"time"
)

type WeatherResponse struct {
	City        string    `json:"city"`
	Weather     string    `json:"weather"`
	UpdateTime  time.Time `json:"updateTime"`
	InterestingFact string `json:"fact"`
}

func UserWeather(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	city := query.Get("gorod")

	format := query.Get("format")

	response, _ := Temper(city, format)

	w.Header().Set("Content-Type", "application/json")


	interesting, err := gigachat.Promt(city)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	response.InterestingFact = interesting

	json.NewEncoder(w).Encode(response)
}

func Temper(gorod string, format string) (weather *WeatherResponse, err error) {

	baseUrl, err := url.Parse("https://wttr.in/" + gorod)
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
	return &WeatherResponse{
		City:       gorod,
		Weather:    string(body),
		UpdateTime: time.Now(),
	}, nil
}
