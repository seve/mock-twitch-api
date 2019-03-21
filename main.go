package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/helix/subscriptions", fakeThatShit)

	e.Logger.Error(e.Start(":1323"))
}

type Sub struct {
	BroadcasterID   string `json:"broadcaster_id"`
	BroadcasterName string `json:"broadcaster_name"`
	IsGift          bool   `json:"is_gift"`
	Tier            string `json:"tier"`
	PlanName        string `json:"plan_name"`
	UserID          string `json:"user_id"`
	UserName        string `json:"user_name"`
}

func fakeThatShit(c echo.Context) error {
	broadcasterID := c.QueryParam("broadcaster_id")
	planNames := [4]string{"levelOne", "levelTwo", "levelThree", "levelFour"}
	tierNum := rand.Intn(3)
	data := make([]Sub, rand.Intn(1000))
	for i := 0; i < len(data); i++ {
		data[i] = Sub{
			BroadcasterID:   broadcasterID,
			BroadcasterName: broadcasterID,
			IsGift:          rand.Float32() < 0.5,
			Tier:            string(tierNum + 1*1000),
			PlanName:        planNames[tierNum],
			UserID:          strconv.Itoa(rand.Intn(1000)),
			UserName:        "dick" + strconv.Itoa(rand.Intn(1000)),
		}
	}
	return c.JSON(http.StatusOK, data)
}
