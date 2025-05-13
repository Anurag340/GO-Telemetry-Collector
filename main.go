package main

import (
	"fmt"

	"github.com/Anurag340/telemetryCollector/Controllers"
	"github.com/Anurag340/telemetryCollector/Database"
	"github.com/gofiber/fiber/v2"
)

func main(){

	app:=fiber.New()

	Database.Connect()

	app.Post("/ingest",Controllers.IngestData)

	err:=app.Listen(":8001")
	if err!=nil{
		panic(err)
	}

	fmt.Println("Server is running on port 8001")

}
