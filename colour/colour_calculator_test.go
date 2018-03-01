package colour

import (
	"encoding/json"
	"testing"

	"github.com/rideways/sensu_status/app/models"
	"github.com/stretchr/testify/assert"
)

func TestEmptyCheckIsGreen(t *testing.T) {
	x := ColourCalculator{OkColour: "green", WarningColour: "amber", CriticalColour: "red", UnknownColour: "purple"}
	data := models.SensuCheck{}
	assert.Equal(t, "green", x.Calculator(&data), "should be green")
}

func TestSingleWarning(t *testing.T) {
	x := ColourCalculator{OkColour: "green", WarningColour: "amber", CriticalColour: "red", UnknownColour: "purple"}
	data := &models.SensuCheck{}
	y := `[{"check": {"handle":true, "status": 1}}]`

	json.Unmarshal([]byte(y), &data)
	assert.Equal(t, "amber", x.Calculator(data), "should be amber")
}

func TestWarningAndOk(t *testing.T) {
	x := ColourCalculator{OkColour: "green", WarningColour: "amber", CriticalColour: "red", UnknownColour: "purple"}
	data := &models.SensuCheck{}
	y := `[{"check": {"handle":true, "status": 1}},{"check": {"handle":true, "status": 0}},{"check": {"handle":true, "status": 3}}]`

	json.Unmarshal([]byte(y), &data)
	assert.Equal(t, "amber", x.Calculator(data), "should be amber")
}

func TestUnknownAndOk(t *testing.T) {
	x := ColourCalculator{OkColour: "green", WarningColour: "amber", CriticalColour: "red", UnknownColour: "purple"}
	data := &models.SensuCheck{}
	y := `[{"check": {"handle":true, "status": 3}},{"check": {"handle":true, "status": 0}}]`

	json.Unmarshal([]byte(y), &data)
	assert.Equal(t, "purple", x.Calculator(data), "should be purple")
}

func TestWarningAndCritical(t *testing.T) {
	x := ColourCalculator{OkColour: "green", WarningColour: "amber", CriticalColour: "red", UnknownColour: "purple"}
	data := &models.SensuCheck{}
	y := `[{"check": {"handle":true, "status": 1}},{"check": {"handle":true, "status": 2}},{"check": {"handle":true, "status": 3}}]`

	json.Unmarshal([]byte(y), &data)
	assert.Equal(t, "red", x.Calculator(data), "should be red")
}

func TestWarningAndUnhandledCritical(t *testing.T) {
	x := ColourCalculator{OkColour: "green", WarningColour: "amber", CriticalColour: "red", UnknownColour: "purple"}
	data := &models.SensuCheck{}
	y := `[{"check": {"handle":true, "status": 1}},{"check": {"handle":false, "status": 2}}]`

	json.Unmarshal([]byte(y), &data)

	assert.Equal(t, "amber", x.Calculator(data), "should be amber")
}

func TestWarningAndSilencedCritical(t *testing.T) {
	x := ColourCalculator{OkColour: "green", WarningColour: "amber", CriticalColour: "red", UnknownColour: "purple"}
	data := &models.SensuCheck{}
	y := `[{"check": {"handle":true, "status": 1}},{"check": {"handle":true, "status": 2}, "silenced":true}]`

	json.Unmarshal([]byte(y), &data)

	assert.Equal(t, "amber", x.Calculator(data), "should be amber")
}