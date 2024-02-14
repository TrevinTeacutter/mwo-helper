package isc

import (
	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages/series"
)

var _ series.Validator = (*Validator)(nil)

type Validator struct{}

func (v *Validator) Name() string {
	return "ISC"
}

func (v *Validator) Validate(_ series.MatchDetails, _ ...series.MatchDetails) map[string]error {
	validations := make(map[string]error)

	validations["Tonnage Limit"] = nil
	validations["Chassis Limit"] = nil
	validations["Hero Limit"] = nil
	validations["Legend Limit"] = nil
	validations["Tech Base Limit"] = nil
	validations["Player Limit"] = nil

	return validations
}
