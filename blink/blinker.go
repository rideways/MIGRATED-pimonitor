package blink

import (
	"fmt"

	"github.com/rideways/pimonitor/rest"
)

// IBlinker blinky blinks
type IBlinker interface {
	Blink(colour string)
}

// Blinker blinky blinks
type Blinker struct {
	BlinkURL string
	Resty    rest.IResty
}

// Blink blink blink
func (b Blinker) Blink(colour string) {
	b.Resty.SetQueryParams(map[string]string{"rgb": fmt.Sprintf("#%s", colour)}).Get(fmt.Sprintf("%s/blink1/fadeToRGB", b.BlinkURL))
}
