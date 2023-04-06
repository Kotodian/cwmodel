package types

import "database/sql/driver"

type ConnectorType int

const (
	ConnectorTypeHouseholdSocket        ConnectorType = 1 // 家用插座
	ConnectorTypeACInterfaceSocket      ConnectorType = 2 // 交流接口插座
	ConnectorTypeACConnectorPlug        ConnectorType = 3 // 交流接口插头
	ConnectorTypeDCConnectorPlug        ConnectorType = 4 // 直流接口枪头
	ConnectorTypeWirelessChargingSocket ConnectorType = 5 // 无线充电插座
	ConnectorTypeOther                  ConnectorType = 6 // 其他
)

type ConnectorState int

const (
	// 不可用
	ConnectorStateUnavailable ConnectorState = 0
	// 不可用
	ConnectorStateAvailable ConnectorState = 1
	// 占用中
	ConnectorStateOccupied ConnectorState = 2
	// 预约中
	ConnectorStateReserved ConnectorState = 3
	// 故障
	ConnectorStateFaulted ConnectorState = 4
)

func (c ConnectorState) Values() []string {
	return []string{"unavailable", "available", "occcupied", "reserved", "faulted"}
}

func (c ConnectorState) String() string {
	return c.Values()[int(c)]
}

func (p ConnectorState) Value() (driver.Value, error) {
	return p, nil
}

// Scan tells our code how to read the enum into our type.
func (p *ConnectorState) Scan(val interface{}) error {
	switch val.(type) {
	case nil:
		return nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		p = val.(*ConnectorState)
	}
	return nil
}
