package main

import "pgStand/internal/app"

func main() {
	pgApp := app.New()
	pgApp.Run()
}
