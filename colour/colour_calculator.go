package colour

import (
	"fmt"
	"log"

	"github.com/rideways/sensu_status/app/models"
)

// IColourCalculator calculates colours
type IColourCalculator interface {
	Calculator(checkResult *models.SensuCheck) string
}

// ColourCalculator calculates colours
type ColourCalculator struct {
	OkColour       string
	WarningColour  string
	CriticalColour string
	UnknownColour  string
}

// Calculator calculates
func (calculator ColourCalculator) Calculator(checkResult *models.SensuCheck) string {
	warn := 0
	critical := 0

	if checkResult == nil {
		log.Println("ok")
		return calculator.OkColour
	}

	for _, v := range *checkResult {
		if v.Check.Handle {
			switch v.Check.Status {
			case 1:
				warn++
			case 2:
				critical++
			}
		}
	}

	if warn == 0 && critical == 0 {
		fmt.Println("ok")
		return calculator.OkColour
	}
	if warn > 0 && critical == 0 {
		fmt.Println("warning")
		return calculator.WarningColour
	}
	if critical > 0 {
		fmt.Println("critical")
		return calculator.CriticalColour
	}
	fmt.Println("unknown")
	return calculator.UnknownColour
}
