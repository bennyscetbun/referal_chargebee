package chargebee

import (
	"encoding/json"
	"os"

	chargebeelib "github.com/chargebee/chargebee-go/v3"
	"github.com/ztrue/tracerr"
)

type Config struct {
	ChargeBeeApiKey       string
	ChargeBeeSite         string
	BrevoApiKey           string
	CreditOffertEnCentime int64
	ReductionEnPourcent   float64
	ReferralEmail         ConfigEmail
	CreditAddedEmail      ConfigEmail
}

type ConfigEmail struct {
	Subject    string
	Address    string
	Name       string
	TemplateID int
}

var cfg Config

func Setup(configFilePath string) error {
	// Read the entire file content
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return tracerr.Wrap(err)
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return tracerr.Wrap(err)
	}
	chargebeelib.Configure(cfg.ChargeBeeApiKey, cfg.ChargeBeeSite)
	setupBrevoClient(cfg.BrevoApiKey)
	return nil
}
