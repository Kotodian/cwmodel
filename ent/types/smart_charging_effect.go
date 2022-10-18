package types

type ChargingSchedulePeriod struct {
	From  int64 `json:"from"`
	To    int64 `json:"to"`
	Limit int32 `json:"limit"`
}
