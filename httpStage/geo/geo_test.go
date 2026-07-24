package geo_test

import (
	"http/geo"
	"testing"
)

func TestGetLocation(t *testing.T) {
	city := "London"
	expected := &geo.Location{
    City: "London",
}

	got, err := geo.GetLocation(city)

	if err != nil{
		t.Error("Ошибка получения города")
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получили %v", expected, got)
	}
}