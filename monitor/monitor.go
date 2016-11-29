package monitor

import (
	"github.com/rideways/pimonitor/blink"
	"github.com/rideways/pimonitor/colour"
	"github.com/rideways/pimonitor/sensu"
)

// Monitor blink blink
type Monitor struct {
	Blinker          blink.IBlinker
	Checker          sensu.IChecker
	ColourCalculator colour.IColourCalculator
}

// Monitor monitors it
func (b Monitor) Monitor() error {
	checks, err := b.Checker.GetChecks()

	if err != nil {
		b.Blinker.Blink("ff00ff")
		return err
	}

	colour := b.ColourCalculator.Calculator(checks)
	b.Blinker.Blink(colour)
	return nil
}
