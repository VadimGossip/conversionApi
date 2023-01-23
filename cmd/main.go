package main

import (
	"github.com/VadimGossip/conversionApi/internal/app"
	"time"
)

func main() {
	convApi := app.NewApp("Conversion api", time.Now())
	convApi.Run()
}
