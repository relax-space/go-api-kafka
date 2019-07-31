package main

import (
	"go-api-producer/adapters"
	"go-api-producer/config"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

func main() {

	c := config.Init(os.Getenv("APP_ENV"))

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/send", func(c echo.Context) error {
		fruit := &Fruit{Name: "apple"}

		adapters.MessagePublisher.Publish(c.Request().Context(), fruit, adapters.EventCreatedFruit)
		return c.String(http.StatusOK, "sended")
	})

	adapters.NewMessagePublisher(c.EventBroker.Kafka)
	defer adapters.MessagePublisher.Close()

	if err := e.Start(":8090"); err != nil {
		log.Println(err)
	}
}

type Fruit struct {
	Id        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Price     int64     `json:"price"`
	StoreCode string    `json:"storeCode"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}
