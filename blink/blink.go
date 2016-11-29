package blink

import (
	"github.com/rideways/pimonitor/colour"
	"github.com/rideways/pimonitor/sensu"
)

// Blink blink blink
type Blink struct {
	Blinker          IBlinker
	Checker          sensu.IChecker
	ColourCalculator colour.IColourCalculator
}

// DoIt does it
func (b Blink) DoIt() error {
	checks, err := b.Checker.GetChecks()

	if err != nil {
		b.Blinker.Blink("ff00ff")
		return err
	}

	colour := b.ColourCalculator.Calculator(checks)
	b.Blinker.Blink(colour)
	return nil
}
