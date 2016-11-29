package sensu

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rideways/pimonitor/rest"
	"github.com/rideways/sensu_status/app/models"
)

// IChecker checks checks
type IChecker interface {
	GetChecks() (*models.SensuCheck, error)
}

// Checker checks checks
type Checker struct {
	Resty         rest.IResty
	SensuAPIURL   string
	Username      string
	Password      string
	UnknownColour string
}

// GetChecks check getter
func (checker Checker) GetChecks() (*models.SensuCheck, error) {
	checkResult := &models.SensuCheck{}

	checker.Resty.SetBasicAuth(checker.Username, checker.Password)
	resp, err := checker.Resty.Get(fmt.Sprintf("%s/events", checker.SensuAPIURL))

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("HTTP Status: %s", resp.Status())
	}

	json.Unmarshal(resp.Body(), &checkResult)
	return checkResult, nil
}
