package main

import (
	"flag"
	"os"

	"github.com/go-resty/resty"
	"github.com/rideways/pimonitor/blink"
	"github.com/rideways/pimonitor/colour"
	"github.com/rideways/pimonitor/sensu"
)

var blinkURL string
var sensuAPIURL string
var username string
var password string
var okColour string
var warningColour string
var criticalColour string

const unknownColour = "800080"

func init() {
	flag.StringVar(&sensuAPIURL, "sensu-api-url", "http://localhost:8080", "sensu api url")
	flag.StringVar(&blinkURL, "blink-url", "http://localhost:4567", "blink(1) server url")
	flag.StringVar(&username, "username", "username", "sensu api username")
	flag.StringVar(&password, "password", "password", "sensu api password")
	flag.StringVar(&okColour, "ok", "00ff00", "ok colour")
	flag.StringVar(&warningColour, "warning", "ffbe00", "warning colour")
	flag.StringVar(&criticalColour, "critical", "ff0000", "critical colour")
	flag.Parse()
}

func main() {
	bb := blink.Blinker{BlinkURL: blinkURL, Resty: resty.New().R()}

	checker := sensu.Checker{
		Resty:         resty.New().R(),
		SensuAPIURL:   sensuAPIURL,
		Username:      username,
		Password:      password,
		UnknownColour: unknownColour}

	colourCalculator := colour.ColourCalculator{
		OkColour:       okColour,
		WarningColour:  warningColour,
		CriticalColour: criticalColour,
		UnknownColour:  unknownColour}

	blink := blink.Blink{
		Blinker:          bb,
		Checker:          checker,
		ColourCalculator: colourCalculator}

	err := blink.DoIt()

	if err != nil {
		os.Exit(1)
	}
}
