package ent

import (
	"time"

	"github.com/Kotodian/cwmodel/ent/types"
	"github.com/Kotodian/protocol/golang/hardware/charger"
)

func (sce *SmartChargingEffect) ParseData() *charger.ChargingProfile {
	periods := make([]*types.ChargingSchedulePeriod, 0)
	targetUnit := sce.Unit
	chargingProfile := &charger.ChargingProfile{
		Id: uint64(sce.SmartID),
		ChargingSchedule: []*charger.ChargingSchedule{{
			ChargingRateUnit:       charger.ChargingRateUnit(charger.ChargingRateUnit_value[targetUnit]),
			ChargingSchedulePeriod: make([]*charger.ChargingSchedule_PeriodType, 0),
		}},
	}
	if sce.ValidFrom != nil {
		validFrom := time.Unix(*sce.ValidFrom, 0)
		chargingProfile.ValidFrom = validFrom.Format(time.RFC3339)
	}
	if sce.ValidTo != nil {
		validTo := time.Unix(*sce.ValidTo, 0)
		chargingProfile.ValidTo = validTo.Format(time.RFC3339)
	}
	smartTime := time.Unix(sce.StartTime, 0)
	chargingProfile.StartTime = smartTime.Format(time.RFC3339)
	for _, period := range periods {
		chargerPeriod := &charger.ChargingSchedule_PeriodType{
			From:  period.From,
			Limit: period.Limit,
			To:    period.To,
		}
		chargingProfile.ChargingSchedule[0].ChargingSchedulePeriod = append(chargingProfile.ChargingSchedule[0].ChargingSchedulePeriod, chargerPeriod)
	}
	return chargingProfile
}
