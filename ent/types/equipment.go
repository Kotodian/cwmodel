package types

import "database/sql/driver"

type EquipmentCategory int

const (
	// 私桩
	PrivateEquipment EquipmentCategory = 0
	// 公桩
	PublicEquipment EquipmentCategory = 1
)

func (c EquipmentCategory) Values() []string {
	return []string{"private", "public"}
}

func (c EquipmentCategory) String() string {
	return c.Values()[int(c)]
}

func (p EquipmentCategory) Value() (driver.Value, error) {
	return p, nil
}

// Scan tells our code how to read the enum into our type.
func (p *EquipmentCategory) Scan(val interface{}) error {
	switch val.(type) {
	case nil:
		return nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		p = val.(*EquipmentCategory)
	}
	return nil
}
