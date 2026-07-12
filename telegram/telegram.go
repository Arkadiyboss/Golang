package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Message struct {
	City    string
	Weather string
	Fact    string
	Time    time.Time
}

func main() {

	token, err := os.ReadFile("token")
	if err != nil {
		fmt.Println("Возникла ошибка")
	}

	if string(token) == "" {
		log.Fatal("Переменная окружения TOKEN не установлена")
	}

	pref := tele.Settings{
		Token:  string(token),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(c tele.Context) error {
		c.Send("Привет! Я бот по показу прогноза погоды")
		b.Handle("/city", func(c tele.Context) error {
			c.Send("Введи город по которому хочешь получить прогноз, пока работаю только с английским языком")

			b.Handle("/format", func(c tele.Context) error {
				c.Send("Введи формат который хочешь получить (от 1 до 4)")

				return c.Send("??????????")
			})
		}) return c.Send(?????????????)
	}) return c.Send(??????????????)

	b.Handle("/city", func(c tele.Context) error {
		return c.Send("Введи город по которому хочешь получить прогноз, пока работаю только с английским языком")
	})

	b.Handle("/format", func(c tele.Context) error {
		return c.Send("Введи формат который хочешь получить (от 1 до 4)")
	})

	log.Println("Бот запущен...")
	b.Start()
}
